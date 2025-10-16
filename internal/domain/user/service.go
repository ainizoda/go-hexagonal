package user

import "context"

type service struct {
	repo OutputPort
}

func NewService(repo OutputPort) *service {
	return &service{repo: repo}
}

func (s *service) Get(ctx context.Context, id string) (*Model, error) {
	return s.repo.Select(ctx, id)
}

func (s *service) Add(ctx context.Context, user *Model) error {
	return s.repo.Save(ctx, user)
}

func (s *service) Delete(ctx context.Context, id string) error {
	return s.repo.Remove(ctx, id)
}

func (s *service) List(ctx context.Context) ([]*Model, error) {
	return s.repo.SelectAll(ctx)
}
