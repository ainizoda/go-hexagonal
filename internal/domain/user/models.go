package user

import (
	"fmt"
	"strings"

	"github.com/ainizoda/go-hexagonal/internal/utils"
	"github.com/google/uuid"
)

type Model struct {
	ID        string   `json:"id"`
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Email     string   `json:"email"`
	Roles     []string `json:"roles"`
}

func New(firstName, lastName, email string, roles []string) (*Model, error) {
	if err := validateRequiredFields(firstName, lastName, email); err != nil {
		return nil, err
	}
	if ok := utils.IsValidEmail(email); !ok {
		return nil, fmt.Errorf("%w: %s", ErrInvalidEmail, email)
	}
	if len(roles) == 0 {
		return nil, fmt.Errorf("user should have at least one role")
	}
	return &Model{
		ID:        uuid.NewString(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Roles:     roles,
	}, nil
}

func validateRequiredFields(firstName, lastName, email string) error {
	fields := map[string]string{
		"firstName": firstName,
		"lastName":  lastName,
		"email":     email,
	}

	for field, value := range fields {
		if strings.TrimSpace(value) == "" {
			return fmt.Errorf("%s: %w", field, ErrEmptyField)
		}
	}
	return nil
}
