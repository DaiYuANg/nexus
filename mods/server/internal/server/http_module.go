package server

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var HttpModule = fx.Module("http",
	fx.Provide(newHttpServer),
	websocket,
	middleware,
	fx.Invoke(
		httpLifecycle,
	),
)

func newHttpServer() *fiber.App {
	return fiber.New(fiber.Config{EnablePrintRoutes: true, Immutable: true})
}

type RegisterParam struct {
	fx.In
	*fiber.App
	*zap.Logger
	SigningKey []byte `name:"jwtKey"`
}
