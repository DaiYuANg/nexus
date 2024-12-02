package service

import (
	"go.uber.org/fx"
	"nexus/internal/service"
)

var Module = fx.Module("service",
	fx.Provide(
		service.NewJWTService,
		service.NewUserService,
		service.NewUpload,
		service.NewAuth,
	),
)
