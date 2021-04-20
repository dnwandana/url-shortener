package middleware

import (
	"github.com/dnwandana/url-shortener/config"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

// JWTRequired is a middleware function that check whether the user has a JWT Token or not.
func JWTRequired() fiber.Handler {
	return jwtware.New(jwtware.Config{
		ErrorHandler: jwtErrorHandler,
		SigningKey:   []byte(config.Env("JWT_SECRET")),
		TokenLookup:  "cookie:token",
	})
}

// Error handler for JWTRequired middleware
func jwtErrorHandler(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		// error handler if the request does not have a JWT Token
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"statusCode": fiber.StatusUnauthorized,
			"error":      "missing or malformed JWT",
		})
	}
	// error handler if the JWT Token invalid or has expired
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"statusCode": fiber.StatusUnauthorized,
		"error":      "invalid or expired JWT",
	})
}
