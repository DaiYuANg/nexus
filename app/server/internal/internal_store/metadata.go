package internal_store

import "time"

type FileMetadata struct {
	ID          string            // 唯一 ID，可使用 UUID 或自定义 snowflake
	Namespace   string            // 所属命名空间
	Name        string            // 原始文件名
	Size        int64             // 文件大小（字节）
	MimeType    string            // 文件类型，如 image/png
	Hash        string            // 文件内容哈希（如 SHA-256）
	Chunks      int               // 若为 chunked 存储，记录块数
	StorageType string            // 存储方式，如 "local", "minio", "memory", "sftp"
	StoragePath string            // 存储路径或对象键值（根据后端定）
	IsChunked   bool              // 是否采用 chunked 存储
	UploadedAt  time.Time         // 上传时间
	UploadedBy  string            // 用户 ID 或 token
	Metadata    map[string]string // 可扩展字段，如 exif、lang、custom_tag
	Status      string            // 状态标记：pending, processing, done, failed
}
