package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func registerRequestId(app *fiber.App) {
	app.Use(requestid.New())
}
