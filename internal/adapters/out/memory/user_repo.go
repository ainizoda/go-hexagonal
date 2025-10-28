package memory

import (
	"context"
	"sync"

	"github.com/ainizoda/go-hexagonal/internal/domain/user"
)

type UserRepo struct {
	data map[string]*user.Model
	mu   sync.RWMutex
}

func NewUserRepo() *UserRepo {
	return &UserRepo{data: make(map[string]*user.Model)}
}

func (ur *UserRepo) Select(ctx context.Context, id string) (*user.Model, error) {
	ur.mu.RLock()
	defer ur.mu.RUnlock()
	data, ok := ur.data[id]
	if !ok {
		return nil, user.ErrUserDoesNotExist
	}
	return data, nil
}
func (ur *UserRepo) Save(ctx context.Context, usr *user.Model) error {
	ur.mu.Lock()
	defer ur.mu.Unlock()
	for _, v := range ur.data {
		if v.Email == usr.Email {
			return user.ErrUserAlreadyExists
		}
	}
	ur.data[usr.ID] = usr
	return nil
}
func (ur *UserRepo) Remove(ctx context.Context, id string) error {
	ur.mu.Lock()
	defer ur.mu.Unlock()
	_, ok := ur.data[id]
	if !ok {
		return user.ErrUserDoesNotExist
	}
	delete(ur.data, id)
	return nil
}
func (ur *UserRepo) SelectAll(ctx context.Context) ([]*user.Model, error) {
	ur.mu.RLock()
	defer ur.mu.RUnlock()
	users := make([]*user.Model, 0, 32)
	for _, v := range ur.data {
		users = append(users, v)
	}
	return users, nil
}
