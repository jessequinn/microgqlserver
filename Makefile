.PHONY: proto data build

proto:
	for d in srv; do \
		for f in $$d/**/proto/**/*.proto; do \
			protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. $$f; \
			echo compiled: $$f; \
		done \
	done

gen:
	cd api/auth/gql && go run github.com/99designs/gqlgen -v

lint:
	./bin/lint.sh

build:
	./bin/build.sh

data:
	go-bindata -o data/bindata.go -pkg data data/*.json

run:
	docker-compose up --build --remove-orphans -d
	docker-compose logs -f authsrv authapi
