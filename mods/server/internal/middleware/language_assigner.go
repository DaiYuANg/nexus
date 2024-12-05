package middleware

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func LanguageAssigner(logger *zap.Logger) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		language := ctx.GetRespHeader(fiber.HeaderAcceptLanguage, "Eng")
		ctx.Set(fiber.HeaderAcceptLanguage, language)
		logger.Info("User", zap.String("language", language))
		return ctx.Next()
	}
}
