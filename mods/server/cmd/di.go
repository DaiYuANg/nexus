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
	"nexus/module/jwt"
	"nexus/module/logger"
	"nexus/module/minio"
	"nexus/module/plugin"
	print2 "nexus/module/print"
	"nexus/module/repository"
	"nexus/module/schedule"
	"nexus/module/server"
	"nexus/module/service"
)

func newDiContainer() *fx.App {
	return fx.New(
		config.Module,
		logger.Module,
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
		},
		),
		db.Module,
		minio.Module,
		print2.Module,
	)
}
