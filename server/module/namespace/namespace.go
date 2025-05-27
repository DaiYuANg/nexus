package namespace

import "time"

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

type Namespace struct {
	Name           string
	Description    string
	AllowDuplicate bool
	OwnerID        string    // 拥有者用户ID，方便权限管理
	CreatedAt      time.Time // 创建时间
	UpdatedAt      time.Time // 更新时间

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
