package entity

type Folder struct {
	BaseModel
	ParentID *uint    `gorm:"index"`               // 父文件夹 ID, NULL 表示根目录
	Name     string   `gorm:"not null"`            // 文件夹名称
	Parent   *Folder  `gorm:"foreignKey:ParentID"` // 外键关联到父文件夹
	Children []Folder `gorm:"foreignKey:ParentID"` // 关联的子文件夹
	Files    []File   `gorm:"foreignKey:FolderID"` // 关联的文件
}
