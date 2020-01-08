FROM alpine:3.11

RUN apk --no-cache add ca-certificates

ADD . /app

WORKDIR /app

ENTRYPOINT [ "/app/authcli" ]

