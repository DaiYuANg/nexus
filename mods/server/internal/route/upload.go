package route

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"nexus/internal/service"
)

type Upload struct {
	*zap.Logger
	*service.Upload
}

func (u *Upload) SingleFile(c *fiber.Ctx) error {
	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		u.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "File upload failed",
		})
	}
	u.UploadFile(c.Context(), file)
	return c.JSON(fiber.Map{
		"message":  "File uploaded successfully",
		"filename": file.Filename,
	})
}

func (u *Upload) Routes() []Info {
	return []Info{
		{
			Path:    "/upload",
			Method:  fiber.MethodPost,
			Handler: u.SingleFile,
		},
	}
}

func NewUpload(
	logger *zap.Logger,
	uploadService *service.Upload,
) *Upload {
	return &Upload{
		Logger: logger,
		Upload: uploadService,
	}
}
