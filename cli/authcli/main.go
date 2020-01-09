package main

import (
	"context"
	pb "github.com/jessequinn/microgqlserver/srv/authsrv/proto/auth"
	"github.com/micro/go-micro"
	"log"
	"os"
)

func main() {
	// Dummy data
	name := "Jesse Quinn"
	email := "me5@jesequinn.info"
	password := "test123"
	company := "CBS"
	// create a new service
	service := micro.NewService(
		micro.Version("1.0.6"),
	)
	// parse command line flags
	service.Init()
	// Create new greeter client
	client := pb.NewUserService("go.micro.srv.user", service.Client())
	// Set arbitrary headers in context
	//ctx := metadata.NewContext(context.Background(), map[string]string{
	//	"X-User-Id": "jesse quinn",
	//	"X-From-Id": "authcli",
	//})
	//rsp, err := client.Create(ctx, &pb.User{
	rsp, err := client.Create(context.Background(), &pb.User{
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
	authResponse, err := client.Auth(context.Background(), &pb.User{
		Email:    email,
		Password: password,
	})
	if err != nil {
		log.Fatalf("Could not authenticate user: %s error: %v\n", email, err)
	}
	log.Printf("Your access token is: %s \n", authResponse.Token)
	validResponse, err := client.ValidateToken(context.Background(), &pb.Token{
		Token: authResponse.Token,
	})
	if err != nil {
		log.Fatalf("Could not validate token: %v\n", err)
	}
	log.Printf("Your token is valid: %v \n", validResponse.Valid)
	// let's just exit because
	os.Exit(0)
}
