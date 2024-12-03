package service

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	"nexus/internal/repository"
)

type AuthParam struct {
	fx.In
	*JWT
	*zap.Logger
	*repository.UserRepository
}

func NewAuth(param AuthParam) *Auth {
	return &Auth{
		JWT:            param.JWT,
		Logger:         param.Logger,
		UserRepository: param.UserRepository,
	}
}
