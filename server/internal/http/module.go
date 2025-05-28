package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"go.uber.org/fx"
)

var Module = fx.Module("http",
	fx.Provide(
		newFiber,
	),
	middlewareModule,
	controllerModule,
	fx.Invoke(startHttp),
)

func newFiber() *fiber.App {
	app := fiber.New(
		fiber.Config{
			EnablePrintRoutes: true,
			BodyLimit:         1024 * 1024 * 1024,
		},
	)
	app.Get("/metrics", monitor.New())
	return app
}
