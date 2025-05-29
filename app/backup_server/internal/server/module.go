package server

import (
	"context"
	"fmt"
	"github.com/panjf2000/gnet/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type BackupServer struct {
	gnet.BuiltinEventEngine
	engine    gnet.Engine
	addr      string
	multicore bool
	logger    *zap.SugaredLogger
}

func (es *BackupServer) OnBoot(eng gnet.Engine) gnet.Action {
	es.engine = eng
	es.logger.Infof("echo server with multi-core=%t is listening on %s\n", es.multicore, es.addr)
	return gnet.None
}

func (es *BackupServer) OnTraffic(c gnet.Conn) gnet.Action {
	buf, _ := c.Next(-1)
	_, err := c.Write(buf)
	if err != nil {
		return 0
	}
	return gnet.None
}

var Module = fx.Module("server", fx.Provide(newBackupServer), fx.Invoke(startBackupServer))

func newBackupServer(logger *zap.Logger) *BackupServer {
	port := 8976
	return &BackupServer{
		multicore: true,
		addr:      fmt.Sprintf("tcp://:%d", port),
		logger:    logger.Sugar(),
	}
}

func startBackupServer(lc fx.Lifecycle, backupServer *BackupServer, logger *zap.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				err := gnet.Run(backupServer, backupServer.addr, gnet.WithMulticore(backupServer.multicore))
				if err != nil {
					panic(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return backupServer.engine.Stop(ctx)
		},
	})
}
