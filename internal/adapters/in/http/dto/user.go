package dto

import "github.com/ainizoda/go-hexagonal/internal/domain/user"

type UserRequestBody struct {
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Email     string   `json: "email"`
	Roles     []string `json:"roles"`
}

type UserResponseBody struct {
	Data []*user.Model `json:"data"`
}
