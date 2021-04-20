package utils

import gonanoid "github.com/matoous/go-nanoid/v2"

func GenerateNanoID() (string, error) {
	var alphabet string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	nanoid, err := gonanoid.Generate(alphabet, 6)
	if err != nil {
		return "", err
	}
	return nanoid, nil
}
