package server

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"nexus/internal/config"
)

type LifecycleParam struct {
	fx.In
	fx.Lifecycle
	*fiber.App
	*config.HttpConfig
	*zap.Logger
}

func httpLifecycle(param LifecycleParam) {
	param.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				lo.Must0(param.Listen(param.GetPort()))
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return param.ShutdownWithContext(ctx)
		},
	})
}
