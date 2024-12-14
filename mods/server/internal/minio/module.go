package minio

import (
	"go.uber.org/fx"
	"nexus/internal/constant"
	"nexus/internal/fs"
	"nexus/vfs/s3"
)

var Module = fx.Module("minio",
	fx.Provide(
		newMinioClient,
		newMinioAdminClient,
		fs.NewWrapper,
	),
	fx.Invoke(checkMajorBucket),
)

func checkMajorBucket(client *fs.Wrapper) error {
	s3.New(s3.VfsConfig{
		Logger:    nil,
		Endpoint:  "",
		AccessKey: "",
		SecretKey: "",
		UseSSL:    false,
	})
	return client.CreateBucket(constant.Major)
}
