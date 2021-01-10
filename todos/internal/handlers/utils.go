package handlers

import (
	"github.com/code7unner/rest-api-test-task/todos/internal/service"
	"time"
)

func errorResponse(message string) map[string]string {
	return map[string]string{
		"error": message,
	}
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

