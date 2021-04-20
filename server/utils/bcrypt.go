package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword return the bcrypt hash of the given password.
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// CompareHashAndPassword compares a bcrypt hashed password with plaintext.
// Return true on success, or false on failure.
func VerifyPassword(hashPwd, plainPwd string) bool {
	byteHashPwd := []byte(hashPwd)
	bytePlainPwd := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHashPwd, bytePlainPwd)
	return err == nil
}
