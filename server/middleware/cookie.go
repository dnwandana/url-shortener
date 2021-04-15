package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func CookieRequired() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Cookies("token")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"statusCode": fiber.StatusUnauthorized,
				"error":      "missing or malformed cookies",
			})
		}
		return c.Next()
	}
}
