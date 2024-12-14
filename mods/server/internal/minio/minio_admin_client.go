package minio

import (
	"github.com/minio/madmin-go/v3"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"nexus/internal/config"
)

type AdminParam struct {
	fx.In
	*config.MinioConfig
	*zap.Logger
}

func newMinioAdminClient(param AdminParam) (*madmin.AdminClient, error) {
	return madmin.New(param.Endpoint, param.AccessKey, param.SecretKey, param.UseSSL)
}
