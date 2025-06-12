package buffer

import (
	"bytes"
	"fmt"
	"go.uber.org/zap"
	"os"
	"path"
	"sync"
	"time"
)

type ChunkBuffer struct {
	sync.Mutex
	buf        []byte
	maxSize    int
	currSize   int
	chunkIndex int
	ns         string
	fileKey    string
	tmpDir     string
	logger     *zap.SugaredLogger
}

func (b *ChunkBuffer) WriteChunk(data []byte) error {
	b.Lock()
	defer b.Unlock()

	if len(data)+b.currSize > b.maxSize {
		if err := b.Flush(); err != nil {
			return err
		}
	}

	b.buf = append(b.buf, data...)
	b.currSize += len(data)
	b.chunkIndex++
	return nil
}

func (b *ChunkBuffer) Flush() error {
	if b.currSize == 0 {
		return nil
	}

	chunkPath := path.Join(b.tmpDir, fmt.Sprintf("%s-batch_%d", b.fileKey, time.Now().UnixNano()))
	err := os.WriteFile(chunkPath, b.buf, 0644)
	if err != nil {
		return err
	}
	b.logger.Infof("flushed chunk buffer to %s", chunkPath)

	// Reset buffer
	b.buf = b.buf[:0]
	b.currSize = 0
	return nil
}

type UploadBuffer struct {
	mu        sync.Mutex
	buffer    map[string]*bytes.Buffer
	capacity  int64
	used      int64
	evictChan chan string
}
