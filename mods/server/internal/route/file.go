package route

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"nexus/internal/model"
)

type File struct {
	*zap.Logger
}

func (f *File) List(ctx *fiber.Ctx) error {
	f.Info("User", zap.Any("u", ctx.Locals("user")))
	return ctx.Status(200).JSON(model.JustOk())
}

func (f *File) Routes() []Info {
	return []Info{
		{
			Path:    "/file/list",
			Method:  GET,
			Handler: f.List,
		},
	}
}
