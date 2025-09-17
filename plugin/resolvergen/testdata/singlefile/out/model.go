package customresolver

import "context"

type Resolver struct{}
type TestInput struct{}

type QueryResolver interface {
	Resolver(ctx context.Context) (*Resolver, error)
}

type ResolverResolver interface {
	Name(ctx context.Context, obj *Resolver) (string, error)
}

type TestInputResolver interface {
	InputName(ctx context.Context, obj *TestInput, data string) error
}
