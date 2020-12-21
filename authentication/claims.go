package authentication

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//struct endcoded to jwt:
type Claims struct {
	userId             primitive.ObjectID `json:"userId" bson:"userId"`
	username           string             `json:"username bson:"username"`
	jwt.StandardClaims                    //for fiels like 'expiry time'
}
