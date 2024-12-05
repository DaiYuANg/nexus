package entity

type File struct {
	BaseModel
	FolderID       uint   `gorm:"index"`                 // 文件所属的文件夹 ID
	Name           string `gorm:"not null"`              // 文件名称
	FileType       string `gorm:"type:varchar(50)"`      // 文件类型（如 image/png, text/plain）
	Size           int64  `gorm:"size:bigint"`           // 文件大小（字节数）
	Path           string `gorm:"type:varchar(1024)"`    // 文件存储路径（相对或绝对路径）
	UploadedBy     uint   `gorm:"index"`                 // 上传者 ID（用户 ID）
	Folder         Folder `gorm:"foreignKey:FolderID"`   // 外键关联到文件夹
	UploadedByUser User   `gorm:"foreignKey:UploadedBy"` // 外键关联到用户表（如果有）
}
