package tcp

import (
	"context"
	"github.com/DaiYuANg/maxio/pkg/gnet_zap_logger"
	"github.com/panjf2000/gnet/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func startServer(lc fx.Lifecycle, server *server, logger *zap.SugaredLogger) {
	myLogger := gnet_zap_logger.NewZapLogger(logger)
	options := gnet.Options{
		Multicore: true,
		Logger:    myLogger,
		ReusePort: true,
		ReuseAddr: true,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				err := gnet.Run(
					server, server.addr, gnet.WithOptions(options),
				)
				if err != nil {
					panic(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Stop(ctx)
		},
	})
}
