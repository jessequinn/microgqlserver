package main

import (
	"context"
	pb "github.com/jessequinn/microgqlserver/srv/proto/auth"
	"github.com/micro/go-micro"
	//"time"
	//microclient "github.com/micro/go-micro/client"
	//"github.com/micro/go-micro/config/cmd"
	"log"
	"os"
)

func main() {
	//cmd.Init()
	// Create new client
	//client := pb.NewUserService("go.micro.srv.user", microclient.DefaultClient)
	// create a new service
	cli := micro.NewService()
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
	log.Printf("Created: %s", r.User.XId)
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
