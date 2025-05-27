package cmd

import (
	"github.com/DaiYuANg/storix/server/module/auth"
	"github.com/DaiYuANg/storix/server/module/config"
	"github.com/DaiYuANg/storix/server/module/entry"
	"github.com/DaiYuANg/storix/server/module/indexer"
	"github.com/DaiYuANg/storix/server/module/logger"
	"github.com/DaiYuANg/storix/server/module/metadata"
	"github.com/DaiYuANg/storix/server/module/namespace"
	"github.com/DaiYuANg/storix/server/module/schedule"
	"github.com/DaiYuANg/storix/server/module/storage"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func Run() {
	fx.New(
		config.Module,
		logger.Module,
		entry.HttpModule,
		entry.TcpModule,
		metadata.Module,
		indexer.Module,
		storage.Module,
		namespace.Module,
		schedule.Module,
		auth.Module,
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	).Run()
}
