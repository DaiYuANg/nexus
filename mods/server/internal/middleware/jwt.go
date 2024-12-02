package middleware

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"nexus/internal/service"
	"strings"
)

func JWTMiddleware(logger *zap.Logger, jwt *service.JWT) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader == "" {
			// 如果没有 Authorization 头，返回错误
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization header missing",
			})
		}

		// Bearer Token 格式检查
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization format must be Bearer <token>",
			})
		}
		// 获取 Token 字符串
		tokenString := parts[1]
		token, err := jwt.Parse(tokenString)
		if err != nil {
			// 如果 token 解析失败，返回错误
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired token",
			})
		}
		// 将 Claims 提取出来并附加到请求上下文
		if token.Valid {
			issuer, err := token.Claims.GetIssuer()
			if err != nil {
				return err
			}
			c.Locals("userID", issuer)
		} else {
			// 如果 token 无效，返回错误
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token claims",
			})
		}
		return c.Next()
	}
}
