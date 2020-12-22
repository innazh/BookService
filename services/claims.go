package services

import (
	"github.com/dgrijalva/jwt-go"
)

//struct endcoded to jwt:
type Claims struct {
	username           string `json:"username bson:"username"`
	jwt.StandardClaims        //for fiels like 'expiry time'
}
