package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/joho/godotenv"
	"github.com/saravase/golang_grpc_mongo_microservice/authentication/repository"
	"github.com/saravase/golang_grpc_mongo_microservice/authentication/service"
	"github.com/saravase/golang_grpc_mongo_microservice/db"
	"github.com/saravase/golang_grpc_mongo_microservice/pb"
	"google.golang.org/grpc"
)

var (
	local bool
	port  int
)

func init() {
	flag.IntVar(&port, "port", 9001, "authentication service port")
	flag.BoolVar(&local, "local", true, "run authentication service local")
	flag.Parse()
}

func main() {
	if local {
		err := godotenv.Load()
		if err != nil {
			log.Panicln(err)
		}
	}

	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	if err != nil {
		log.Panicln(err)
	}
	defer conn.Close()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	usersRepository := repository.NewUsersRepository(conn)
	authService := service.NewAuthService(usersRepository)

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, authService)

	log.Printf("Authentication service running on [::]:%d\n", port)

	grpcServer.Serve(lis)
}
