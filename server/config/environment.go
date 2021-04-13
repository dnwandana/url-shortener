package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func Env(key string) string {
	return os.Getenv(key)
}
