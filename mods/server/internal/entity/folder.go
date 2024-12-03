package entity

import "time"

type Folder struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`                 // 文件夹的唯一 ID
	ParentID  *uint     `gorm:"index"`                                    // 父文件夹 ID, NULL 表示根目录
	Name      string    `gorm:"not null"`                                 // 文件夹名称
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`                // 创建时间
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime"` // 更新时间
	Parent    *Folder   `gorm:"foreignKey:ParentID"`                      // 外键关联到父文件夹
	Children  []Folder  `gorm:"foreignKey:ParentID"`                      // 关联的子文件夹
	Files     []File    `gorm:"foreignKey:FolderID"`                      // 关联的文件
}
