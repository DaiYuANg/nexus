package service

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	"nexus/internal/model"
)

type Auth struct {
	jwt    *JWT
	logger *zap.Logger
}

func (s *Auth) Login(loginUser model.LoginUser) (string, error) {
	sign, err := s.jwt.Sign("test")
	if err != nil {
		return "", err
	}
	s.logger.Info("sign", zap.String("sign", sign))
	return sign, nil
}

type AuthParam struct {
	fx.In
	Jwt    *JWT
	Logger *zap.Logger
}

func NewAuth(param AuthParam) *Auth {
	return &Auth{
		jwt:    param.Jwt,
		logger: param.Logger,
	}
}
