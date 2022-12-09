# Petstore-go-client

Example of a go client consuming a swagger api.

## prerequsites

### Golang

To compile this program you need to install Go: https://go.dev/dl/

### Oapi-codegen

To automatically regenerate the petstore package from the openapi specs, you will need oapi-codegen:

```shell
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
```

## Deploy petstore swagger instance

To deploy a petstore API, please follow readme on [swagger-petstore](https://github.com/swagger-api/swagger-petstore) GitHub repository.

## Build

### Regenerate petstore library

The petstore package was automatically generated from the openapi specifications of the petstore API.

Here is the command to regenerate the petstore package :

```shell
oapi-codegen -generate types,client -package petstore http://localhost:8080/api/v3/openapi.json > petstore/petstore.gen.go
```

### Run application

To build application, execute `go run` command :

```shell
go run *.go
```

### Build executable

To build an executable, execute `go build` command :

```shell
go build -o petstore.bin *.go
```
