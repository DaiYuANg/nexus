package schedule

import (
	"context"
	"github.com/go-co-op/gocron/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"time"
)

var Module = fx.Module("scheduler",
	fx.Provide(
		newScheduler,
	),
	fx.Invoke(
		startScheduler,
	),
)

func newScheduler(logger *zap.Logger) (gocron.Scheduler, error) {
	zapLogger := schedulerZapLogger{zapLogger: logger.Sugar()}
	return gocron.NewScheduler(
		gocron.WithLogger(zapLogger),
		gocron.WithStopTimeout(time.Second*5),
	)
}

func startScheduler(lc fx.Lifecycle, s gocron.Scheduler) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go s.Start()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return s.Shutdown()
		},
	})
}
