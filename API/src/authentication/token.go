package authentication

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

//ValidateToken verifies if the token passed in the request is valid
func ValidateToken(r *http.Request) error {
	tokenString := getToken(r)
	token, err := jwt.Parse(tokenString, returnAuthentificationKey)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return errors.New("invalid token")
}

//GetUserID returns the user's id that is saved in the token
func GetUserID(r *http.Request) (uint64, error) {
	tokenString := getToken(r)
	token, err := jwt.Parse(tokenString, returnAuthentificationKey)
	if err != nil {
		return 0, err
	}

	if perimitions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, err := strconv.ParseUint(fmt.Sprintf("%.0f", perimitions["userId"]), 10, 64)
		if err != nil {
			return 0, err
		}
		return userID, nil
	}
	return 0, errors.New("invalid token")
}

func getToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func returnAuthentificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}
