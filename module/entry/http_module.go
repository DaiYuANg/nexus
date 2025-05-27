package entry

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/template/html/v2"
	"go.uber.org/fx"
	"net/http"
	"storix/view"
)

var HttpModule = fx.Module("http",
	fx.Provide(
		newTemplateEngine,
		newFiber,
	),
	controllerModule,
	fx.Invoke(startHttp),
)

func newTemplateEngine() *html.Engine {
	return html.NewFileSystem(http.FS(view.ViewFS), ".html")
}

func newFiber(engine *html.Engine) *fiber.App {

	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
		Views:             engine,
		ViewsLayout:       "layouts/main",
		PassLocalsToViews: false,
	})
	app.Get("/metrics", monitor.New())
	return app
}

func startHttp(lc fx.Lifecycle, app *fiber.App) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				err := app.Listen(":3000")
				if err != nil {
					panic(err)
				}
			}()
			return nil
		},
	})
}
