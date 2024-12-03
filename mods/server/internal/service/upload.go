package service

import (
	"context"
	"github.com/codingsince1985/checksum"
	"github.com/gabriel-vasile/mimetype"
	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"
	"mime/multipart"
	"nexus/internal/constant"
	"nexus/internal/entity"
	"nexus/internal/repository"
)

type Upload struct {
	*minio.Client
	*zap.Logger
	*repository.FileResourceRepository
}

func (u *Upload) UploadFile(ctx context.Context, file *multipart.FileHeader) error {
	fileSource, err := file.Open()
	if err != nil {
		return nil
	}
	defer func(fileSource multipart.File) {
		err := fileSource.Close()
		if err != nil {
			u.Error("", zap.Error(err))
		}
	}(fileSource)
	mtype, _ := mimetype.DetectReader(fileSource)
	fileMd5, err := checksum.MD5sumReader(fileSource)
	if err != nil {
		return err
	}
	u.Info("File", zap.String("md5", fileMd5), zap.String("type", mtype.String()))
	go u.PutObject(ctx, constant.Major, file.Filename, fileSource, file.Size, minio.PutObjectOptions{})

	var fileResource = entity.FileResource{
		Md5:    fileMd5,
		Bucket: constant.Major,
		Object: file.Filename,
		Mime:   mtype.String(),
		Size:   file.Size,
	}
	go u.Create(&fileResource)
	return nil
}
