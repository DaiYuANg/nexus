package entry

import (
	"github.com/gofiber/fiber/v2"
)

type UploadController struct {
}

func (u UploadController) RegisterRoutes(app *fiber.App) {
	app.Post("/upload", u.Upload)
}

func (u UploadController) Upload(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	// 保存到指定路径
	err = c.SaveFile(file, "./uploads/"+file.Filename)
	if err != nil {
		return err
	}

	return c.SendString("upload ok: " + file.Filename)
}

func newUploadController() *UploadController {
	return &UploadController{}
}
