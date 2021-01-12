package models

import (
	"github.com/go-pg/pg/v10"
	"time"
)

//go:generate mockgen -source=$GOFILE -destination=../mocks/model_user_mock.go -package=mocks UserImpl

type TodosImpl interface {
	Get(id int) (*Todos, error)
	GetAll(userID int) ([]Todos, error)
	Create(todo *Todos) (*Todos, error)
	Update(todo *Todos) (*Todos, error)
	Delete(todo *Todos) error
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

func (r TodosRepo) Get(id int) (*Todos, error) {
	todo := &Todos{}
	if err := r.db.Model(todo).Where("id = ?", id).Select(); err != nil {
		return nil, err
	}

	return todo, nil
}

func (r TodosRepo) GetAll(userID int) ([]Todos, error) {
	todos := make([]Todos, 0)
	err := r.db.Model(&todos).Where("user_id = ?", userID).Select()
	if err != nil {
		return nil, err
	}

	return todos, nil
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

func (r TodosRepo) Delete(todo *Todos) error {
	_, err := r.db.Model(todo).WherePK().Delete()

	return err
}