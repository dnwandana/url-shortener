package middleware

import (
	"github.com/dnwandana/url-shortener/model"
	"github.com/gofiber/fiber/v2"
)

// CookieRequired is a middleware function that check whether the user has a cookie header or not.
func CookieRequired() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := ctx.Cookies("token")
		// if the user does not have cookie, the application will return a JSON error
		if token == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponse{
				StatusCode: fiber.StatusUnauthorized,
				Error:      "Missing or Malformed Cookies",
			})
		}
		// on some route, the application will allow and serve the request
		return ctx.Next()
	}
}
