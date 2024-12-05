package model

type RegisterUser struct {
	Password string
	Email    string
}

type LoginUser struct {
	Email    string
	Password string
}

type UserVerified struct {
	Token  string `json:"token"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
	UserId string `json:"userId"`
}
