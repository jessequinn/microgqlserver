FROM golang:1.13.5-alpine as builder

WORKDIR /go/src/github.com/jessequinn/microgqlserver/srv

COPY . .
RUN rm -r scripts .realize.yaml Makefile

ENV GO111MODULE=on

RUN CGO_ENABLED=0 GOOS=linux go build -o rpcserver


FROM alpine:3.11

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/jessequinn/microgqlserver/srv/rpcserver .

CMD ["./rpcserver"]
