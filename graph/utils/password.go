package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hash original password
func HashPassword(password string) (string, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(pass), err
}

// ComparePassword compare password with hash
func ComparePassword(password string, hash string) bool {
	byteHash := []byte(hash)
	bytePassword := []byte(password)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePassword)
	if err != nil {
		return false
	}

	return true
}
