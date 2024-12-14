package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
	"go.uber.org/fx"
	"go.uber.org/zap"
	ws2 "nexus/internal/server/ws"
)

var websocket = fx.Module("websocket",
	fx.Provide(
		fx.Annotate(
			ws2.NewChatWebsocket,
			fx.As(new(ws2.Websocket)),
			fx.ResultTags(ws2.Tag),
		),
		fx.Annotate(
			ws2.NewLiveNotification,
			fx.As(new(ws2.Websocket)),
			fx.ResultTags(ws2.Tag),
		),
	),
	fx.Invoke(registerWebsocket),
)

type RegisterWebsocketParam struct {
	fx.In
	Engine    *fiber.App
	Websocket []ws2.Websocket `group:"websocket"`
	*zap.Logger
}

func registerWebsocket(param RegisterWebsocketParam) {
	param.Info("Websockets", zap.Any("ws", len(param.Websocket)))
	app := param.Engine
	lo.ForEach(param.Websocket, func(item ws2.Websocket, index int) {
		app.Get(item.Path(), item.UpgradeHandler())
		app.Get(item.Path(), item.ConnectHandler())
	})
}
