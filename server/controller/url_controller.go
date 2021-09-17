package controller

import (
	"fmt"
	"os"

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

func (controller *URLController) Create(ctx *fiber.Ctx) error {
	var request model.URLCreateRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.ResponseMessage{
			Code:    fiber.StatusBadRequest,
			Status:  model.StatusBadRequest,
			Message: err.Error(),
		})
	}

	response, err := controller.URLService.Create(&request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.ResponseMessage{
			Code:    fiber.StatusBadRequest,
			Status:  model.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(model.ResponseData{
		Code:   fiber.StatusCreated,
		Status: model.StatusCreated,
		Data:   response,
	})
}

func (controller *URLController) FindOne(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	url, err := controller.URLService.FindOne(id)
	notFoundPage := fmt.Sprintf("%s/404", os.Getenv("DOMAIN"))
	if err != nil {
		return ctx.Redirect(notFoundPage, fiber.StatusMovedPermanently)
	}

	return ctx.Redirect(url, fiber.StatusMovedPermanently)
}

func (controller *URLController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	secret_key := ctx.Query("secret_key")
	err := controller.URLService.Delete(id, secret_key)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.ResponseMessage{
			Code:    fiber.StatusBadRequest,
			Status:  model.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(model.ResponseMessage{
		Code:    fiber.StatusOK,
		Status:  model.StatusOK,
		Message: "successfully deleted url",
	})
}
