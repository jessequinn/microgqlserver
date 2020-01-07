TAG     := $$(git rev-parse --short HEAD)
GOPATH	:= $(shell go env GOPATH)

.PHONY: build
build:
	go build -o authcli main.go

.PHONY: run
run:
	go run main.go --registry=mdns

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	@docker build -t authcli:prod .
	@docker tag authcli:prod authcli:${TAG}
	@docker tag authcli:prod authcli:latest

