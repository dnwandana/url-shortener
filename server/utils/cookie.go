package utils

import (
	"log"
	"strconv"
	"time"

	"github.com/dnwandana/url-shortener/config"
	"github.com/gofiber/fiber/v2"
)

func SetCookies(token string) *fiber.Cookie {
	jwtLife, jwtLifeErr := strconv.Atoi(config.Env("JWT_LIFE"))
	if jwtLifeErr != nil {
		log.Fatal("=> jwtLifeErr error:", jwtLifeErr)
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(time.Duration(jwtLife) * time.Hour)
	// cookie.Secure = true
	// cookie.HTTPOnly = true

	return cookie
}
