package http

import (
	"github.com/DaiYuANg/maxio/server/view"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/template/html/v2"
	"go.uber.org/fx"
	"net/http"
)

var Module = fx.Module("http",
	fx.Provide(
		newTemplateEngine,
		newFiber,
	),
	middlewareModule,
	controllerModule,
	fx.Invoke(startHttp),
)

func newTemplateEngine() *html.Engine {
	return html.NewFileSystem(http.FS(view.FS), ".html")
}

func newFiber(engine *html.Engine) *fiber.App {
	app := fiber.New(
		fiber.Config{
			Views:             engine,
			PassLocalsToViews: true,
			EnablePrintRoutes: true,
			BodyLimit:         1024 * 1024 * 1024,
		},
	)
	app.Get("/metrics", monitor.New())
	return app
}
