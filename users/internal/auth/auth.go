package auth

import (
	"github.com/code7unner/rest-api-test-task/users/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func CreateToken(user *models.Users, secret string, expires int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(expires)).Unix()

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, nil
}