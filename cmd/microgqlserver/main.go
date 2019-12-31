package main

import (
	"github.com/99designs/gqlgen/handler"
	helloProto "github.com/jessequinn/microgqlserver/internal/proto"
	gql "github.com/micro/examples/greeter/api/graphql/graphql"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"
)

func main() {
	// create new web service
	service := web.NewService(
		web.Name("go.micro.api.microgqlserver"),
		web.Version("0.1"),
		web.Address(":8085"),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// RPC client
	cl := helloProto.NewSayService("go.micro.srv.greeter", client.DefaultClient)

	// register graphql handlers
	service.Handle("/", handler.Playground("GraphQL playground", "/query"))
	service.Handle("/query", handler.GraphQL(gql.NewExecutableSchema(gql.Config{Resolvers: &gql.Resolver{Client: cl}})))

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)

	}
}
