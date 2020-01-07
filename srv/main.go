package main

import (
	ds "github.com/jessequinn/microgqlserver/srv/datastores"
	hs "github.com/jessequinn/microgqlserver/srv/handlers"
	pb "github.com/jessequinn/microgqlserver/srv/proto/auth"
	rs "github.com/jessequinn/microgqlserver/srv/repositories"
	ss "github.com/jessequinn/microgqlserver/srv/services"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"log"
	"os"
	"time"
)

const (
	defaultHost = "localhost:27017"
)

func createDummyData(repo rs.Repository) {
	defer repo.Close()
	users := []*pb.User{
		{Name: "Dummy Name", Company: "Dummy Company", Email: "dummy@dummy.com", Password: "dummy"},
	}
	for _, v := range users {
		repo.Create(v)
	}
}

func main() {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = defaultHost
	}
	session, err := ds.CreateSession(host)
	defer session.Close()
	if err != nil {
		log.Fatalf("Error connecting to datastore: %v", err)
	}
	repo := &rs.AuthRepository{session.Copy()}
	tokenService := &ss.TokenService{repo}
	createDummyData(repo)
	// Configure 'log' package to give file name and line number on eg. log.Fatal
	// just the filename & line number:
	// log.SetFlags(log.Lshortfile)
	// Or add timestamps and pipe file name and line number to it:
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("0.1"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
	// optionally setup command line usage
	srv.Init()
	// Graceful shutdown
	srv.Server().Init(
		server.Wait(nil),
	)
	// Register Handlers
	pb.RegisterUserServiceHandler(srv.Server(), &hs.Service{session, tokenService})
	// Run server
	if err := srv.Run(); err != nil {
		log.Fatalf("Error running server: %v", err)
	}
}
