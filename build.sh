#!/bin/bash

go clean --cache && go test -v -cover github.com/saravase/golang_grpc_mongo_microservice/...
go build -o authentication/authsvc authentication/main.go
go build -o api/apisvc api/main.go