package entity

import "gorm.io/gorm"

type FileStructure struct {
	gorm.Model
	userId int
}
