package server

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"nexus/internal/route"
	"strings"
)

type JWTMiddleware struct {
	fx.In
	*fiber.App
	SigningKey []byte        `name:"jwtKey"`
	Routes     []route.Route `group:"route"`
}

func registerJwt(middleware JWTMiddleware) {
	middleware.App.Use(jwtware.New(jwtware.Config{
		Filter: func(ctx *fiber.Ctx) bool {
			result := strings.ReplaceAll(ctx.Path(), "/api", "")
			for _, r := range middleware.Routes {
				for _, info := range r.Routes() {
					if info.Path == result && info.PermitAll {
						return true
					}
				}
			}
			return false
		},
		SigningKey: jwtware.SigningKey{Key: middleware.SigningKey},
		SuccessHandler: func(ctx *fiber.Ctx) error {
			return ctx.Next()
		},
	}))
}
