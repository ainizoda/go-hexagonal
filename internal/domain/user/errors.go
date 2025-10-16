package user

import "errors"

var (
	ErrUserDoesNotExist  = errors.New("user does not exist")
	ErrUserAlreadyExists = errors.New("user with this email already exists")
)
