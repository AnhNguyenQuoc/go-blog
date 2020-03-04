package lib

import (
	"golang.org/x/crypto/bcrypt"
)

var secret = []byte{1, 2, 1, 2, 1, 9, 9, 6}

func HashPassword(pwd string) (string, error) {
	cost, _ := bcrypt.Cost(secret)
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), cost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func AuthenticatePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}
	return true
}
