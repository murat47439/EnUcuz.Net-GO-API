package utils

import (
	"Store-Dio/config"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		config.Logger.Printf("Unsuccesful HashPassword")
		return "", err
	}
	return string(hash), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if err != nil {
		config.Logger.Printf("Unsuccesful ChechPasswordHash")
		return false
	}
	return true
}
