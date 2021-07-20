package main

import (
	"github.com/dnwandana/url-shortener/config"
	"github.com/dnwandana/url-shortener/controller"
	"github.com/dnwandana/url-shortener/model"
	"github.com/dnwandana/url-shortener/repository"
	"github.com/dnwandana/url-shortener/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
)

func main() {
	// get database connection
	db, dbErr := config.DatabaseConnection()
	if dbErr != nil {
		log.Fatal("=> database error:", dbErr.Error())
	}

	// setup url collection, instantiate repository, service, and controller
	urlCollection := db.Collection(config.Env("URL_COLLECTION"))
	urlRepository := repository.NewUrlRepository(urlCollection)
	urlService := service.NewUrlService(&urlRepository)
	urlController := controller.NewUrlController(&urlService)

	// setup user collection, instantiate repository, service, and controller
	userCollection := db.Collection(config.Env("USER_COLLECTION"))
	userRepository := repository.NewUserRepository(userCollection)
	userService := service.NewUserService(&userRepository)
	userController := controller.NewUserController(&userService)

	// instantiate fiber application
	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Send custom error page
			if err != nil {
				// In case the SendFile fails
				return ctx.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
					StatusCode: fiber.StatusBadRequest,
					Error:      err.Error(),
				})
			}
			// Return from handler
			return nil
		},
	})

	// register cors middleware and allow browser expose credentials
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// receiver custom error handler
	app.Use(recover.New())

	// setup controller
	urlController.SetupRoutes(app)
	userController.SetupRoutes(app)

	// listen to port `:5000` and log any errors
	log.Fatal(app.Listen(":5000"))
}
