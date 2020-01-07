package main

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/jessequinn/microgqlserver/api/internal/gql"
	"github.com/jessequinn/microgqlserver/api/internal/gql/resolvers"
	hello "github.com/jessequinn/microgqlserver/srv/proto/hello"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"
)

func main() {
	// create new web service
	service := web.NewService(
		web.Name("go.micro.api.gqlserver"),
		web.Version("0.1"),
		web.Address(":8000"),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// RPC client
	cl := hello.NewSayService("go.micro.srv.rpcserver", client.DefaultClient)

	// register graphql handlers
	service.Handle("/", handler.Playground("GraphQL playground", "/query"))
	service.Handle("/query", handler.GraphQL(gql.NewExecutableSchema(gql.Config{Resolvers: &resolvers.Resolver{Client: cl}})))

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
