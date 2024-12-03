package minio

import (
	"github.com/minio/madmin-go/v3"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"nexus/internal/conf"
)

type AdminParam struct {
	fx.In
	*conf.MinioConfig
	*zap.Logger
}

func newMinioAdminClient(param AdminParam) (*madmin.AdminClient, error) {
	return madmin.New(param.Endpoint, param.AccessKey, param.SecretKey, param.UseSSL)
}
