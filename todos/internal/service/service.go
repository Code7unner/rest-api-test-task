package service

import (
	"github.com/code7unner/rest-api-test-task/todos/internal/models"
	"sort"
	"time"
)

//go:generate mockgen -source=$GOFILE -destination=mock/service_mock.go -package=service_mock Service

type Service interface {
	CreateTodo(userID int, title string, desc string, t time.Time) (*models.Todos, error)
	UpdateTodo(id int, userID int, title string, desc string, t time.Time) (*models.Todos, error)
	DeleteTodo(id int) error
	GetAllTodos(userID int) ([]models.Todos, error)
	GetAllCurrentTodos(userID int, t time.Time) ([]models.Todos, error)
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

func (s service) DeleteTodo(id int) error {
	todo, err := s.todos.Get(id)
	if err != nil {
		// TODO
		return err
	}

	if err := s.todos.Delete(todo); err != nil {
		// TODO
		return err
	}

	return nil
}

func (s service) GetAllTodos(userID int) ([]models.Todos, error) {
	todos, err := s.todos.GetAll(userID)
	if err != nil {
		// TODO
		return nil, err
	}

	sort.Slice(todos, func(i, j int) bool {
		return todos[i].TimeToComplete.After(todos[j].TimeToComplete)
	})

	return todos, nil
}

func (s service) GetAllCurrentTodos(userID int, t time.Time) ([]models.Todos, error) {
	todos, err := s.todos.GetAll(userID)
	if err != nil {
		return nil, err
	}

	parsedTodos := make([]models.Todos, 0)
	for _, todo := range todos {
		if t.Before(todo.TimeToComplete) {
			parsedTodos = append(parsedTodos, todo)
		}
	}

	sort.Slice(parsedTodos, func(i, j int) bool {
		return parsedTodos[i].TimeToComplete.After(parsedTodos[j].TimeToComplete)
	})

	return parsedTodos, nil
}
