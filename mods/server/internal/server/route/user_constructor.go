package route

import (
	"go.uber.org/zap"
	"nexus/internal/service"
)

func NewUserRoute(userService *service.User, logger *zap.Logger, auth *service.Auth) *User {
	return &User{
		User:   userService,
		Logger: logger,
		Auth:   auth,
	}
}
