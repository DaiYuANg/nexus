package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/samber/lo"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
	"nexus/internal/route"
	"nexus/web"
)

var RouteModule = fx.Module("route",
	fx.Provide(
		fx.Annotate(
			route.NewRuntimeRoute,
			fx.As(new(route.Route)),
			fx.ResultTags(route.Tag),
		),
		fx.Annotate(
			route.NewUserRoute,
			fx.As(new(route.Route)),
			fx.ResultTags(route.Tag),
		),
		fx.Annotate(
			route.NewUpload,
			fx.As(new(route.Route)),
			fx.ResultTags(route.Tag),
		),
		fx.Annotate(
			route.NewFileRoute,
			fx.As(new(route.Route)),
			fx.ResultTags(route.Tag),
		),
	),
	fx.Invoke(registerRoute, registerUI),
)

type RegisterJwtRouteParam struct {
	fx.In
	Engine *fiber.App
	Routes []route.Route `group:"route"`
	Logger *zap.Logger
}

func registerRoute(param RegisterJwtRouteParam) {
	log, engine, routes := param.Logger, param.Engine, param.Routes
	log.Info("Routes:", zap.Any("Routes", routes))
	routeInfos := lo.FlatMap[route.Route, route.Info](routes, func(item route.Route, index int) []route.Info {
		return item.Routes()
	})
	group := engine.Group("/api")
	lo.ForEach(
		routeInfos,
		func(item route.Info, index int) {
			switch item.Method {
			case route.GET:
				group.Get(item.Path, item.Handler)
			case route.POST:
				group.Post(item.Path, item.Handler)
			case route.PUT:
				group.Put(item.Path, item.Handler)
			case route.DELETE:
				group.Delete(item.Path, item.Handler)
			}
		},
	)
}

func registerUI(app *fiber.App) {
	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(web.Content),
		Index:      "index.html",
		PathPrefix: "dist",
		Browse:     true,
	}))
}
