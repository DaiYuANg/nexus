package s3

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
)

type VfsConfig struct {
	*zap.Logger
	Endpoint  string
	AccessKey string
	SecretKey string
	UseSSL    bool
}

func New(config VfsConfig) *VFS {
	client, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKey, config.SecretKey, ""),
		Secure: config.UseSSL,
	})
	if err != nil {
		return nil
	}
	return &VFS{
		client,
		config.Logger,
	}
}
