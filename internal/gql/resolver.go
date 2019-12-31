package gql

import (
	"context"

	go_micro_srv_greeter "github.com/micro/examples/greeter/srv/proto/hello"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Hello(ctx context.Context, name string) (*go_micro_srv_greeter.Response, error) {
	panic("not implemented")
}
