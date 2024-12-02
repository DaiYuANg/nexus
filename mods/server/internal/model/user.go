package model

type RegisterUser struct {
	Password string
	Email    string
}

type LoginUser struct {
	Email    string
	Password string
}
