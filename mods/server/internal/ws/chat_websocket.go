package ws

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"log"
)

type ChatWs struct {
	*zap.Logger
}

func (cw *ChatWs) UpgradeHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(ctx) {
			ctx.Locals("allowed", true)
			return ctx.Next()
		}
		return fiber.ErrUpgradeRequired
	}
}

func (cw *ChatWs) Path() string {
	return "/chat/:userid"
}

func (cw *ChatWs) ConnectHandler() fiber.Handler {
	return websocket.New(
		func(c *websocket.Conn) {
			// c.Locals is added to the *websocket.Conn
			log.Println(c.Locals("allowed"))  // true
			log.Println(c.Params("userid"))   // 123
			log.Println(c.Query("v"))         // 1.0
			log.Println(c.Cookies("session")) // ""

			var (
				mt  int
				msg []byte
				err error
			)
			for {
				if mt, msg, err = c.ReadMessage(); err != nil {
					log.Println("read:", err)
					break
				}
				log.Printf("recv: %s", msg)

				if err = c.WriteMessage(mt, msg); err != nil {
					log.Println("write:", err)
					break
				}
			}
		},
		websocket.Config{
			EnableCompression: true,
			RecoverHandler: func(conn *websocket.Conn) {
				cw.Info("recover")
			},
		},
	)
}
