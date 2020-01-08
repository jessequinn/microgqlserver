package resolvers

import (
	"context"
	"github.com/jessequinn/microgqlserver/api/authapi/gql"
	pb "github.com/jessequinn/microgqlserver/srv/authsrv/proto/auth"
	"log"
	"time"
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
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	res, err := r.Client.GetAll(ctx, &pb.Request{})

	if err != nil {
		return nil, err
	}

	var responses []*pb.Response

	//responses = append(responses, &pb.Response{
	//	Users: res.Users,
	//})

	for _, v := range res.Users {
		log.Println(v.Id)
		responses = append(responses, &pb.Response{
			User: v,
			//User: &pb.User{
			//	Id:       v.Id,
			//	//Name:     v.Name,
			//	//Company:  v.Company,
			//	//Email:    v.Email,
			//	//Password: v.Password,
			//},
		})
	}
	return responses, nil
}

type userResolver struct{ *Resolver }

func (r *userResolver) ID(ctx context.Context, obj *pb.Response) (string, error) {
	panic(obj.User.Id)
	//return obj.User.Id, nil
}

func (r *userResolver) Name(ctx context.Context, obj *pb.Response) (string, error) {
	return obj.User.Name, nil
}

func (r *userResolver) Company(ctx context.Context, obj *pb.Response) (string, error) {
	return obj.User.Company, nil
}

func (r *userResolver) Email(ctx context.Context, obj *pb.Response) (string, error) {
	return obj.User.Email, nil
}

func (r *userResolver) Password(ctx context.Context, obj *pb.Response) (string, error) {
	return obj.User.Password, nil
}
