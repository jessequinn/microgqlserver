package resolvers

import (
	"context"
	"github.com/jessequinn/microgqlserver/api/authapi/gql"
	pb "github.com/jessequinn/microgqlserver/srv/authsrv/proto/auth"
	"time"
)

type Resolver struct {
	Client pb.UserService
}

func (r *Resolver) Mutation() gql.MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Query() gql.QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) User() gql.UserResolver {
	return &userResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, user pb.User) (*pb.GetUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	res, err := r.Client.Create(ctx, &user)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserResponse{
		User: &pb.User{
			Id:       res.User.Id,
			Name:     res.User.Name,
			Company:  res.User.Company,
			Email:    res.User.Email,
			Password: res.User.Password,
		},
	}, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]*pb.GetUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	res, err := r.Client.GetAll(ctx, &pb.Request{})
	if err != nil {
		return nil, err
	}
	var responses []*pb.GetUserResponse
	for _, v := range res.Users {
		response := &pb.GetUserResponse{
			User: &pb.User{
				Id:       v.Id,
				Name:     v.Name,
				Company:  v.Company,
				Email:    v.Email,
				Password: v.Password,
			},
		}
		responses = append(responses, response)
	}
	return responses, nil
}

type userResolver struct{ *Resolver }

func (r *userResolver) ID(ctx context.Context, obj *pb.GetUserResponse) (string, error) {
	return obj.User.Id, nil
}

func (r *userResolver) Name(ctx context.Context, obj *pb.GetUserResponse) (string, error) {
	return obj.User.Name, nil
}

func (r *userResolver) Company(ctx context.Context, obj *pb.GetUserResponse) (string, error) {
	return obj.User.Company, nil
}

func (r *userResolver) Email(ctx context.Context, obj *pb.GetUserResponse) (string, error) {
	return obj.User.Email, nil
}

func (r *userResolver) Password(ctx context.Context, obj *pb.GetUserResponse) (string, error) {
	return obj.User.Password, nil
}
