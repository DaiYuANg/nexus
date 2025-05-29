package http

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func startHttp(lc fx.Lifecycle, app *fiber.App) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				err := app.Listen(":3000")
				if err != nil {
					panic(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.ShutdownWithContext(ctx)
		},
	})
}
