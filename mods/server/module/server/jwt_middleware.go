package server

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
	"go.uber.org/fx"
	"nexus/internal/model"
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

			if lo.Contains(ignoreJwt, result) {
				return true
			}

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
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(model.Err())
		},
	}))
}
