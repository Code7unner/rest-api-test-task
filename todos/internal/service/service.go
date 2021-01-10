package service

import (
	"github.com/code7unner/rest-api-test-task/todos/internal/models"
	"time"
)

//go:generate mockgen -source=$GOFILE -destination=mock/service_mock.go -package=service_mock Service

type Service interface {
	CreateTodo(userID int, title string, desc string, t time.Time) (*models.Todos, error)
	UpdateTodo(id int, userID int, title string, desc string, t time.Time) (*models.Todos, error)
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

func (s service) CreateTodo(userID int, title string, desc string, t time.Time) (*models.Todos, error) {
	todo := &models.Todos{
		UserID:         userID,
		Title:          title,
		Description:    desc,
		TimeToComplete: t,
	}

	if _, err := s.todos.Create(todo); err != nil {
		// TODO
		return nil, err
	}

	return todo, nil
}

func (s service) UpdateTodo(id int, userID int, title string, desc string, t time.Time) (*models.Todos, error) {
	todo := &models.Todos{
		ID:             id,
		UserID:         userID,
		Title:          title,
		Description:    desc,
		TimeToComplete: t,
	}

	if _, err := s.todos.Update(todo); err != nil {
		// TODO
		return nil, err
	}

	return todo, nil
}
