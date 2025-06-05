package endpoint

import (
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

	if !u.Service.NamespaceExists(ns) {
		return Fail(c, fiber.StatusNotFound, "bucket not found")
	}

	fileHeader, err := c.FormFile("file")
	if err != nil {
		u.Errorf("failed to read file: %v", err)
		return Fail(c, fiber.StatusBadRequest, "invalid file upload")
	}

	chunkCount, err := u.ChunkSave(fileHeader, ns, u.SugaredLogger)
	if err != nil {
		u.Errorf("chunk save failed: %v", err)
		return Fail(c, fiber.StatusInternalServerError, "failed to save file")
	}

	return Success(c, fiber.Map{
		"message":    "upload success",
		"chunkCount": chunkCount,
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
