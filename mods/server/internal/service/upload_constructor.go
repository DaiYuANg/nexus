package service

import (
	"github.com/minio/minio-go/v7"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"nexus/internal/repository"
)

type UploadParam struct {
	fx.In
	*minio.Client
	*zap.Logger
	*repository.FileResourceRepository
}

func NewUpload(param UploadParam) *Upload {
	return &Upload{
		Client:                 param.Client,
		Logger:                 param.Logger,
		FileResourceRepository: param.FileResourceRepository,
	}
}
