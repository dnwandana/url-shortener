package controller

import (
	"fmt"
	"os"

	"github.com/dnwandana/url-shortener/exception"
	"github.com/dnwandana/url-shortener/model"
	"github.com/dnwandana/url-shortener/service"
	"github.com/gofiber/fiber/v2"
)

type URLController struct {
	URLService service.URLService
}

func NewURLController(urlService *service.URLService) URLController {
	return URLController{URLService: *urlService}
}

func (controller *URLController) SetupRoutes(router fiber.Router) {
	router.Post("/go", controller.Create)
	router.Get("/go/:id", controller.FindOne)
	router.Delete("/go/:id", controller.Delete)
}

// Shorten URL
func (controller *URLController) Create(ctx *fiber.Ctx) error {
	// parse requestBody
	var request model.URLCreateRequest
	err := ctx.BodyParser(&request)
	exception.PanicIfNeeded(err)

	// execute request to shorten URL
	response := controller.URLService.Create(&request)
	// if there are no error, send created response
	return ctx.Status(fiber.StatusCreated).JSON(model.ResponseData{
		Code:   fiber.StatusCreated,
		Status: model.StatusCreated,
		Data:   response,
	})
}

// Go to Destination URL
func (controller *URLController) FindOne(ctx *fiber.Ctx) error {
	// get id from params
	id := ctx.Params("id")
	// execute request to get url
	url := controller.URLService.FindOne(id)
	// if no url found redirect to domain.tld/404
	if url == "" {
		notFoundPage := fmt.Sprintf("%s/404", os.Getenv("DOMAIN"))
		return ctx.Redirect(notFoundPage, fiber.StatusMovedPermanently)
	}

	// if url found send to destination url
	return ctx.Redirect(url, fiber.StatusMovedPermanently)
}

// Delete URL
func (controller *URLController) Delete(ctx *fiber.Ctx) error {
	// get id from params
	id := ctx.Params("id")
	// get secret_key from query
	secret_key := ctx.Query("secret_key")
	// execute request to delete url
	controller.URLService.Delete(id, secret_key)
	// if there are no error, send ok response
	return ctx.Status(fiber.StatusOK).JSON(model.ResponseMessage{
		Code:    fiber.StatusOK,
		Status:  model.StatusOK,
		Message: "Successfully deleted URL",
	})
}
