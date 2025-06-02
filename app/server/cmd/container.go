package cmd

import (
	"github.com/DaiYuANg/maxio/server/internal/auth"
	"github.com/DaiYuANg/maxio/server/internal/bucket"
	"github.com/DaiYuANg/maxio/server/internal/config"
	"github.com/DaiYuANg/maxio/server/internal/indexer"
	"github.com/DaiYuANg/maxio/server/internal/internal_store"
	"github.com/DaiYuANg/maxio/server/internal/logger"
	"github.com/DaiYuANg/maxio/server/internal/protocol"
	"github.com/DaiYuANg/maxio/server/internal/schedule"
	"github.com/DaiYuANg/maxio/server/internal/stash"
	"github.com/DaiYuANg/maxio/server/internal/storage"
	"github.com/DaiYuANg/maxio/server/internal/worker"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func container() *fx.App {
	return fx.New(
		config.Module,
		logger.Module,
		internal_store.Module,
		stash.Module,
		protocol.Module,
		indexer.Module,
		storage.Module,
		bucket.Module,
		schedule.Module,
		auth.Module,
		worker.Module,
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	)
}
