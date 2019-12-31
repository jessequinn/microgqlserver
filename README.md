# Microgqlserver Service

This is the Microgqlserver service

Generated with

```
micro new github.com/jessequinn/microgqlserver --namespace=go.micro --type=api
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.api.microgqlserver
- Type: api
- Alias: microgqlserver

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./microgqlserver-api
```

Build a docker image
```
make docker
```