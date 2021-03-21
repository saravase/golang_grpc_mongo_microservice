#!/bin/bash

go clean --cache && go test -v -cover github.com/saravase/golang_grpc_mongo_microservice/authentication/...
go build -o authentication/authsvc authentication/main.go