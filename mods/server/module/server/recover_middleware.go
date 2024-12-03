package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func registerRecover(app *fiber.App) {
	app.Use(recover.New())

}
