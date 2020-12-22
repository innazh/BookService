package services

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

/*Creates a token with claims, it expires in 5 minutes*/
func CreateToken(appKey []byte, username string) (string, time.Time, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), // In JWT, the expiry time is expressed as unix milliseconds
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString(appKey)
	if err != nil {
		return "", expirationTime, err
	}

	return tokenStr, expirationTime, nil
}

func VerifyToken(tokenStr string) (*jwt.Token, error) {
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		//return tokenStr, nil //i think it's the app key
	})

	return tkn, err
}
