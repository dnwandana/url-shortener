package routes

import (
	"time"

	"github.com/dnwandana/url-shortener/models"
	"github.com/dnwandana/url-shortener/services"
	"github.com/dnwandana/url-shortener/utils"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app fiber.Router, service services.UserService) {
	app.Post("/sign-up", signUp(service))
	app.Post("/sign-in", signIn(service))
}

func signUp(service services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var reqBody *models.UserSignUp
		err := c.BodyParser(&reqBody)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      err.Error(),
			})
		}
		isEmailExist, _ := service.Find("email", reqBody.Email)
		if isEmailExist != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      "email already exist",
			})
		}
		if reqBody.Password != reqBody.ConfirmationPassword {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      "confirmation password don't match",
			})
		}
		hashedPassword, hashErr := utils.HashPassword(reqBody.Password)
		if hashErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"statusCode": fiber.StatusInternalServerError,
				"error":      hashErr.Error(),
			})
		}
		data := models.User{
			Fullname:  reqBody.Fullname,
			Email:     reqBody.Email,
			Password:  hashedPassword,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		result, dbErr := service.Create(&data)
		if dbErr != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      err.Error(),
			})
		}
		response := models.UserSignUpResponse{
			ID:       result.ID,
			Fullname: result.Fullname,
		}
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"statusCode": fiber.StatusCreated,
			"user":       response,
		})
	}
}

func signIn(service services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var reqBody *models.User
		err := c.BodyParser(&reqBody)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      err.Error(),
			})
		}
		user, _ := service.Find("email", reqBody.Email)
		if user == nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      "no user exist",
			})
		}
		isPasswordMatch := utils.VerifyPassword(user.Password, reqBody.Password)
		if !isPasswordMatch {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      "invalid password",
			})
		}
		// TODO: return Cookies(JWT)
		return c.SendStatus(fiber.StatusOK)
	}
}
