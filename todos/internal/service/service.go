package service

import (
	"github.com/code7unner/rest-api-test-task/todos/models"
)

//go:generate mockgen -source=$GOFILE -destination=mock/service_mock.go -package=service_mock Service

type Service interface {
}

type service struct {
	todos     models.TodosImpl
	jwtSecret string
	expires   int
}

func New(todosModel models.TodosImpl, jwtSecret string, expires int) Service {
	return &service{
		todos:     todosModel,
		jwtSecret: jwtSecret,
		expires:   expires,
	}
}