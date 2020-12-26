package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

//struct endcoded to jwt:
type Claims struct {
	username           string
	jwt.StandardClaims //for fiels like 'expiry time'
}

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

/*Uses app's signing key to parse the jwt string, returns the jwt token and nil as error if eveything is fine */
func VerifyToken(signingKey []byte, jwtStr string) (*jwt.Token, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(jwtStr, claims, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	return token, err
}
