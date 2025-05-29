package tcp

import (
	"context"
	"github.com/panjf2000/gnet/v2"
	"go.uber.org/zap"
)

type server struct {
	*gnet.BuiltinEventEngine
	engine    gnet.Engine
	addr      string
	multicore bool
	*zap.SugaredLogger
}

func (s *server) OnBoot(eng gnet.Engine) gnet.Action {
	s.engine = eng
	s.Info("echo server with multi-core=%t is listening on %s\n", s.multicore, s.addr)
	return gnet.None
}

func (s *server) OnTraffic(c gnet.Conn) gnet.Action {
	state := c.Context().(*ConnState)
	data, _ := c.Next(-1)
	state.Buffer = append(state.Buffer, data...)

	for {
		msg, remaining, ok := decodeMessage(state.Buffer)
		if !ok {
			break
		}
		state.Buffer = remaining

		if msg.Version != ProtocolVersion {
			return gnet.Close
		}

		if act := s.handleMessage(c, state, msg); act != nil {
			return *act
		}
	}
	return gnet.None
}

func (s *server) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	c.SetContext(&ConnState{
		Authenticated: false,
		FileOffset:    0,
		Buffer:        make([]byte, 0),
	})
	return
}

func (s *server) OnClose(c gnet.Conn, err error) (action gnet.Action) {
	//log.Printf("Connection closed: %s, err: %v", c.RemoteAddr().String(), err)
	return
}

func newStorixServer(logger *zap.SugaredLogger) *server {
	return &server{
		multicore:     true,
		addr:          "tcp://:8080",
		SugaredLogger: logger,
	}
}

func (s *server) Stop(ctx context.Context) error {
	return s.engine.Stop(ctx)
}
