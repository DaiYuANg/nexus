package minio

import (
	"go.uber.org/fx"
	"nexus/internal/constant"
	minio2 "nexus/internal/minio"
)

var Module = fx.Module("minio",
	fx.Provide(
		newMinioClient,
		newMinioAdminClient,
		minio2.NewWrapper,
	),
	fx.Invoke(checkMajorBucket),
)

func checkMajorBucket(client *minio2.Wrapper) error {
	return client.CreateBucket(constant.Major)
}
