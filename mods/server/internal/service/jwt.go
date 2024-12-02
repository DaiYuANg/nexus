package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"time"
)

type JWT struct {
	logger     *zap.Logger
	signingKey []byte
}
type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func (s *JWT) Sign(userID string) (string, error) {
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

	signedToken, err := token.SignedString(s.signingKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (s *JWT) Parse(tokenString string) (*jwt.Token, error) {
	// 解析并验证 JWT
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法是否匹配
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.signingKey, nil
	})
	if err != nil {
		// 如果 token 解析失败，返回错误
		return nil, err
	}
	return token, nil
}

type JwtServiceParam struct {
	fx.In
	SigningKey []byte `name:"jwtKey"`
	Logger     *zap.Logger
}

func NewJWTService(param JwtServiceParam) *JWT {
	return &JWT{
		logger:     param.Logger,
		signingKey: param.SigningKey,
	}
}
