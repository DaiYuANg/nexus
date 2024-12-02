package service

import (
	"context"
	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"
	"mime/multipart"
	"nexus/internal/constant"
)

type Upload struct {
	client *minio.Client
	logger *zap.Logger
}

func (u Upload) UploadFile(ctx context.Context, file *multipart.FileHeader) error {
	fileSource, err := file.Open()
	if err != nil {
		return nil
	}
	defer fileSource.Close()
	//go io.CalculateMD5File(fileSource)
	//md5File, err := io.CalculateMD5File(fileSource)
	//if err != nil {
	//	return err
	//}
	u.client.PutObject(ctx, constant.Major, file.Filename, fileSource, file.Size, minio.PutObjectOptions{})
	return nil
}

func NewUpload(client *minio.Client, logger *zap.Logger) *Upload {
	return &Upload{
		client: client,
		logger: logger,
	}
}
