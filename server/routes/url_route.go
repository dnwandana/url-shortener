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

// Setup endpoint, parameter, middleware, and handler.
func UrlRoutes(app fiber.Router, service services.UrlService) {
	app.Get("/", middleware.CookieRequired(), middleware.JWTRequired(), getUrls(service))
	app.Post("/", addUrl(service))
	app.Get("/:id", getUrl(service))
	app.Put("/:id", middleware.CookieRequired(), middleware.JWTRequired(), updateUrl(service))
	app.Delete("/:id", middleware.CookieRequired(), middleware.JWTRequired(), deleteUrl(service))
}

// getUrls handler which handle request to list all shortUrls belonging to that user.
func getUrls(service services.UrlService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// getting token from cookies
		token := c.Locals("user").(*jwt.Token)
		userID := utils.ExtractIDFromJWT(token)
		// execute request
		result, err := service.ListAllShortUrl(userID)
		// check if there is an error
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      err.Error(),
			})
		}
		// return list all shortUrls belonging to that user
		return c.JSON(fiber.Map{
			"statusCode": fiber.StatusOK,
			"url":        result,
		})
	}
}

// addUrl handler which handle request for creating a new shortUrl.
func addUrl(service services.UrlService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// parse data from request body
		var data *models.UrlForm
		parserErr := c.BodyParser(&data)
		// check if there is an error
		if parserErr != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      parserErr.Error(),
			})
		}
		// validate the request body
		validationErr := utils.Validate(data)
		// check if there is an error
		if validationErr != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      validationErr,
			})
		}
		// getting userId from cookies
		// if there is no userId cookies then the UserID will be set to an empty string
		userID := c.Cookies("userId")
		// generate nanoid
		nanoid, nanoidErr := utils.GenerateNanoID()
		// check if there is an error
		if nanoidErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"statusCode": fiber.StatusInternalServerError,
				"error":      nanoidErr.Error(),
			})
		}
		// set data to Url struct
		url := models.Url{
			UserID:    userID,
			ID:        nanoid,
			Title:     data.Title,
			URL:       data.URL,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		// execute the request
		result, dbErr := service.CreateShortUrl(&url)
		// check if there is an error
		if dbErr != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      dbErr.Error(),
			})
		}
		// send another struct for JSON response
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

// getUrl handler which handle request for getting to specific URL Resource.
func getUrl(service services.UrlService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// getting id from request parameter
		id := c.Params("id")
		// execute the request
		result, err := service.GetShortUrl(id)
		// check if there is an error
		if err != nil {
			// if there is an error, application will send to `/404` endpoint
			return c.Redirect("/404", fiber.StatusNotFound)
		}
		// if there are no error, application will send to specific URL
		return c.Redirect(result.URL)
	}
}

// updateUrl handler which handle request for updating the existing shortUrl.
func updateUrl(service services.UrlService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// getting id from request parameter
		id := c.Params("id")
		// parse data from request body
		var data *models.UrlForm
		parserErr := c.BodyParser(&data)
		// check if there is an error
		if parserErr != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      parserErr.Error(),
			})
		}
		// validate the request body
		validationErr := utils.Validate(data)
		// check if there is an error
		if validationErr != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      validationErr,
			})
		}
		// set new data to Url struct
		newUrl := models.Url{
			ID:        data.ID,
			Title:     data.Title,
			URL:       data.URL,
			UpdatedAt: time.Now(),
		}
		// execute the request
		result, dbErr := service.UpdateShortUrl(id, &newUrl)
		// check if there is an error
		if dbErr != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      dbErr.Error(),
			})
		}
		// get updated data
		url, urlErr := service.GetShortUrl(data.ID)
		// check if there is an error
		if urlErr != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      urlErr.Error(),
			})
		}
		// send another struct for JSON response
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

// deleteUrl handler which handle reqeuest for deleting the existing shortUrl.
func deleteUrl(service services.UrlService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// getting id from request parameter
		id := c.Params("id")
		// execute the request
		err := service.DeleteShortUrl(id)
		// check if there is an error
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      err.Error(),
			})
		}
		return c.SendStatus(fiber.StatusNoContent)
	}
}
