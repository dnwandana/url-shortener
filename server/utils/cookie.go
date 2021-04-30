package utils

import (
	"log"
	"strconv"
	"time"

	"github.com/dnwandana/url-shortener/config"
	"github.com/gofiber/fiber/v2"
)

// SetCookies return pointer cookie, which can be used to access private routes.
func SetCookies(name, value string) *fiber.Cookie {
	// get ttl token in hour
	jwtLife, jwtLifeErr := strconv.Atoi(config.Env("JWT_LIFE"))
	// check if there is an errror
	if jwtLifeErr != nil {
		log.Fatal("=> jwtLifeErr error:", jwtLifeErr)
	}
	// setting cookie
	cookie := new(fiber.Cookie)
	cookie.Name = name
	cookie.Value = value
	cookie.Expires = time.Now().Add(time.Duration(jwtLife) * time.Hour)
	// disable secure and HttpOnly cookie when development
	stage := config.Env("STAGE")
	if stage != "DEVELOPMENT" {
		cookie.Secure = true
		cookie.HTTPOnly = true
	}
	return cookie
}
