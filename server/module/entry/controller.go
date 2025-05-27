package entry

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
	"go.uber.org/fx"
)

type Controller interface {
	RegisterRoutes(app *fiber.App)
}

var controllerModule = fx.Module("controller_module",
	fx.Provide(
		fx.Annotate(
			newUploadController,
			fx.As(new(Controller)),
			fx.ResultTags(`group:"controllers"`),
		),
		fx.Annotate(
			newNamespaceController,
			fx.As(new(Controller)),
			fx.ResultTags(`group:"controllers"`),
		),
	),
	fx.Invoke(bindingController))

type BindingParams struct {
	fx.In
	App        *fiber.App
	Controller []Controller `group:"controllers"`
}

func bindingController(params BindingParams) {
	lo.ForEach(params.Controller, func(item Controller, index int) {
		item.RegisterRoutes(params.App)
	})
}
