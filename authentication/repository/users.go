package repository

import (
	"golang_grpc_mongo_microservice/db"

	"gopkg.in/mgo.v2"
)

const userCollection = "users"

type UserRepository interface {
}

type userRepository struct {
	c *mgo.Collection
}

func NewUserRepository(conn db.Connection) UserRepository {
	return &userRepository{c: conn.DB().C(userCollection)}
}
