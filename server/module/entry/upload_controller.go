package entry

import (
	"github.com/gofiber/fiber/v2"
)

type uploadController struct {
}

func (u uploadController) RegisterRoutes(app *fiber.App) {
	app.Post("/upload/:namespace", u.Upload)
}

func (u uploadController) Upload(c *fiber.Ctx) error {
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

func newUploadController() *uploadController {
	return &uploadController{}
}
