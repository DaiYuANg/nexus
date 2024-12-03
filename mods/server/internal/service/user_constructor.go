package service

import (
	goeventbus "github.com/stanipetrosyan/go-eventbus"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"nexus/internal/repository"
)

type UserServiceParam struct {
	fx.In
	goeventbus.EventBus
	*gorm.DB
	*zap.Logger
	*repository.UserRepository
}

func NewUserService(param UserServiceParam) *User {
	return &User{
		db:             param.DB,
		Logger:         param.Logger,
		EventBus:       param.EventBus,
		UserRepository: param.UserRepository,
	}
}
