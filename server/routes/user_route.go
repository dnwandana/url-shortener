package routes

import (
	"time"

	"github.com/dnwandana/url-shortener/models"
	"github.com/dnwandana/url-shortener/services"
	"github.com/dnwandana/url-shortener/utils"
	"github.com/gofiber/fiber/v2"
)

// Setup endpoint, parameter, middleware, and handler.
func UserRoutes(app fiber.Router, service services.UserService) {
	app.Post("/sign-up", signUp(service))
	app.Post("/sign-in", signIn(service))
}

// signUp handler which handle request for creating a new user.
func signUp(service services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// parse data from request body
		var data *models.UserSignUp
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
		// check if the email is already registered
		isEmailExist, _ := service.Find("email", data.Email)
		// if the email is already registered send a JSON error
		if isEmailExist != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      "email already exist",
			})
		}
		// password encryption
		hashedPassword, hashErr := utils.HashPassword(data.Password)
		// check if there is an error
		if hashErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"statusCode": fiber.StatusInternalServerError,
				"error":      hashErr.Error(),
			})
		}
		// set data to User struct
		user := models.User{
			Fullname:  data.Fullname,
			Email:     data.Email,
			Password:  hashedPassword,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		// execute the request
		result, dbErr := service.Create(&user)
		// check if there is an error
		if dbErr != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      dbErr.Error(),
			})
		}
		// send another struct for JSON response
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

// signIn handler which handle request for getting cookies and JWT token
func signIn(service services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// parse data from request body
		var data *models.UserSignIn
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
		// check if the user exist with the given email
		user, _ := service.Find("email", data.Email)
		// send an error if the user does not exist
		if user == nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      "no user exist",
			})
		}
		// compare if the given password is the same as the user password from the database
		isPasswordMatch := utils.VerifyPassword(user.Password, data.Password)
		// send an error if the provided password are not the same
		if !isPasswordMatch {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      "invalid password",
			})
		}
		// generate JWT Token
		token, tokenErr := utils.GenerateJWT(user)
		// check if there is an error
		if tokenErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"statusCode": fiber.StatusInternalServerError,
				"error":      tokenErr.Error(),
			})
		}
		// set cookies
		userIdCookie := utils.SetCookies("userId", user.ID.Hex())
		jwtCookie := utils.SetCookies("token", token)
		c.Cookie(userIdCookie)
		c.Cookie(jwtCookie)
		return c.SendStatus(fiber.StatusOK)
	}
}
