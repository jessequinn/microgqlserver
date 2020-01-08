This Golang project contains the following structure (at the moment)

```sh
.
├── api
│   └── auth
│       ├── Dockerfile
│       ├── gql
│       │   ├── generated.go
│       │   ├── gqlgen.yml
│       │   ├── resolvers
│       │   │   ├── generated
│       │   │   │   └── resolver.go
│       │   │   └── resolver.go
│       │   └── schema.graphql
│       └── main.go
├── bin
│   ├── build.sh
│   └── lint.sh
├── cli
│   └── authcli
│       ├── authcli
│       ├── Dockerfile
│       └── main.go
├── docker
│   ├── app
│   │   └── dev.Dockerfile
│   ├── docker
│   │   └── consule
│   │       ├── config
│   │       │   └── consul-conf.json
│   │       └── Dockerfile
│   └── mongo
│       └── root
│           └── 000_init_replicaSet.js
├── docker-compose.yml
├── docker.sh
├── go.mod
├── go.sum
├── Makefile
├── README.md
└── srv
    └── authsrv
        ├── authsrv
        ├── datastores
        │   └── datastore.go
        ├── Dockerfile
        ├── handlers
        │   └── handler.go
        ├── main.go
        ├── proto
        │   └── auth
        │       ├── auth.pb.go
        │       ├── auth.pb.micro.go
        │       └── auth.proto
        ├── repositories
        │   └── repository.go
        └── services
            └── token.go
```

To build the protobuffs

    $ make proto

To build the graphql dependencies

    $ make gen
    
To build the apps

    $ make build

To run docker-compose and apps

    $ make run
    
For discovery, `etcd` is used. Therefore to talk to a service the following is required

    $ ./authcli --registry=etcd --registry_address=http://localhost:2379 api # an example

This command is ran from the `cli` directory if not obvious.

To connect to the GraphQL playground `http://localhost:8000/`
