package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

// Env retrieves the value of the OS environment variable named by the key.
// Env will return empty string if the variable is not present.
func Env(key string) string {
	return os.Getenv(key)
}
