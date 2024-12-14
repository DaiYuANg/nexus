package server

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"nexus/internal/model"
	"nexus/internal/server/route"
	"strings"
)

type JWTMiddleware struct {
	fx.In
	*fiber.App
	SigningKey []byte        `name:"jwtKey"`
	Routes     []route.Route `group:"route"`
	*zap.SugaredLogger
}

func registerJwt(middleware JWTMiddleware) {
	middleware.App.Use(jwtware.New(jwtware.Config{
		Filter: func(ctx *fiber.Ctx) bool {
			if strings.HasPrefix(ctx.Path(), "/swagger/") {
				return true
			}
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
			middleware.Info("Context:%s %e", ctx.Path(), zap.Error(err))
			return ctx.Status(fiber.StatusUnauthorized).JSON(model.Err())
		},
	}))
}
