package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"nexus/internal/ws"
)

var websocket = fx.Module("websocket",
	fx.Provide(
		fx.Annotate(
			ws.NewChatWebsocket,
			fx.As(new(ws.Websocket)),
			fx.ResultTags(ws.Tag),
		),
		fx.Annotate(
			ws.NewLiveNotification,
			fx.As(new(ws.Websocket)),
			fx.ResultTags(ws.Tag),
		),
	),
	fx.Invoke(registerWebsocket),
)

type RegisterWebsocketParam struct {
	fx.In
	Engine    *fiber.App
	Websocket []ws.Websocket `group:"websocket"`
	*zap.Logger
}

func registerWebsocket(param RegisterWebsocketParam) {
	param.Info("Websockets", zap.Any("ws", len(param.Websocket)))
	app := param.Engine
	lo.ForEach(param.Websocket, func(item ws.Websocket, index int) {
		app.Get(item.Path(), item.UpgradeHandler())
		app.Get(item.Path(), item.ConnectHandler())
	})
}
