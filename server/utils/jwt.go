package utils

import (
	"log"
	"strconv"
	"time"

	"github.com/dnwandana/url-shortener/config"
	"github.com/dnwandana/url-shortener/models"
	"github.com/form3tech-oss/jwt-go"
)

// GenerateJWT return string JWT Token, which can be used to access private routes.
func GenerateJWT(user *models.User) (string, error) {
	// get ttl token in hour
	jwtLife, jwtLifeErr := strconv.Atoi(config.Env("JWT_LIFE"))
	// check if there is an error
	if jwtLifeErr != nil {
		log.Fatal("=> jwtLifeErr error:", jwtLifeErr)
	}
	// sign token with signin algorithms HS256
	data := jwt.New(jwt.SigningMethodHS256)
	// setting token payload
	claims := data.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["fullname"] = user.Fullname
	claims["exp"] = time.Now().Add(time.Duration(jwtLife) * time.Hour).Unix()
	// Get the complete signed token
	token, tokenErr := data.SignedString([]byte(config.Env("JWT_SECRET")))
	// check if there is an error
	if tokenErr != nil {
		return "", tokenErr
	}
	return token, nil
}

// ExtractIDFromJWT return string id, extracted from payload JWT Token.
func ExtractIDFromJWT(token *jwt.Token) string {
	claims := token.Claims.(jwt.MapClaims)
	id := claims["id"].(string)
	return id
}
