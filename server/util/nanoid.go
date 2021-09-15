package util

import gonanoid "github.com/matoous/go-nanoid/v2"

func GenerateNanoID(size int) (string, error) {
	alphabet := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	nanoid, err := gonanoid.Generate(alphabet, size)
	if err != nil {
		return "", err
	}

	return nanoid, nil
}
