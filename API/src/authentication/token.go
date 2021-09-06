package authentication

import (
	"api/src/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//CreateToken returns a signed token with the user's perimitions
func CreateToken(userId uint64) (string, error) {
	permitions := jwt.MapClaims{}
	permitions["userId"] = userId
	permitions["authorized"] = true
	permitions["exp"] = time.Now().Add(time.Hour * 2).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permitions)

	return token.SignedString([]byte(config.SecretKey)) //Secret
}
