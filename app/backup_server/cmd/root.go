package cmd

import (
	"github.com/DaiYuANg/maxio/backup-server/internal/logger"
	"github.com/DaiYuANg/maxio/backup-server/internal/server"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func Run() {
	fx.New(
		logger.Module,
		server.Module,
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		})).
		Run()
}
