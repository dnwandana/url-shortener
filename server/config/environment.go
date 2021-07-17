package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

// Env retrieves the value of the OS environment variable named by the key.
// Env will return empty string if the variable is not present.
func Env(key string) string {
	return os.Getenv(key)
}
