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
		var data *models.UserSignUp
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
		isEmailExist, _ := service.Find("email", data.Email)
		if isEmailExist != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      "email already exist",
			})
		}
		hashedPassword, hashErr := utils.HashPassword(data.Password)
		if hashErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"statusCode": fiber.StatusInternalServerError,
				"error":      hashErr.Error(),
			})
		}
		user := models.User{
			Fullname:  data.Fullname,
			Email:     data.Email,
			Password:  hashedPassword,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		result, dbErr := service.Create(&user)
		if dbErr != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      dbErr.Error(),
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
		var data *models.UserSignIn
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
		user, _ := service.Find("email", data.Email)
		if user == nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      "no user exist",
			})
		}
		isPasswordMatch := utils.VerifyPassword(user.Password, data.Password)
		if !isPasswordMatch {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      "invalid password",
			})
		}
		token, tokenErr := utils.GenerateJWT(user)
		if tokenErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"statusCode": fiber.StatusInternalServerError,
				"error":      tokenErr.Error(),
			})
		}
		userIdCookie := utils.SetCookies("userId", user.ID.Hex())
		jwtCookie := utils.SetCookies("token", token)
		c.Cookie(userIdCookie)
		c.Cookie(jwtCookie)
		return c.SendStatus(fiber.StatusOK)
	}
}
