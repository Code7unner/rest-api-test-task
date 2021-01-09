package service

import "errors"

var (
	// ErrUserNotFound user with given id not found.
	ErrUserNotFound = errors.New("user not found")
	// ErrUserPasswordInvalid user password could not encrypted.
	ErrUserPasswordInvalid = errors.New("user password is invalid")
	// ErrUserCreateJWTToken user token could not created
	ErrUserCreateJWTToken = errors.New("could not create user jwt-token")
	// ErrUserCreating create user in db is fail
	ErrUserCreating = errors.New("could not create user")
)
