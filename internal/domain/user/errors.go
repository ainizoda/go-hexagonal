package user

import (
	"errors"
)

var (
	ErrUserDoesNotExist  = errors.New("user does not exist")
	ErrUserAlreadyExists = errors.New("user with this email already exists")
	ErrEmptyField        = errors.New("field should not be empty")
	ErrInvalidEmail      = errors.New("provided email is invalid")
)
