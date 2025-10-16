package memory

import (
	"context"

	"github.com/ainizoda/go-hexagonal/internal/domain/user"
)

type userRepo struct {
	data map[string]*user.Model
}

func NewUserRepo() *userRepo {
	return &userRepo{data: make(map[string]*user.Model)}
}

func (ur *userRepo) Select(ctx context.Context, id string) (*user.Model, error) {
	data, ok := ur.data[id]
	if !ok {
		return nil, user.ErrUserDoesNotExist
	}
	return data, nil
}
func (ur *userRepo) Save(ctx context.Context, usr *user.Model) error {
	for _, v := range ur.data {
		if v.Email == usr.Email {
			return user.ErrUserAlreadyExists
		}
	}
	ur.data[usr.ID] = usr
	return nil
}
func (ur *userRepo) Remove(ctx context.Context, id string) error {
	_, ok := ur.data[id]
	if !ok {
		return user.ErrUserDoesNotExist
	}
	delete(ur.data, id)
	return nil
}
func (ur *userRepo) SelectAll(ctx context.Context) ([]*user.Model, error) {
	users := make([]*user.Model, 0, 32)
	for _, v := range ur.data {
		users = append(users, v)
	}
	return users, nil
}
