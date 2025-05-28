package storage

import (
	"encoding/binary"
	"fmt"
	"github.com/dgraph-io/badger/v4"
	"github.com/google/uuid"
	"github.com/influxdata/influxdb/pkg/snowflake"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"os"
	"path"
)

type Chunk struct {
	Index int    // 块编号，0开始
	Data  []byte // 块数据
	Size  int64  // 块大小
}

type FileMeta struct {
	FileName       string
	TotalSize      int64
	ChunkSize      int64
	ChunkCount     int
	UploadedChunks map[int]bool // 已上传的块索引
}

const ChunkSize = 5 * 1024 * 1024

type FileChunker struct {
	store       *badger.DB
	idGenerator *snowflake.Generator
}

func (f FileChunker) ChunkSave(header *multipart.FileHeader, ns string, logger *zap.SugaredLogger) (int, error) {
	file, err := header.Open()
	if err != nil {
		return 0, err
	}

	fileKey := fmt.Sprintf("%s-%s", ns, uuid.New().String())
	chunkDir := path.Join(os.TempDir(), "chunks", ns, uuid.New().String())
	if err := os.MkdirAll(chunkDir, os.ModePerm); err != nil {
		return 0, err
	}

	buf := make([]byte, ChunkSize)
	part := 0

	var keys []string
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			return 0, err
		}
		if n == 0 {
			break
		}

		chunkKey := fmt.Sprintf("%s-chunk-%d", fileKey, part)
		keys = append(keys, chunkKey)
		chunkPath := path.Join(chunkDir, fmt.Sprintf("chunk_%d", part))
		logger.Debugf("uploading chunk: %s", chunkPath)

		if writeErr := os.WriteFile(chunkPath, buf[:n], 0644); writeErr != nil {
			return 0, writeErr
		}

		if err != nil {
			return 0, err
		}
		part++
	}

	id := f.idGenerator.Next()

	logger.Debugf("uploading chunk: %d", id)

	buf = make([]byte, 8)

	binary.BigEndian.PutUint64(buf, id)

	err = f.store.Update(func(txn *badger.Txn) error {
		for _, key := range keys {
			err := txn.Set(buf, []byte(key))
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return part, nil
}
