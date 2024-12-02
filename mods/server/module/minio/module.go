package minio

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"nexus/internal/constant"
	minio2 "nexus/internal/minio"
	"nexus/internal/model"
)

var Module = fx.Module("minio", fx.Provide(newMinioClient, minio2.NewWrapper), fx.Invoke(checkMajorBucket))

type Param struct {
	fx.In
	MinioConfig *model.MinioConfig
	Logger      *zap.Logger
}

func newMinioClient(param Param) (*minio.Client, error) {
	config, logger := param.MinioConfig, param.Logger
	logger.Info("Minio", zap.Any("Config", config))
	return minio.New(param.MinioConfig.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKey, config.SecretKey, ""),
		Secure: config.UseSSL,
	})
}

func checkMajorBucket(client *minio2.Wrapper) error {
	return client.CreateBucket(constant.Major)
}
