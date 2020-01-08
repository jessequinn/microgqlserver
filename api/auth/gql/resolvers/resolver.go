package resolvers

import (
	"context"
	"github.com/jessequinn/microgqlserver/api/auth/gql"
	pb "github.com/jessequinn/microgqlserver/srv/auth/proto/auth"
)

type Resolver struct {
	Client pb.UserService
}

func (r *Resolver) Query() gql.QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) User() gql.UserResolver {
	return &userResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]*pb.Response, error) {
	res, err := r.Client.GetAll(ctx, &pb.Request{})
	if err != nil {
		return &pb.Response{}, err
	}
	return res, nil
}

type userResolver struct{ *Resolver }

func (r *userResolver) ID(ctx context.Context, obj *pb.Response) (string, error) {
	panic("not implemented")
}

func (r *userResolver) Name(ctx context.Context, obj *pb.Response) (string, error) {
	panic("not implemented")
}

func (r *userResolver) Company(ctx context.Context, obj *pb.Response) (string, error) {
	panic("not implemented")
}

func (r *userResolver) Email(ctx context.Context, obj *pb.Response) (string, error) {
	panic("not implemented")
}

func (r *userResolver) Password(ctx context.Context, obj *pb.Response) (string, error) {
	panic("not implemented")
}
