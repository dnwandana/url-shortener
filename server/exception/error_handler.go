package exception

import (
	"github.com/dnwandana/url-shortener/model"
	"github.com/gofiber/fiber/v2"
)

// ErrorHandler is a function that create custom error handler.
func ErrorHandler(ctx *fiber.Ctx, err error) error {
	// check if the error message from BadRequestError struct
	_, ok := err.(BadRequestError)
	if ok {
		// if it is from BadRequestError
		// client will see 400 BadRequest response
		return ctx.Status(fiber.StatusBadRequest).JSON(model.ResponseMessage{
			Code:    fiber.StatusBadRequest,
			Status:  model.StatusBadRequest,
			Message: err.Error(),
		})
	}

	// the default error
	// client will see 500 InternalServerError response
	return ctx.Status(fiber.StatusInternalServerError).JSON(model.ResponseMessage{
		Code:    fiber.StatusInternalServerError,
		Status:  model.StatusInternalServerError,
		Message: err.Error(),
	})
}
