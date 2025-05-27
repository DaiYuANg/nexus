package tcp

import (
	"context"
	"github.com/panjf2000/gnet/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func startStorixServer(lc fx.Lifecycle, server *storixServer, logger *zap.SugaredLogger) {
	myLogger := newZapLogger(logger)
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
