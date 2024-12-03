package server

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"nexus/internal/conf"
)

type LifecycleParam struct {
	fx.In
	fx.Lifecycle
	*fiber.App
	*conf.HttpConfig
	*zap.Logger
}

func httpLifecycle(param LifecycleParam) {
	param.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				err := param.Listen(param.GetPort())
				if err != nil {
					param.Error("Start http server failed", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return param.ShutdownWithContext(ctx)
		},
	})
}
