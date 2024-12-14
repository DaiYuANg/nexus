package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func registerLogger(app *fiber.App) {
	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format:        "${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
		DisableColors: false,
	}))
}
