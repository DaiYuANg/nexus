package controller

import (
	"github.com/DaiYuANg/storix/server/internal/namespace"
	"github.com/DaiYuANg/storix/server/internal/storage"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type UploadController struct {
	*zap.SugaredLogger
	*namespace.Service
	*storage.FileChunker
}

func (u UploadController) RegisterRoutes(app *fiber.App) {
	app.Post("/upload/:namespace", u.Upload)
}

func (u UploadController) Upload(c *fiber.Ctx) error {
	ns := c.Params("namespace")
	u.Debugf("namespace: %s", ns)
	if u.Service.NamespaceExists(ns) == false {
		return Fail(c, fiber.StatusNotFound, "namespace not found")
	}
	fileHeader, err := c.FormFile("file")
	if err != nil {
		u.Errorf("failed to upload file: %v", err)
		return Fail(c, fiber.StatusInternalServerError, "failed to upload file")
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
	service *namespace.Service,
	fileChunker *storage.FileChunker,
) *UploadController {
	return &UploadController{
		logger,
		service,
		fileChunker,
	}
}
