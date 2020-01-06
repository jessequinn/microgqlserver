TAG     := $$(git rev-parse --short HEAD)
GOPATH	:=$(shell go env GOPATH)

.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/vessel/*.proto

.PHONY: build
build: proto

	go build -o rpcserver main.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	@docker build -t rpcserver:prod .
	@docker tag rpcserver:prod rpcserver:${TAG}
	@docker tag rpcserver:prod rpcserver:latest

