package service

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var mySigningKey = []byte("mysecretkey")

type JWTService struct {
}
type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func (s *JWTService) Sign(userID string) (string, error) {
	// 创建自定义的Claims
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "myapp",                                            // JWT的发行者
			Subject:   "user authentication",                              // JWT的主题
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 发行时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // 过期时间（24小时后）
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func NewJWTService() *JWTService {
	return &JWTService{}
}
