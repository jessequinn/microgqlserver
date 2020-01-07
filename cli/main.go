package main

import (
	"context"
	pb "github.com/jessequinn/microgqlserver/srv/proto/auth"
	//pb "github.com/jessequinn/microgqlserver/cli/proto/auth"
	"github.com/micro/go-micro"
	"log"
	"os"
)

func main() {
	// create a new service
	cli := micro.NewService(
		micro.Name("go.micro.cli.user"),
		micro.Version("0.1"),
	)
	// parse command line flags
	cli.Init()
	// Create new greeter client
	client := pb.NewUserService("go.micro.srv.user", cli.Client())
	name := "Ewan Valentine"
	email := "ewan.valentine89@gmail.com"
	password := "test123"
	company := "BBC"
	log.Println(name, email, password)
	r, err := client.Create(context.TODO(), &pb.User{
		Name:     name,
		Email:    email,
		Password: password,
		Company:  company,
	})
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created: %s", r.User.Id)
	getAll, err := client.GetAll(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("Could not list users: %v", err)
	}
	for _, v := range getAll.Users {
		log.Println(v)
	}
	authResponse, err := client.Auth(context.TODO(), &pb.User{
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