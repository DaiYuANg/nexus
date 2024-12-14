package server

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	middleware2 "nexus/internal/middleware"
)

func registerUaRecorder(app *fiber.App, logger *zap.Logger) {
	app.Use(middleware2.AgentRecorder(logger))
}
