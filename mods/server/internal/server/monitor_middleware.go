package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func registerMonitor(app *fiber.App) {
	app.Get("/metrics", monitor.New())
}
