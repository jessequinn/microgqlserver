package handlers

import (
	"errors"
	"github.com/globalsign/mgo"
	pb "github.com/jessequinn/microgqlserver/srv/proto/auth"
	rs "github.com/jessequinn/microgqlserver/srv/repositories"
	ss "github.com/jessequinn/microgqlserver/srv/services"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"log"
)

// Our gRPC Service handler
type Service struct {
	Session      *mgo.Session
	TokenService ss.Authable
}

func (srv *Service) GetRepo() rs.Repository {
	return &rs.AuthRepository{srv.Session.Clone()}
}

func (srv *Service) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	repo := srv.GetRepo()
	defer repo.Close()
	user, err := repo.Get(req.Id)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}

func (srv *Service) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	repo := srv.GetRepo()
	defer repo.Close()
	users, err := repo.GetAll()
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}

func (srv *Service) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
	repo := srv.GetRepo()
	defer repo.Close()
	log.Println("Logging in with:", req.Email, req.Password)
	user, err := repo.GetByEmail(req.Email)
	log.Println(user)
	if err != nil {
		return err
	}
	// Compares our given password against the hashed password
	// stored in the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return err
	}
	token, err := srv.TokenService.Encode(user)
	if err != nil {
		return err
	}
	res.Token = token
	return nil
}

func (srv *Service) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
	repo := srv.GetRepo()
	defer repo.Close()
	// Generates a hashed version of our password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(hashedPass)
	if err := repo.Create(req); err != nil {
		return err
	}
	res.User = req
	return nil
}

func (srv *Service) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	// Decode token
	claims, err := srv.TokenService.Decode(req.Token)
	if err != nil {
		return err
	}
	log.Println(claims)
	if claims.User.Id == "" {
		return errors.New("invalid user")
	}
	res.Valid = true
	return nil
}
