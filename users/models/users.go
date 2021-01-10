package models

import (
	"github.com/go-pg/pg/v10"
)

//go:generate mockgen -source=$GOFILE -destination=../mocks/model_user_mock.go -package=mocks UserImpl

type UsersImpl interface {
	GetByID(id int) (*Users, error)
	GetByName(name string) (*Users, error)
	Create(user *Users) (*Users, error)
	Update(user *Users) (*Users, error)
}

type Users struct {
	tableName struct{} `pg:"users,alias:u"`
	ID        int
	Username  string
	Password  string
}

type UsersRepo struct {
	db *pg.DB
}

func NewUsersModel(db *pg.DB) *UsersRepo {
	return &UsersRepo{db: db}
}

func (r *UsersRepo) GetByName(username string) (*Users, error) {
	user := &Users{}
	err := r.db.Model(user).
		Where("username = ?", username).
		Select()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UsersRepo) GetByID(id int) (*Users, error) {
	user := &Users{}
	err := r.db.Model(user).
		Where("id = ?", id).
		Select()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UsersRepo) Create(user *Users) (*Users, error) {
	_, err := r.db.Model(user).
		OnConflict("DO NOTHING").
		Insert()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UsersRepo) Update(user *Users) (*Users, error) {
	_, err := r.db.Model(user).WherePK().UpdateNotZero()
	if err != nil {
		return nil, err
	}

	return user, nil
}
