package entry

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"go.uber.org/fx"
)

var HttpModule = fx.Module("http",
	fx.Provide(
		newFiber,
	),
	controllerModule,
	fx.Invoke(startHttp),
)

func newFiber() *fiber.App {
	app := fiber.New(
		fiber.Config{
			EnablePrintRoutes: true,
		},
	)
	app.Get("/metrics", monitor.New())
	return app
}

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
	})
}
