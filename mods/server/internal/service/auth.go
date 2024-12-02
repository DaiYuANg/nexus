package service

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	"nexus/internal/model"
)

type Auth struct {
	*JWT
	*zap.Logger
}

func (s *Auth) Login(loginUser model.LoginUser) (string, error) {
	sign, err := s.Sign("test")
	if err != nil {
		return "", err
	}
	s.Info("sign", zap.String("sign", sign))
	return sign, nil
}

type AuthParam struct {
	fx.In
	Jwt    *JWT
	Logger *zap.Logger
}

func NewAuth(param AuthParam) *Auth {
	return &Auth{
		JWT:    param.Jwt,
		Logger: param.Logger,
	}
}
