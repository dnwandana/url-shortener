package exception

import (
	"github.com/dnwandana/url-shortener/model"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	_, ok := err.(BadRequestError)
	if ok {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.ResponseMessage{
			Code:    fiber.StatusBadRequest,
			Status:  model.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(model.ResponseMessage{
		Code:    fiber.StatusInternalServerError,
		Status:  model.StatusInternalServerError,
		Message: err.Error(),
	})
}