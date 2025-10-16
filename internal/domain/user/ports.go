package user

import "context"

type OutputPort interface {
	Select(ctx context.Context, id string) (*Model, error)
	Save(ctx context.Context, user *Model) error
	Remove(ctx context.Context, id string) error
	SelectAll(ctx context.Context) ([]*Model, error)
}

type InputPort interface {
	Get(ctx context.Context, id string) (*Model, error)
	Add(ctx context.Context, user *Model) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*Model, error)
}
