package service

import (
	"github.com/code7unner/rest-api-test-task/users/internal/auth"
	"github.com/code7unner/rest-api-test-task/users/internal/models"
	"golang.org/x/crypto/bcrypt"
)

//go:generate mockgen -source=$GOFILE -destination=mock/service_mock.go -package=service_mock Service

type Service interface {
	Login(username, password string) (string, error)
	Register(username, password string) (*models.Users, error)
	GetUser(id int) (*models.Users, error)
}

type service struct {
	users     models.UsersImpl
	jwtSecret string
	expires   int
}

func New(usersModel models.UsersImpl, jwtSecret string, expires int) Service {
	return &service{
		users:     usersModel,
		jwtSecret: jwtSecret,
		expires:   expires,
	}
}

func (s service) Login(username, password string) (string, error) {
	user, err := s.users.GetByName(username)
	if err != nil {
		return "", ErrUserNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", ErrUserPasswordInvalid
	}

	user.Password = password
	token, err := auth.CreateToken(user, s.jwtSecret, s.expires)
	if err != nil {
		return "", ErrUserCreateJWTToken
	}

	return token, nil
}

func (s service) Register(username, password string) (*models.Users, error) {
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return nil, ErrUserPasswordInvalid
	}

	newUser := &models.Users{
		Username: username,
		Password: string(hashPwd),
	}

	if _, err := s.users.Create(newUser); err != nil {
		return nil, ErrUserCreating
	}

	return newUser, nil
}

func (s service) GetUser(id int) (*models.Users, error) {
	user, err := s.users.GetByID(id)
	if err != nil {
		return nil, ErrUserNotFound
	}

	return user, nil
}
