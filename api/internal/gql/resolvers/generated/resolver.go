package resolvers

import (
	"context"

	"github.com/jessequinn/microgqlserver/api/internal/gql"
	go_micro_srv_microgqlserver "github.com/jessequinn/microgqlserver/srv/proto/hello"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Query() gql.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Hello(ctx context.Context, name string) (*go_micro_srv_microgqlserver.Response, error) {
	panic("not implemented")
}
