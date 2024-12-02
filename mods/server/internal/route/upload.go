package route

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"
)

type Upload struct {
	client *minio.Client
	logger *zap.Logger
}

func (u Upload) SingleFile(c *fiber.Ctx) error {
	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "File upload failed",
		})
	}
	// 打印文件名
	fmt.Println("Received file:", file.Filename)
	// 返回成功的响应
	return c.JSON(fiber.Map{
		"message":  "File uploaded successfully",
		"filename": file.Filename,
	})
}

func (Upload) Routes() []Info {
	return []Info{}
}

func NewUpload(client *minio.Client) Upload {
	return Upload{client: client}
}
