package entity

import "gorm.io/gorm"

type UserGroup struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null"`
}
