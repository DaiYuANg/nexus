package minio

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"nexus/internal/constant"
	"nexus/internal/model"
)

var Module = fx.Module("minio", fx.Provide(newMinioClient), fx.Invoke(checkMajorBucket))

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

func checkMajorBucket(client *minio.Client) error {
	ctx := context.Background()
	exists, err := client.BucketExists(ctx, constant.Major)
	if err != nil {
		return err
	}
	if !exists {
		err := client.MakeBucket(ctx, constant.Major, minio.MakeBucketOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}
