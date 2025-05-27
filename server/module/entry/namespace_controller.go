package entry

import (
	"github.com/DaiYuANg/storix/server/module/namespace"
	"github.com/gofiber/fiber/v2"
)

type namespaceController struct {
	store *namespace.Store
}

func (n namespaceController) RegisterRoutes(app *fiber.App) {
	group := app.Group("namespace")
	group.Post("/create/:name", n.createNamespace)
}

func (n namespaceController) createNamespace(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	return n.store.CreateNamespace(name)
}

func newNamespaceController(store *namespace.Store) *namespaceController {
	return &namespaceController{
		store: store,
	}
}
