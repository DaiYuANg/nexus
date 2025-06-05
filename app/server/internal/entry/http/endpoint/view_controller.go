package endpoint

import (
	"github.com/gofiber/fiber/v2"
)

type ViewController struct {
}

func (v ViewController) RegisterRoutes(app *fiber.App) {
	app.Get("/", v.index)
}

func (v ViewController) index(ctx *fiber.Ctx) error {
	return ctx.Render("index", fiber.Map{})
}

func NewViewController() *ViewController {
	return &ViewController{}
}
