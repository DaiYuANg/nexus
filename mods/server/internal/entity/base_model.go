package entity

import (
	"gorm.io/gorm"
	"strconv"
)

type BaseModel struct {
	gorm.Model
}

func (m BaseModel) Id2String() string {
	return strconv.Itoa(int(m.ID))
}
