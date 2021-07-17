package controller

import (
	"github.com/dnwandana/url-shortener/entities"
	"github.com/dnwandana/url-shortener/middleware"
	"github.com/dnwandana/url-shortener/models"
	"github.com/dnwandana/url-shortener/services"
	"github.com/dnwandana/url-shortener/utils"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"time"
)

type UrlController struct {
	UrlService services.UrlService
}

func NewUrlController(urlService *services.UrlService) UrlController {
	return UrlController{
		UrlService: *urlService,
	}
}

// SetupRoutes Setup endpoint, parameter, middleware, and handler.
func (controller *UrlController) SetupRoutes (app *fiber.App) {
	app.Get("/go", middleware.CookieRequired(), middleware.JWTRequired(), controller.List())
	app.Post("/go", controller.Create())
	app.Get("/go/:id", controller.Get())
	app.Put("/go/:id", middleware.CookieRequired(), middleware.JWTRequired(), controller.Update())
	app.Delete("/go/:id", middleware.CookieRequired(), middleware.JWTRequired(), controller.Remove())

}

// List handler which handle request to list all shortUrls belonging to that user.
func (controller *UrlController) List() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// getting token from cookies
		token := ctx.Locals("user").(*jwt.Token)
		userID := utils.ExtractIDFromJWT(token)
		// execute request
		result, err := controller.UrlService.List(userID)
		// check if there is an error
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      err.Error(),
			})
		}
		// return list all shortUrls belonging to that user
		return ctx.JSON(fiber.Map{
			"statusCode": fiber.StatusOK,
			"url":        result,
		})
	}
}

// Create handler which handle request for creating a new shortUrl.
func (controller *UrlController) Create() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// parse data from request body
		var data *models.UrlData
		parserErr := ctx.BodyParser(&data)
		// check if there is an error
		if parserErr != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      parserErr.Error(),
			})
		}
		// validate the request body
		validationErr := utils.Validate(data)
		// check if there is an error
		if validationErr != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      validationErr,
			})
		}
		// getting userId from cookies
		// if there is no userId cookies then the UserID will be set to an empty string
		userID := ctx.Cookies("userId")
		// generate nanoid
		nanoid, nanoidErr := utils.GenerateNanoID()
		// check if there is an error
		if nanoidErr != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"statusCode": fiber.StatusInternalServerError,
				"error":      nanoidErr.Error(),
			})
		}
		// set data to Url struct
		url := entities.Url{
			UserID:    userID,
			ID:        nanoid,
			Title:     data.Title,
			URL:       data.URL,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		// execute the request
		result, dbErr := controller.UrlService.Create(&url)
		// check if there is an error
		if dbErr != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
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
		return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
			"statusCode": fiber.StatusCreated,
			"url":        response,
		})
	}
}

// Get handler which handle request for getting to specific URL Resource.
func (controller *UrlController) Get() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// getting id from request parameter
		id := ctx.Params("id")
		// execute the request
		result, err := controller.UrlService.Get(id)
		// check if there is an error
		if err != nil {
			// if there is an error, application will send to `/404` endpoint
			return ctx.Redirect("/404")
		}
		// if there are no error, application will send to specific URL
		return ctx.Redirect(result.URL, fiber.StatusMovedPermanently)
	}
}

// Update handler which handle request for updating the existing shortUrl.
func (controller *UrlController) Update() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// getting id from request parameter
		id := ctx.Params("id")
		// parse data from request body
		var data *models.UrlData
		parserErr := ctx.BodyParser(&data)
		// check if there is an error
		if parserErr != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      parserErr.Error(),
			})
		}
		// validate the request body
		validationErr := utils.Validate(data)
		// check if there is an error
		if validationErr != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      validationErr,
			})
		}
		// set new data to Url struct
		newUrl := entities.Url{
			ID:        data.ID,
			Title:     data.Title,
			URL:       data.URL,
			UpdatedAt: time.Now(),
		}
		// execute the request
		result, dbErr := controller.UrlService.Update(id, &newUrl)
		// check if there is an error
		if dbErr != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      dbErr.Error(),
			})
		}
		// get updated data
		url, urlErr :=  controller.UrlService.Get(data.ID)
		// check if there is an error
		if urlErr != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
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
		return ctx.JSON(fiber.Map{
			"statusCode": fiber.StatusOK,
			"url":        response,
		})
	}
}

// Remove handler which handle request for deleting the existing shortUrl.
func (controller *UrlController) Remove() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// getting id from request parameter
		id := ctx.Params("id")
		// execute the request
		err := controller.UrlService.Remove(id)
		// check if there is an error
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      err.Error(),
			})
		}
		return ctx.SendStatus(fiber.StatusOK)
	}
}
