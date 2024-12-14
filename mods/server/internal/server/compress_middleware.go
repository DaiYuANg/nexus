package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

func registerCompress(app *fiber.App) {
	app.Use(compress.New(
		compress.Config{
			Level: compress.LevelBestSpeed,
		}),
	)
}
