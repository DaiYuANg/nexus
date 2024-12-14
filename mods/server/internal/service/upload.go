package service

import (
	"context"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"nexus/internal/repository"
	"nexus/vfs/local"
)

type Upload struct {
	vfs *local.VFS
	*zap.Logger
	*repository.FileResourceRepository
}

func (u *Upload) UploadFile(ctx context.Context, file *multipart.FileHeader) error {
	fileSource, err := file.Open()
	if err != nil {
		return nil
	}

	user := ctx.Value("user")
	u.Info("user", zap.Any("user", user))
	all, err := io.ReadAll(fileSource)
	if err != nil {
		return err
	}

	err = u.vfs.WriteFile("/test/"+file.Filename, all)
	if err != nil {
		return err
	}

	defer func(fileSource multipart.File) {
		err := fileSource.Close()
		if err != nil {
			u.Error("", zap.Error(err))
		}
	}(fileSource)
	//go u.Create(&fileResource)
	return nil
}
