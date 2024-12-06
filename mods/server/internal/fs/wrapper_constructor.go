package fs

import (
	"context"
	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"
)

func NewWrapper(client *minio.Client, logger *zap.Logger) *Wrapper {
	return &Wrapper{Client: client, Logger: logger, Context: context.Background()}
}
