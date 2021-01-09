package auth

import (
	"errors"
	"github.com/code7unner/rest-api-test-task/users/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func CreateToken(user *models.Users, secret string, expires int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["username"] = user.Username
	claims["password"] = user.Password
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(expires)).Unix()

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, nil
}

func GetPasswordFromToken(tokenStr, secret string) (string, error) {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return "", err
	}

	password, ok := claims["password"].(string)
	if !ok {
		return "", errors.New("could not parse claims")
	}

	return password, nil
}
