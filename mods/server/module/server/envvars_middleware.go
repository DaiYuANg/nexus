package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/envvar"
)

func registerEnvvars(app *fiber.App) {
	app.Use("/expose/envvars", envvar.New())
}
