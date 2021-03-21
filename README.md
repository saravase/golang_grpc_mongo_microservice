# golang_grpc_mongo_microservice

$ go version

    go version go1.16 linux/amd64

$ docker version

    Client: Docker Engine - Community
    Version:           20.10.5
    API version:       1.41
    Go version:        go1.13.15
    Git commit:        55c4c88
    Built:             Tue Mar  2 20:18:20 2021
    OS/Arch:           linux/amd64
    Context:           default
    Experimental:      true
    Got permission denied while trying to connect to the Docker daemon socket at unix:///var/run/docker.sock: Get http://%2Fvar%2Frun%2Fdocker.sock/v1.24/version: dial unix /var/run/docker.sock: connect: permission denied

$ minikube version

    minikube version: v1.18.1
    commit: 09ee84d530de4a92f00f1c5dbc34cead092b95bc

$ kubectl version

    Client Version: version.Info{Major:"1", Minor:"20", GitVersion:"v1.20.4", GitCommit:"e87da0bd6e03ec3fea7933c4b5263d151aafd07c", GitTreeState:"clean", BuildDate:"2021-02-18T16:12:00Z", GoVersion:"go1.15.8", Compiler:"gc", Platform:"linux/amd64"}
    The connection to the server localhost:8080 was refused - did you specify the right host or port?

$ sudo docker run -it --rm --name mongo_container -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=admin -v mongodata:/data/db -d -p 27017:27017 mongo

    b384d43582ac96e5257f5961358ac190c08f41a1980f94dd29ce2ad6280b70b8

$ sudo docker ps

    CONTAINER ID   IMAGE     COMMAND                  CREATED          STATUS          PORTS                      NAMES
    b384d43582ac   mongo     "docker-entrypoint.s…"   29 seconds ago   Up 14 seconds   0.0.0.0:27017->27017/tcp   mongo_container

$ sudo docker exec -it mongo_container /bin/bash

    root@b384d43582ac:/# mongo -u admin -p admin --authenticationDatabase admin
    
    > show dbs;

    admin   0.000GB
    config  0.000GB
    local   0.000GB

    > use microservice;
    
    switched to db microservice

    > db.createUser({'user': 'user','pwd': 'pass','roles': [{ 'role': 'readWrite', 'db': 'microservice'}]});

        Successfully added user: {
	    "user" : "user",
	    "roles" : [
		    {
			    "role" : "readWrite",
			    "db" : "microservice"
		    }
	    ]
    }

    root@b384d43582ac:/# mongo -u user -p pass --authenticationDatabase microservice
    
    MongoDB shell version v4.4.4
    connecting to: mongodb://127.0.0.1:27017/?authSource=microservice&compressors=disabled&gssapiServiceName=mongodb
    Implicit session: session { "id" : UUID("0ffd77a1-9882-486e-8d45-3f407ce32377") }
    MongoDB server version: 4.4.4
    
    > use microservice
    
    switched to db microservice
    
    > show dbs;
    > show collections;

### Install Protocol Buffer Compiler:
    
    $ apt install -y protobuf-compiler
    $ protoc --version  # Ensure compiler version is 3+

### Install Protocol Buffer Generate Plugin:

    $ go get -u github.com/golang/protobuf/protoc-gen-go

### Get gRPC package:

    $ go get google.golang.org/grpc

### Add GO_PATH in ~/.bashrc:

    $ code ~/.bashrc

        export GOROOT=/usr/local/go
        export GOPATH=$HOME/go
        export GOBIN=$GOPATH/bin
        export PATH=$PATH:$GOROOT:$GOPATH:$GOBIN

    $ source ~/.bashrc

### Generate code using protoc command:

    $ protoc -I=./messages ./messages/*.proto --go_out=plugins=grpc:.
