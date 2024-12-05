package ws

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type ChatWebsocketParam struct {
	fx.In
	*zap.Logger
}

func NewChatWebsocket(param ChatWebsocketParam) *ChatWs {
	return &ChatWs{Logger: param.Logger}
}
