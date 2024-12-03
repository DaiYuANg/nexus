package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/pprof"
)

func registerPprof(app *fiber.App) {
	app.Use(pprof.New())
}
