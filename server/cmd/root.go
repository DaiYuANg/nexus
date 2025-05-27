package cmd

import (
	"github.com/DaiYuANg/storix/server/internal/auth"
	"github.com/DaiYuANg/storix/server/internal/config"
	"github.com/DaiYuANg/storix/server/internal/http"
	"github.com/DaiYuANg/storix/server/internal/indexer"
	"github.com/DaiYuANg/storix/server/internal/internal_store"
	"github.com/DaiYuANg/storix/server/internal/logger"
	"github.com/DaiYuANg/storix/server/internal/namespace"
	"github.com/DaiYuANg/storix/server/internal/schedule"
	"github.com/DaiYuANg/storix/server/internal/storage"
	"github.com/DaiYuANg/storix/server/internal/tcp"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func Run() {
	fx.New(
		config.Module,
		logger.Module,
		http.Module,
		tcp.Module,
		internal_store.Module,
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
