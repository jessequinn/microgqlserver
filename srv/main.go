package main

import (
	ds "github.com/jessequinn/microgqlserver/srv/datastores"
	hs "github.com/jessequinn/microgqlserver/srv/handlers"
	pb "github.com/jessequinn/microgqlserver/srv/proto/vessel"
	rs "github.com/jessequinn/microgqlserver/srv/repositories"
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
	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "Kane's Salty Secret", MaxWeight: 200000, Capacity: 500},
	}
	for _, v := range vessels {
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
	repo := &rs.VesselRepository{session.Copy()}
	createDummyData(repo)
	// Configure 'log' package to give file name and line number on eg. log.Fatal
	// just the filename & line number:
	// log.SetFlags(log.Lshortfile)
	// Or add timestamps and pipe file name and line number to it:
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	srv := micro.NewService(
		micro.Name("go.micro.srv.rpcserver"),
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
	pb.RegisterVesselServiceHandler(srv.Server(), &hs.Service{session})
	// Run server
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
