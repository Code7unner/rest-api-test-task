package handlers

import (
	"errors"
	"github.com/code7unner/rest-api-test-task/todos/internal/service"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func errorResponse(message string) map[string]string {
	return map[string]string{
		"error": message,
	}
}

func getUserIDFromToken(t interface{}) (int, error) {
	userToken, ok := t.(*jwt.Token)
	if !ok {
		return 0, errors.New("invalid token")
	}
	claims := userToken.Claims.(jwt.MapClaims)
	userID := claims["id"].(float64)

	return int(userID), nil
}

func parseDate(value string) (time.Time, error) {
	if value != "" {
		return time.Parse(service.DateLayout, value)
	}

	return time.Time{}, nil
}

func parseDateTime(value string) (time.Time, error) {
	if value != "" {
		return time.Parse(service.DateTimeLayout, value)
	}

	return time.Time{}, nil
}

