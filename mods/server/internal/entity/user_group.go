package entity

type UserGroup struct {
	BaseModel
	Name string `gorm:"type:varchar(20);not null"`
}
