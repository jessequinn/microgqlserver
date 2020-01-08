package main

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/jessequinn/microgqlserver/api/authapi/gql"
	rs "github.com/jessequinn/microgqlserver/api/authapi/gql/resolvers"
	pb "github.com/jessequinn/microgqlserver/srv/authsrv/proto/auth"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"
)

func main() {
	service := web.NewService(
		web.Name("go.micro.api.user"),
		web.Version("1.0.6"),
		web.Address(":8000"),
	)
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}
	client := pb.NewUserService("go.micro.srv.user", client.DefaultClient)
	service.Handle("/", handler.Playground("GraphQL playground", "/query"))
	service.Handle("/query", handler.GraphQL(gql.NewExecutableSchema(gql.Config{Resolvers: &rs.Resolver{Client: client}})))
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
