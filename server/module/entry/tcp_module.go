package entry

import (
	"context"
	"github.com/panjf2000/gnet/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var TcpModule = fx.Module("tcp", fx.Provide(newStorixServer), fx.Invoke(startStorixServer))

type storixServer struct {
	*gnet.BuiltinEventEngine
}

func (es *storixServer) OnTraffic(c gnet.Conn) gnet.Action {
	buf, _ := c.Next(-1)
	_ = c.AsyncWrite(buf, func(c gnet.Conn, err error) error {

		return nil
	})
	return gnet.None
}

func newStorixServer() *storixServer {
	return &storixServer{}
}

type ZapLogger struct {
	logger *zap.SugaredLogger
}

func (l *ZapLogger) Fatalf(format string, args ...any) {
	l.logger.Fatalf(format, args...)
}

func (l *ZapLogger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}
func (l *ZapLogger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}
func (l *ZapLogger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}
func (l *ZapLogger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func startStorixServer(lc fx.Lifecycle, server *storixServer, logger *zap.Logger) {
	myLogger := &ZapLogger{
		logger: logger.Sugar(),
	}
	options := gnet.Options{
		Multicore: true,
		Logger:    myLogger,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				err := gnet.Run(server, "tcp://:8080", gnet.WithOptions(options))
				if err != nil {
					panic(err)
				}
			}()
			return nil
		},
	})
}
