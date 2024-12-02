package schedule_task

import (
	"github.com/go-co-op/gocron/v2"
	"github.com/samber/lo"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"nexus/internal/minio"
	"time"
)

type newScanParams struct {
	fx.In
	*zap.Logger
	gocron.Scheduler
	*minio.Wrapper
}

func Scan(params newScanParams) error {
	_, err := params.NewJob(
		gocron.DurationJob(
			10*time.Second,
		),
		gocron.NewTask(
			func() {
				bucket, err := params.ListBucket()
				if err != nil {
					return
				}
				lo.ForEach(bucket, func(item string, index int) {
					params.Info("scan bucket", zap.String("bucket", item))
				})
			},
		),
	)
	return err
}