package middleware

import (
	"github.com/dnwandana/url-shortener/config"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func JWTRequired() fiber.Handler {
	return jwtware.New(jwtware.Config{
		ErrorHandler: jwtErrorHandler,
		SigningKey:   []byte(config.Env("JWT_SECRET")),
		TokenLookup:  "cookie:token",
	})
}

func jwtErrorHandler(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"statusCode": fiber.StatusUnauthorized,
			"error":      "missing or malformed JWT",
		})
	}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"statusCode": fiber.StatusUnauthorized,
		"error":      "invalid or expired JWT",
	})
}
