package minio

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"nexus/internal/conf"
)

type Param struct {
	fx.In
	*conf.MinioConfig
	*zap.Logger
}

func newMinioClient(param Param) (*minio.Client, error) {
	param.Info("Minio", zap.Any("Config", param.MinioConfig))
	return minio.New(param.MinioConfig.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(param.AccessKey, param.SecretKey, ""),
		Secure: param.UseSSL,
	})
}
