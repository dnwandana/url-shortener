package controller

import (
	"github.com/dnwandana/url-shortener/middleware"
	"github.com/dnwandana/url-shortener/model"
	"github.com/dnwandana/url-shortener/service"
	"github.com/dnwandana/url-shortener/util"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type UrlController struct {
	UrlService service.UrlService
}

func NewUrlController(urlService *service.UrlService) UrlController {
	return UrlController{
		UrlService: *urlService,
	}
}

// SetupRoutes Setup endpoint, parameter, middleware, and handler.
func (controller *UrlController) SetupRoutes(app *fiber.App) {
	app.Get("/go", middleware.CookieRequired(), middleware.JWTRequired(), controller.List)
	app.Post("/go", controller.Create)
	app.Get("/go/:id", controller.Get)
	app.Put("/go/:id", middleware.CookieRequired(), middleware.JWTRequired(), controller.Update)
	app.Delete("/go/:id", middleware.CookieRequired(), middleware.JWTRequired(), controller.Remove)

}

// List handler which handle request to list all shortUrls belonging to that user.
func (controller *UrlController) List(ctx *fiber.Ctx) error {
	// getting token from cookies
	token := ctx.Locals("user").(*jwt.Token)
	userID := util.ExtractIDFromJWT(token)

	// execute request
	result := controller.UrlService.List(userID)

	// return list all shortUrls belonging to that user
	return ctx.JSON(model.SuccessResponse{
		StatusCode: fiber.StatusOK,
		Data:       result,
	})
}

// Create handler which handle request for creating a new shortUrl.
func (controller *UrlController) Create(ctx *fiber.Ctx) error {
	// parse data from request body
	var request *model.UrlData
	parserErr := ctx.BodyParser(&request)
	// check if there is an error
	util.ReturnErrorIfNeeded(parserErr)

	// getting userId from cookies
	// if there is no userId cookies then the UserID will be set to an empty string
	userID := ctx.Cookies("userId")

	// execute the request
	result := controller.UrlService.Create(userID, request)

	// send another struct for JSON response
	return ctx.Status(fiber.StatusCreated).JSON(model.SuccessResponse{
		StatusCode: fiber.StatusCreated,
		Data:       result,
	})
}

// Get handler which handle request for getting to specific URL Resource.
func (controller *UrlController) Get(ctx *fiber.Ctx) error {
	// getting id from request parameter
	id := ctx.Params("id")
	// execute the request
	result, err := controller.UrlService.Get(id)
	// check if there is an error
	if err != nil {
		// if there is an error, application will send to `/404` endpoint
		return ctx.Redirect("/404", fiber.StatusMovedPermanently)
	}
	// if there are no error, application will send to specific URL
	return ctx.Redirect(result.URL, fiber.StatusMovedPermanently)
}

// Update handler which handle request for updating the existing shortUrl.
func (controller *UrlController) Update(ctx *fiber.Ctx) error {
	// getting id from request parameter
	id := ctx.Params("id")
	// parse data from request body
	var request *model.UrlData
	parserErr := ctx.BodyParser(&request)
	// check if there is an error
	util.ReturnErrorIfNeeded(parserErr)

	// execute the request
	result := controller.UrlService.Update(id, request)

	// send another struct for JSON response
	return ctx.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		StatusCode: fiber.StatusOK,
		Data:       result,
	})
}

// Remove handler which handle request for deleting the existing shortUrl.
func (controller *UrlController) Remove(ctx *fiber.Ctx) error {
	// getting id from request parameter
	id := ctx.Params("id")
	// execute the request
	controller.UrlService.Remove(id)
	// data deleted
	return ctx.SendStatus(fiber.StatusOK)
}
