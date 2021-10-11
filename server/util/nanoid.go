package util

import (
	"github.com/dnwandana/url-shortener/exception"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

// Generate random id based on request size
func GenerateNanoID(size int) string {
	alphabet := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	nanoid, err := gonanoid.Generate(alphabet, size)
	exception.PanicIfNeeded(err)
	return nanoid
}
