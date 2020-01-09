package main

import (
	"github.com/globalsign/mgo"
	ds "github.com/jessequinn/microgqlserver/srv/authsrv/datastores"
	hs "github.com/jessequinn/microgqlserver/srv/authsrv/handlers"
	pb "github.com/jessequinn/microgqlserver/srv/authsrv/proto/auth"
	rs "github.com/jessequinn/microgqlserver/srv/authsrv/repositories"
	ss "github.com/jessequinn/microgqlserver/srv/authsrv/services"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"log"
	"os"
	"time"
)

// https://docs.mongodb.com/manual/reference/limits/#restrictions-on-db-names
// https://stackoverflow.com/questions/5916080/what-are-naming-conventions-for-mongodb
const (
	dbName          = "service"
	usersCollection = "users"
	defaultHost     = "localhost:27017"
)

// Testing purposes only
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
	// Index
	index := mgo.Index{
		Key:        []string{"email"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	c := session.DB(dbName).C(usersCollection)
	err = c.EnsureIndex(index)
	if err != nil {
		log.Fatalf("Error creating unique index: %v", err)
	}

	repo := &rs.AuthRepository{session.Copy()}
	tokenService := &ss.TokenService{repo}
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	service := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("1.0.6"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
	service.Init()
	service.Server().Init(
		server.Wait(nil),
	)
	pb.RegisterUserServiceHandler(service.Server(), &hs.Service{session, tokenService})
	if err := service.Run(); err != nil {
		log.Fatalf("Error running server: %v", err)
	}
}
