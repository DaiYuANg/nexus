package endpoint

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"golang.org/x/net/webdav"
)

type WebDavController struct {
}

func (c *WebDavController) RegisterRoutes(app *fiber.App) {
	handler := &webdav.Handler{
		Prefix:     "/webdav/",
		FileSystem: webdav.Dir("./data"),
		LockSystem: webdav.NewMemLS(),
	}

	app.All("/webdav/*", adaptor.HTTPHandler(handler))
}

func NewWebDavController() *WebDavController {
	return &WebDavController{}
}
