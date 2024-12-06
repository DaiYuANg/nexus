package fs

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type Wrapper struct {
	*minio.Client
	*zap.Logger
	context.Context
}

func (w *Wrapper) CreateBucket(bucket string) error {
	exists, err := w.BucketExists(w, bucket)
	if err != nil {
		return err
	}
	if !exists {
		err := w.MakeBucket(w, bucket, minio.MakeBucketOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}

func (w *Wrapper) ListBucket() ([]string, error) {
	buckets, err := w.ListBuckets(w)
	if err != nil {
		return nil, err
	}
	result := lo.Map(buckets, func(item minio.BucketInfo, index int) string {
		return item.Name
	})
	return result, nil
}
