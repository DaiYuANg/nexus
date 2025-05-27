package storage

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
