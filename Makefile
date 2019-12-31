GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. internal/proto/*/*.proto

.PHONY: build
build: proto

	go build -o build/microgqlserver cmd/microgqlserver/main.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t microgqlserver:latest
