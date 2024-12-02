package entity

import "gorm.io/gorm"

type FileResource struct {
	gorm.Model
	Md5    string
	Bucket string
	Object string
}
