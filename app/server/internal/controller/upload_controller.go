package controller

import (
	"fmt"
	"github.com/DaiYuANg/maxio/server/internal/bucket"
	"github.com/DaiYuANg/maxio/server/internal/storage"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type UploadController struct {
	*zap.SugaredLogger
	*bucket.Service
	*storage.FileChunker
}

func (u UploadController) RegisterRoutes(app *fiber.App) {
	app.Post("/upload/:bucket", u.Upload)
}

func (u UploadController) Upload(c *fiber.Ctx) error {
	ns := c.Params("bucket")
	u.Debugf("bucket: %s", ns)
	if u.Service.NamespaceExists(ns) == false {
		return Fail(c, fiber.StatusNotFound, "bucket not found")
	}
	fileHeader, err := c.FormFile("file")
	if err != nil {
		u.Errorf("failed to upload file: %v", err)
		return Fail(c, fiber.StatusInternalServerError, fmt.Sprintf("failed to upload file:%s", err.Error()))
	}

	part, err := u.ChunkSave(fileHeader, ns, u.SugaredLogger)
	if err != nil {
		u.Errorf("failed to upload file: %v", err)
		return Fail(c, fiber.StatusInternalServerError, "failed to upload file")
	}

	return Success(c, fiber.Map{
		"message":    "upload success",
		"chunkCount": part,
	})
}

func NewUploadController(
	logger *zap.SugaredLogger,
	service *bucket.Service,
	fileChunker *storage.FileChunker,
) *UploadController {
	return &UploadController{
		logger,
		service,
		fileChunker,
	}
}
