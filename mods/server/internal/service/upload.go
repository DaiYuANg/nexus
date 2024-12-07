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
	//mtype, _ := mimetype.DetectReader(fileSource)
	//fileMd5, err := checksum.MD5sumReader(fileSource)
	//if err != nil {
	//	return err
	//}
	//u.Info("File", zap.String("md5", fileMd5), zap.String("type", mtype.String()))
	//go u.PutObject(ctx, constant.Major, file.Filename, fileSource, file.Size, minio.PutObjectOptions{})
	all, err := io.ReadAll(fileSource)
	if err != nil {
		return err
	}
	err = u.vfs.WriteFile("/test/"+file.Filename, all)
	if err != nil {
		return err
	}
	//var fileResource = entity.FileResource{
	//	Md5:    fileMd5,
	//	Bucket: constant.Major,
	//	Object: file.Filename,
	//	Mime:   mtype.String(),
	//	Size:   file.Size,
	//}
	defer func(fileSource multipart.File) {
		err := fileSource.Close()
		if err != nil {
			u.Error("", zap.Error(err))
		}
	}(fileSource)
	//go u.Create(&fileResource)
	return nil
}
