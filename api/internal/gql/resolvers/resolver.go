package resolvers

import (
	"context"

	"github.com/jessequinn/microgqlserver/api/internal/gql"
	go_micro_srv_microgqlserver "github.com/jessequinn/microgqlserver/srv/proto/hello"
)

type Resolver struct {
	Client go_micro_srv_microgqlserver.SayService
}

func (r *Resolver) Query() gql.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Hello(ctx context.Context, name string) (*go_micro_srv_microgqlserver.Response, error) {
	res, err := r.Client.Hello(ctx, &go_micro_srv_microgqlserver.Request{Name: name})
	if err != nil {
		return &go_micro_srv_microgqlserver.Response{}, err
	}
	return res, nil
}
