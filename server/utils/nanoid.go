package utils

import gonanoid "github.com/matoous/go-nanoid/v2"

// GenerateNanoID return string 6 character unique id
func GenerateNanoID() (string, error) {
	alphabet := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	nanoid, err := gonanoid.Generate(alphabet, 6)
	// check if there is an error
	if err != nil {
		return "", err
	}
	return nanoid, nil
}
