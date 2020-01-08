package resolvers

import (
	"context"

	"github.com/jessequinn/microgqlserver/api/authapi/gql"
	go_micro_srv_user "github.com/jessequinn/microgqlserver/srv/authsrv/proto/auth"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Query() gql.QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) User() gql.UserResolver {
	return &userResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]*go_micro_srv_user.GetUserResponse, error) {
	panic("not implemented")
}

type userResolver struct{ *Resolver }

func (r *userResolver) ID(ctx context.Context, obj *go_micro_srv_user.GetUserResponse) (string, error) {
	panic("not implemented")
}
func (r *userResolver) Name(ctx context.Context, obj *go_micro_srv_user.GetUserResponse) (string, error) {
	panic("not implemented")
}
func (r *userResolver) Company(ctx context.Context, obj *go_micro_srv_user.GetUserResponse) (string, error) {
	panic("not implemented")
}
func (r *userResolver) Email(ctx context.Context, obj *go_micro_srv_user.GetUserResponse) (string, error) {
	panic("not implemented")
}
func (r *userResolver) Password(ctx context.Context, obj *go_micro_srv_user.GetUserResponse) (string, error) {
	panic("not implemented")
}