package service

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	"nexus/internal/repository"
	"nexus/vfs/local"
)

type UploadParam struct {
	fx.In
	*local.VFS
	//*minio.Client
	*zap.Logger
	*repository.FileResourceRepository
}

func NewUpload(param UploadParam) *Upload {
	return &Upload{
		vfs:                    param.VFS,
		Logger:                 param.Logger,
		FileResourceRepository: param.FileResourceRepository,
	}
}
