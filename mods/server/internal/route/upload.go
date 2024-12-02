package route

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"nexus/internal/service"
)

type Upload struct {
	logger        *zap.Logger
	uploadService *service.Upload
}

func (u *Upload) SingleFile(c *fiber.Ctx) error {
	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		u.logger.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "File upload failed",
		})
	}
	u.uploadService.UploadFile(c.Context(), file)
	//// 打印文件名
	//src, err := file.Open()
	//if err != nil {
	//	return c.Status(fiber.StatusInternalServerError).SendString("Failed to open file")
	//}
	//defer src.Close()
	//fmt.Println("Received file:", file.Filename)
	//_, err = u.client.PutObject(c.Context(), constant.Major, file.Filename, src, file.Size, minio.PutObjectOptions{})
	// 返回成功的响应
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
		logger:        logger,
		uploadService: uploadService,
	}
}
