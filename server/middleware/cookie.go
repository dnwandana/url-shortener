package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// CookieRequired is a middleware function that check whether the user has a cookie header or not.
func CookieRequired() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Cookies("token")
		// if the user does not have cookie, the application will return a JSON error
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"statusCode": fiber.StatusUnauthorized,
				"error":      "missing or malformed cookies",
			})
		}
		// on some route, the application will allow and serve the request
		return c.Next()
	}
}
