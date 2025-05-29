package bucket

import (
	"github.com/samber/lo"
	"time"
)

type StorageBackendType string

const (
	StorageBackendMemory StorageBackendType = "memory"
	StorageBackendLocal  StorageBackendType = "local"
	StorageBackendMinIO  StorageBackendType = "minio"
	StorageBackendSFTP   StorageBackendType = "sftp"
)

type StorageMode string

const (
	StorageModeWholeFile StorageMode = "whole_file" // 整文件存储
	StorageModeChunk     StorageMode = "chunk"      // chunk分块存储
	StorageModeObject    StorageMode = "object"     // 对象存储（类似S3）
)

type Bucket struct {
	Name           string
	Description    string
	AllowDuplicate bool
	OwnerID        string
	CreatedAt      time.Time
	UpdatedAt      time.Time

	IsActive   bool
	QuotaBytes int64
	UsedBytes  int64

	AccessControl string

	Tags           []string
	StorageBackend StorageBackendType `json:"storage_backend"`
	StorageConfig  map[string]string  `json:"storage_config"`

	// 存储模式
	StorageMode StorageMode `json:"storage_mode"`
}

func New(name string, opts ...func(*Bucket)) *Bucket {
	ns := &Bucket{
		Name:           name,
		CreatedAt:      time.Now(),
		IsActive:       true,
		AllowDuplicate: true,
	}
	lo.ForEach(opts, func(opt func(*Bucket), index int) {
		opt(ns)
	})
	return ns
}

func WithQuota(bytes int64) func(*Bucket) {
	return func(n *Bucket) {
		n.QuotaBytes = bytes
	}
}

func WithStorage(backend StorageBackendType, config map[string]string) func(*Bucket) {
	return func(n *Bucket) {
		n.StorageBackend = backend
		n.StorageConfig = config
	}
}
