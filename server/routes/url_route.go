package routes

import (
	"time"

	"github.com/dnwandana/url-shortener/middleware"
	"github.com/dnwandana/url-shortener/models"
	"github.com/dnwandana/url-shortener/services"
	"github.com/dnwandana/url-shortener/utils"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func UrlRoutes(app fiber.Router, service services.UrlService) {
	app.Get("/", middleware.CookieRequired(), middleware.JWTRequired(), getUrls(service))
	app.Post("/", addUrl(service))
	app.Get("/:id", getUrl(service))
	app.Put("/:id", middleware.CookieRequired(), middleware.JWTRequired(), updateUrl(service))
	app.Delete("/:id", middleware.CookieRequired(), middleware.JWTRequired(), deleteUrl(service))
}

var alphabet string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func getUrls(service services.UrlService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Locals("user").(*jwt.Token)
		userID := utils.ExtractIDFromJWT(token)
		result, err := service.ListAllShortUrl(userID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      err.Error(),
			})
		}
		return c.JSON(&fiber.Map{
			"statusCode": fiber.StatusOK,
			"url":        result,
		})
	}
}

func addUrl(service services.UrlService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var reqBody *models.Url
		err := c.BodyParser(&reqBody)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      err.Error(),
			})
		}
		userID := c.Cookies("userId")
		if userID == "" {
			userID = ""
		}
		nanoid, nanoidErr := gonanoid.Generate(alphabet, 6)
		if nanoidErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"statusCode": fiber.StatusInternalServerError,
				"error":      nanoidErr.Error(),
			})
		}
		data := models.Url{
			UserID:    userID,
			ID:        nanoid,
			URL:       reqBody.URL,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		result, dberr := service.CreateShortUrl(&data)
		response := models.UrlResponse{
			ID:    result.ID,
			Title: result.Title,
			URL:   result.URL,
		}
		if dberr != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      dberr.Error(),
			})
		}
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"statusCode": fiber.StatusCreated,
			"url":        response,
		})
	}
}

func getUrl(service services.UrlService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		result, err := service.GetShortUrl(id)
		if err != nil {
			return c.Redirect("/404")
		}
		return c.Redirect(result.URL)
	}
}

func updateUrl(service services.UrlService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var reqBody *models.Url
		err := c.BodyParser(&reqBody)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      err.Error(),
			})
		}
		result, dberr := service.UpdateShortUrl(id, reqBody)
		if dberr != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      dberr.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"statusCode": fiber.StatusOK,
			"url":        result,
		})
	}
}

func deleteUrl(service services.UrlService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		err := service.DeleteShortUrl(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      err.Error(),
			})
		}
		return c.SendStatus(fiber.StatusNoContent)
	}
}
