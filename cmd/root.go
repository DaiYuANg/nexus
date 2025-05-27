package cmd

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"storix/module/config"
	"storix/module/entry"
	"storix/module/indexer"
	"storix/module/logger"
	"storix/module/metadata"
	"storix/module/storage"
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
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	).Run()
}
