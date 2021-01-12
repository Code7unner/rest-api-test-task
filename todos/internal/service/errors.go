package service

import "errors"

var (
	// ErrTodoNotCreated could not create todo
	ErrTodoNotCreated = errors.New("todo not created")
	// ErrTodoNotUpdated could not update todo
	ErrTodoNotUpdated = errors.New("todo not updated")
	// ErrTodoNotFound could not find todo
	ErrTodoNotFound = errors.New("todo not found")
	// ErrTodoNotDeleted could not delete todo
	ErrTodoNotDeleted = errors.New("todo not deleted")
	// ErrTodosNotFound could not find todos
	ErrTodosNotFound = errors.New("todos not found")
)
