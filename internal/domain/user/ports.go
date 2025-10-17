package user

import "context"

/*
http, grpc, or any client that needs to
get access to domain, will use this port
*/
type InputPort interface {
	Get(ctx context.Context, id string) (*Model, error)
	Add(ctx context.Context, user *Model) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*Model, error)
}

/*
data-store adapters like: redis, postgres
or anything else should implement this port
*/
type OutputPort interface {
	Select(ctx context.Context, id string) (*Model, error)
	Save(ctx context.Context, user *Model) error
	Remove(ctx context.Context, id string) error
	SelectAll(ctx context.Context) ([]*Model, error)
}
