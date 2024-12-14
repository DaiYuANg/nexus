package schedule

import (
	"github.com/go-co-op/gocron/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"nexus/vfs/local"
	"time"
)

type newScanParams struct {
	fx.In
	*zap.Logger
	gocron.Scheduler
	*local.VFS
	//*fs.Wrapper
}

func Scan(params newScanParams) error {
	_, err := params.NewJob(
		gocron.DurationJob(
			60*time.Second,
		),
		gocron.NewTask(
			func() {
				//bucket, err := params.ListBucket()
				params.Info("test")
				//if err != nil {
				//	params.Error("failed to list bucket", zap.Error(err))
				//	return
				//}
				//lo.ForEach(bucket, func(item string, index int) {
				//	params.Info("scan bucket", zap.String("bucket", item))
				//})
			},
		),
	)
	return err
}
