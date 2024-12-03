package route

import (
	"go.uber.org/zap"
	"nexus/internal/service"
)

func NewUpload(
	logger *zap.Logger,
	uploadService *service.Upload,
) *Upload {
	return &Upload{
		Logger: logger,
		Upload: uploadService,
	}
}
