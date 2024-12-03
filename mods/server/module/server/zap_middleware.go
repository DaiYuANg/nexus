package server

import (
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func registerZap(app *fiber.App, log *zap.Logger) {
	app.Use(fiberzap.New(fiberzap.Config{
		Logger: log,
	}))
}
