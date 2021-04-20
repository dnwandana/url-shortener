package routes

import (
	"time"

	"github.com/dnwandana/url-shortener/middleware"
	"github.com/dnwandana/url-shortener/models"
	"github.com/dnwandana/url-shortener/services"
	"github.com/dnwandana/url-shortener/utils"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func UrlRoutes(app fiber.Router, service services.UrlService) {
	app.Get("/", middleware.CookieRequired(), middleware.JWTRequired(), getUrls(service))
	app.Post("/", addUrl(service))
	app.Get("/:id", getUrl(service))
	app.Put("/:id", middleware.CookieRequired(), middleware.JWTRequired(), updateUrl(service))
	app.Delete("/:id", middleware.CookieRequired(), middleware.JWTRequired(), deleteUrl(service))
}

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
		var data *models.UrlForm
		parserErr := c.BodyParser(&data)
		if parserErr != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      parserErr.Error(),
			})
		}
		validationErr := utils.Validate(data)
		if validationErr != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      validationErr,
			})
		}
		userID := c.Cookies("userId")
		nanoid, nanoidErr := utils.GenerateNanoID()
		if nanoidErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"statusCode": fiber.StatusInternalServerError,
				"error":      nanoidErr.Error(),
			})
		}
		url := models.Url{
			UserID:    userID,
			ID:        nanoid,
			Title:     data.Title,
			URL:       data.URL,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		result, dbErr := service.CreateShortUrl(&url)
		if dbErr != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      dbErr.Error(),
			})
		}
		response := models.UrlResponse{
			ID:        result.ID,
			Title:     result.Title,
			URL:       result.URL,
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.CreatedAt,
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
		var data *models.UrlForm
		parserErr := c.BodyParser(&data)
		if parserErr != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      parserErr.Error(),
			})
		}
		validationErr := utils.Validate(data)
		if validationErr != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      validationErr,
			})
		}
		newUrl := models.Url{
			ID:        data.ID,
			Title:     data.Title,
			URL:       data.URL,
			UpdatedAt: time.Now(),
		}
		result, dbErr := service.UpdateShortUrl(id, &newUrl)
		if dbErr != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      dbErr.Error(),
			})
		}
		url, urlErr := service.GetShortUrl(data.ID)
		if urlErr != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      urlErr.Error(),
			})
		}
		response := models.UrlResponse{
			ID:        result.ID,
			Title:     result.Title,
			URL:       result.URL,
			CreatedAt: url.CreatedAt,
			UpdatedAt: result.UpdatedAt,
		}
		return c.JSON(fiber.Map{
			"statusCode": fiber.StatusOK,
			"url":        response,
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
