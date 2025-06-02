package buffer

import (
	"bytes"
	"sync"
)

type UploadBuffer struct {
	mu        sync.Mutex
	buffer    map[string]*bytes.Buffer
	capacity  int64
	used      int64
	evictChan chan string
}
