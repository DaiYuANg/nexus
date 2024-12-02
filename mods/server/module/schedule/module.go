package schedule

import (
	"context"
	"github.com/go-co-op/gocron/v2"
	"github.com/samber/lo"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("schedule", fx.Provide(newScheduler), fx.Invoke(startScheduler))

func newScheduler() gocron.Scheduler {
	return lo.Must(
		gocron.NewScheduler(
			gocron.WithLogger(
				gocron.NewLogger(gocron.LogLevelInfo),
			),
		),
	)
}

func startScheduler(lc fx.Lifecycle, cron gocron.Scheduler, log *zap.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go cron.Start()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return cron.Shutdown()
		},
	})
}
