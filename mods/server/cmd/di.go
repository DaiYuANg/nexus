package cmd

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"nexus/module/cache"
	"nexus/module/config"
	"nexus/module/db"
	"nexus/module/email"
	"nexus/module/event"
	"nexus/module/fs"
	"nexus/module/i18n"
	"nexus/module/jwt"
	"nexus/module/logger"
	"nexus/module/plugin"
	"nexus/module/print"
	"nexus/module/repository"
	"nexus/module/runtime"
	"nexus/module/schedule"
	"nexus/module/server"
	"nexus/module/service"
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
