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

## Petstore API

You can use the publicly available petstore API : https://petstore3.swagger.io

To deploy an instance of petstore yourself, please follow readme on [swagger-petstore](https://github.com/swagger-api/swagger-petstore) GitHub repository.

## Build

### Regenerate petstore library

The petstore package was automatically generated from the openapi specifications of the petstore API.

Here is the command to regenerate the petstore package :

```shell
oapi-codegen -generate types,client -package petstore https://petstore3.swagger.io/api/v3/openapi.json > petstore/petstore.gen.go
```

### Run application

To run application, execute `go run` command :

```shell
go run *.go
```

### Build executable

To build an executable, execute `go build` command :

```shell
go build -o petstore.bin *.go
```
