package services

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

//function used to hash passwords
func HashAndSalt(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	return string(hash), nil
}

//function used to compare passwords upon sign in
func ComparePasswords(registeredPass, signInPass []byte) bool {
	result := bcrypt.CompareHashAndPassword(registeredPass, signInPass)
	if result != nil {
		return false
	}

	return true
}
