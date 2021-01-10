package models

import (
	"github.com/go-pg/pg/v10"
	"time"
)

//go:generate mockgen -source=$GOFILE -destination=../mocks/model_user_mock.go -package=mocks UserImpl

type TodosImpl interface {
	Create(todo *Todos) (*Todos, error)
	Update(todo *Todos) (*Todos, error)
}

type Todos struct {
	tableName      struct{} `pg:"todos,alias:t"`
	ID             int
	UserID         int
	Title          string
	Description    string
	TimeToComplete time.Time
}

type TodosRepo struct {
	db *pg.DB
}

func NewTodosModel(db *pg.DB) *TodosRepo {
	return &TodosRepo{db: db}
}

func (r TodosRepo) Create(todo *Todos) (*Todos, error) {
	_, err := r.db.Model(todo).Insert()
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (r TodosRepo) Update(todo *Todos) (*Todos, error) {
	_, err := r.db.Model(todo).WherePK().UpdateNotZero()
	if err != nil {
		return nil, err
	}

	return todo, nil
}
