package tcp

import (
	"context"
	"github.com/panjf2000/gnet/v2"
	"go.uber.org/zap"
)

type storixServer struct {
	*gnet.BuiltinEventEngine
	engine    gnet.Engine
	addr      string
	multicore bool
	*zap.SugaredLogger
}

func (es *storixServer) OnBoot(eng gnet.Engine) gnet.Action {
	es.engine = eng
	es.Info("echo server with multi-core=%t is listening on %s\n", es.multicore, es.addr)
	return gnet.None
}

func (es *storixServer) OnTraffic(c gnet.Conn) gnet.Action {
	buf, _ := c.Next(-1)
	_ = c.AsyncWrite(buf, func(c gnet.Conn, err error) error {

		return nil
	})
	return gnet.None
}

func newStorixServer(logger *zap.SugaredLogger) *storixServer {
	return &storixServer{
		multicore:     true,
		addr:          "tcp://:8080",
		SugaredLogger: logger,
	}
}

func (es *storixServer) Stop(ctx context.Context) error {
	return es.engine.Stop(ctx)
}
