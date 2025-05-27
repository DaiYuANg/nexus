package http

import (
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/envvar"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var middlewareModule = fx.Module("middleware",
	fx.Invoke(
		loggingMiddleware,
		requestIdMiddleware,
		pprofMiddleware,
		envvarMiddleware,
		compressMiddleware,
	),
)

func loggingMiddleware(app *fiber.App, zapLogger *zap.Logger) {
	app.Use(
		fiberzap.New(
			fiberzap.Config{
				Logger: zapLogger,
			},
		),
	)
}

func requestIdMiddleware(app *fiber.App) {
	app.Use(requestid.New())
}

func pprofMiddleware(app *fiber.App) {
	app.Use(pprof.New())
}

func envvarMiddleware(app *fiber.App) {
	app.Use("/expose/envvars", envvar.New())
}

func compressMiddleware(app *fiber.App) {
	app.Use(compress.New())
}
