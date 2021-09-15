package main

import (
	"github.com/dnwandana/url-shortener/config"
	"github.com/dnwandana/url-shortener/controller"
	"github.com/dnwandana/url-shortener/repository"
	"github.com/dnwandana/url-shortener/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// get database connection
	db, err := config.DatabaseConnection()
	if err != nil {
		panic(err)
	}

	// setup repository, service, and controller
	urlRepo := repository.NewURLRepository(db)
	urlService := service.NewURLService(&urlRepo)
	urlController := controller.NewURLController(&urlService)

	// instantitate fiber application
	app := fiber.New()
	// enable cors
	app.Use(cors.New())

	// setup routes
	urlController.SetupRoutes(app)

	// listen to port `:5000` and log any errors
	panic(app.Listen(":5000"))
}
