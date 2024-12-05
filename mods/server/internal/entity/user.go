package entity

type User struct {
	BaseModel
	Username string
	Password string
	Email    string
	Avatar   string
}
