package cmd

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"nexus/internal/cache"
	"nexus/internal/config"
	"nexus/internal/db"
	"nexus/internal/email"
	"nexus/internal/event"
	"nexus/internal/fs"
	"nexus/internal/i18n"
	"nexus/internal/jwt"
	"nexus/internal/logger"
	"nexus/internal/plugin"
	"nexus/internal/print"
	"nexus/internal/repository"
	"nexus/internal/runtime"
	"nexus/internal/schedule"
	"nexus/internal/server"
	"nexus/internal/service"
)

func newDiContainer() *fx.App {
	return fx.New(
		config.Module,
		runtime.Module,
		logger.Module,
		fs.Module,
		server.HttpModule,
		server.RouteModule,
		jwt.Module,
		repository.Module,
		service.Module,
		schedule.Module,
		event.Module,
		plugin.Module,
		email.Module,
		cache.Module,
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		db.Module,
		//minio.Module,
		i18n.Module,
		print.Module,
	)
}
