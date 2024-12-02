package minio

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"nexus/internal/model"
)

var Module = fx.Module("minio", fx.Provide(newMinioClient))

type Param struct {
	fx.In
	MinioConfig *model.MinioConfig
	Logger      *zap.Logger
}

func newMinioClient(param Param) (*minio.Client, error) {
	config, logger := param.MinioConfig, param.Logger
	logger.Info("Minio", zap.Any("Config", config))
	// Initialize minio client object.
	return minio.New(param.MinioConfig.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKey, config.SecretKey, ""),
		Secure: config.UseSSL,
	})
}
