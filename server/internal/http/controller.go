package http

import (
	"github.com/DaiYuANg/storix/server/internal/controller"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
	"go.uber.org/fx"
)

type Controller interface {
	RegisterRoutes(app *fiber.App)
}

var controllerModule = fx.Module("controller_module",
	fx.Provide(
		newValidator,
		fx.Annotate(
			controller.NewUploadController,
			fx.As(new(Controller)),
			fx.ResultTags(`group:"controllers"`),
		),
		fx.Annotate(
			controller.NewNamespaceController,
			fx.As(new(Controller)),
			fx.ResultTags(`group:"controllers"`),
		),
		fx.Annotate(
			controller.NewNamespaceController,
			fx.As(new(Controller)),
			fx.ResultTags(`group:"controllers"`),
		),
		fx.Annotate(
			controller.NewWebDavController,
			fx.As(new(Controller)),
			fx.ResultTags(`group:"controllers"`),
		),
	),
	fx.Invoke(bindingController))

func newValidator() *validator.Validate {
	return validator.New(validator.WithRequiredStructEnabled())
}

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
