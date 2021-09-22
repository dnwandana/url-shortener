package main

import (
	"github.com/dnwandana/url-shortener/config"
	"github.com/dnwandana/url-shortener/controller"
	"github.com/dnwandana/url-shortener/exception"
	"github.com/dnwandana/url-shortener/repository"
	"github.com/dnwandana/url-shortener/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// get database connection
	db := config.DatabaseConnection()

	// setup repository, service, and controller
	urlRepo := repository.NewURLRepository(db)
	urlService := service.NewURLService(&urlRepo)
	urlController := controller.NewURLController(&urlService)

	// instantitate fiber application
	app := fiber.New(
		fiber.Config{
			ErrorHandler: exception.ErrorHandler,
		},
	)

	// enable cors
	app.Use(cors.New())
	// recover panic
	app.Use(recover.New())

	// setting group prefix api v1
	v1 := app.Group("/api/v1")

	// setup routes
	urlController.SetupRoutes(v1)

	// listen to port `:5000` and log any errors
	panic(app.Listen(":5000"))
}
