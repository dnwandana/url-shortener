package utils

import (
	"log"
	"strconv"
	"time"

	"github.com/dnwandana/url-shortener/config"
	"github.com/dnwandana/url-shortener/models"
	"github.com/form3tech-oss/jwt-go"
)

func GenerateJWT(user *models.User) (string, error) {
	jwtLife, jwtLifeErr := strconv.Atoi(config.Env("JWT_LIFE"))
	if jwtLifeErr != nil {
		log.Fatal("=> jwtLifeErr error:", jwtLifeErr)
	}

	data := jwt.New(jwt.SigningMethodHS256)

	claims := data.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["fullname"] = user.Fullname
	claims["exp"] = time.Now().Add(time.Duration(jwtLife) * time.Hour).Unix()

	token, tokenErr := data.SignedString([]byte(config.Env("JWT_SECRET")))
	if tokenErr != nil {
		return "", tokenErr
	}
	return token, nil
}

func ExtractIDFromJWT(token *jwt.Token) string {
	claims := token.Claims.(jwt.MapClaims)
	id := claims["id"].(string)
	return id
}
