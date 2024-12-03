package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mileusna/useragent"
	"go.uber.org/zap"
)

func AgentRecorder(logger *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		agent := c.Get("User-Agent")
		ua := useragent.Parse(agent)
		logger.Info("UserAgent", zap.Any("ua", ua))
		return c.Next()
	}
}
