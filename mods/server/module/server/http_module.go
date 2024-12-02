package server

import (
	"context"
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/envvar"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go.uber.org/fx"
	"nexus/internal/model"
)

var HttpModule = fx.Module("http", fx.Provide(newHttpServer), fx.Invoke(httpLifecycle))

func newHttpServer() *fiber.App {
	app := fiber.New(fiber.Config{EnablePrintRoutes: true, Immutable: true})
	app.Use(compress.New())
	app.Use("/expose/envvars", envvar.New())
	app.Get("/metrics", monitor.New())
	app.Use(requestid.New())
	app.Use(pprof.New())
	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format:        "${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
		DisableColors: false,
	}))
	app.Use(recover.New())
	prometheus := fiberprometheus.New("my-service-name")
	prometheus.RegisterAt(app, "/metrics")
	prometheus.SetSkipPaths([]string{"/ping"}) // Optional: Remove some paths from metrics
	app.Use(prometheus.Middleware)

	return app
}

type LifecycleParam struct {
	fx.In
	Lc         fx.Lifecycle
	App        *fiber.App
	HttpConfig *model.HttpConfig
}

func httpLifecycle(param LifecycleParam) {
	lc, app, config := param.Lc, param.App, param.HttpConfig
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return app.Listen(config.GetPort())
		},
	})
}
