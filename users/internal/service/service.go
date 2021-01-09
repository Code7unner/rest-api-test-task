package service

import (
	"github.com/code7unner/rest-api-test-task/users/internal/auth"
	"github.com/code7unner/rest-api-test-task/users/models"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(username, password string) (string, error)
	Register(username, password string) error
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
		// TODO
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		// TODO
		return "", err
	}

	user.Password = password
	token, err := auth.CreateToken(user, s.jwtSecret, s.expires)
	if err != nil {
		// TODO
		return "", err
	}

	return token, nil
}

func (s service) Register(username, password string) error {
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		// TODO
		return err
	}

	newUser := &models.Users{
		Username: username,
		Password: string(hashPwd),
	}

	if err := s.users.Create(newUser); err != nil {
		// TODO
		return err
	}

	return nil
}

func (s service) GetUser(id int) (*models.Users, error) {
	user, err := s.users.GetByID(id)
	if err != nil {
		// TODO
		return nil, err
	}

	return user, nil
}
