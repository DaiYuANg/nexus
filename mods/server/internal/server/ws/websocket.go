package ws

import "github.com/gofiber/fiber/v2"

const Tag = `group:"websocket"`

type Websocket interface {
	UpgradeHandler() fiber.Handler
	Path() string
	ConnectHandler() fiber.Handler
}
