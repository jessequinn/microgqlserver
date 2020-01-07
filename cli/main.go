package main

import (
	"context"
	pb "github.com/jessequinn/microgqlserver/srv/proto/auth"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/service/grpc"
	"log"
	"os"
)

func main() {
	// Dummy data
	name := "Jesse Quinn"
	email := "me@jesequinn.info"
	password := "test123"
	company := "CBS"
	// create a new service
	service := grpc.NewService(
		micro.Version("1.0.5"),
	)
	// parse command line flags
	service.Init()
	// Create new greeter client
	client := pb.NewUserService("go.micro.srv.user", service.Client())
	// Set arbitrary headers in context
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "jesse quinn",
		"X-From-Id": "authcli",
	})
	rsp, err := client.Create(ctx, &pb.User{
		Name:     name,
		Email:    email,
		Password: password,
		Company:  company,
	})
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created: %s", rsp.User.Name)
	getAll, err := client.GetAll(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("Could not list users: %v", err)
	}
	for _, v := range getAll.Users {
		log.Println(v.Id)
	}
	authResponse, err := client.Auth(ctx, &pb.User{
		Email:    email,
		Password: password,
	})
	if err != nil {
		log.Fatalf("Could not authenticate user: %s error: %v\n", email, err)
	}
	log.Printf("Your access token is: %s \n", authResponse.Token)
	// let's just exit because
	os.Exit(0)
}
