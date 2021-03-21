package main

import (
	"flag"
	"log"

	"github.com/joho/godotenv"
	"github.com/saravase/golang_grpc_mongo_microservice/db"
)

var (
	local bool
)

func init() {
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
}
