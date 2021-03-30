#!/bin/bash

cp ../../authentication/authsvc .
cp ../../api/apisvc .

docker build -t saravase/microservices:v1 .
docker inspect saravase/microservices:v1