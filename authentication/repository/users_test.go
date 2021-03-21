package repository

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/saravase/golang_grpc_mongo_microservice/authentication/models"
	"github.com/saravase/golang_grpc_mongo_microservice/db"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Panicln(err)
	}

	config := db.NewConfig()
	con, err := db.NewConnection(config)
	if err != nil {
		log.Panicln(err)
	}
	defer con.Close()

	r := NewUsersRepository(con)
	r.DeleteAll()
}

func TestUsersRepositorySave(t *testing.T) {
	config := db.NewConfig()
	con, err := db.NewConnection(config)
	assert.NoError(t, err)
	defer con.Close()

	id := bson.NewObjectId()

	user := &models.User{
		Id:       id,
		Name:     "TEST",
		Email:    fmt.Sprintf("%s@gmail.com", id.Hex()),
		Password: "123456789",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	r := NewUsersRepository(con)
	err = r.Save(user)
	assert.NoError(t, err)

	found, err := r.GetById(id.Hex())
	assert.NoError(t, err)
	assert.NotNil(t, found)
}

func TestUsersRepositoryGetById(t *testing.T) {
	config := db.NewConfig()
	con, err := db.NewConnection(config)
	assert.NoError(t, err)
	defer con.Close()

	id := bson.NewObjectId()

	user := &models.User{
		Id:       id,
		Name:     "TEST",
		Email:    fmt.Sprintf("%s@gmail.com", id.Hex()),
		Password: "123456789",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	r := NewUsersRepository(con)
	err = r.Save(user)
	assert.NoError(t, err)

	found, err := r.GetById(id.Hex())
	assert.NoError(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, user.Id, found.Id)
	assert.Equal(t, user.Name, found.Name)
	assert.Equal(t, user.Email, found.Email)
	assert.Equal(t, user.Password, found.Password)

	found, err = r.GetById(bson.NewObjectId().Hex())
	assert.Error(t, err)
	assert.EqualError(t, mgo.ErrNotFound, err.Error())
	assert.Nil(t, found)
}

func TestUsersRepositoryGetByEmail(t *testing.T) {
	config := db.NewConfig()
	con, err := db.NewConnection(config)
	assert.NoError(t, err)
	defer con.Close()

	id := bson.NewObjectId()

	user := &models.User{
		Id:       id,
		Name:     "TEST",
		Email:    fmt.Sprintf("%s@gmail.com", id.Hex()),
		Password: "123456789",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	r := NewUsersRepository(con)
	err = r.Save(user)
	assert.NoError(t, err)

	found, err := r.GetByEmail(user.Email)
	assert.NoError(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, user.Id, found.Id)
	assert.Equal(t, user.Name, found.Name)
	assert.Equal(t, user.Email, found.Email)
	assert.Equal(t, user.Password, found.Password)

	found, err = r.GetByEmail("")
	assert.Error(t, err)
	assert.EqualError(t, mgo.ErrNotFound, err.Error())
	assert.Nil(t, found)
}

func TestUsersRepositoryUpdate(t *testing.T) {
	config := db.NewConfig()
	con, err := db.NewConnection(config)
	assert.NoError(t, err)
	defer con.Close()

	id := bson.NewObjectId()

	user := &models.User{
		Id:       id,
		Name:     "TEST",
		Email:    fmt.Sprintf("%s@gmail.com", id.Hex()),
		Password: "123456789",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	r := NewUsersRepository(con)
	err = r.Save(user)
	assert.NoError(t, err)

	user.Name = "UPDATE"
	err = r.Update(user)
	assert.NoError(t, err)

	found, err := r.GetById(id.Hex())
	assert.NoError(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, user.Name, found.Name)

}

func TestUsersRepositoryDelete(t *testing.T) {
	config := db.NewConfig()
	con, err := db.NewConnection(config)
	assert.NoError(t, err)
	defer con.Close()

	id := bson.NewObjectId()

	user := &models.User{
		Id:       id,
		Name:     "TEST",
		Email:    fmt.Sprintf("%s@gmail.com", id.Hex()),
		Password: "123456789",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	r := NewUsersRepository(con)
	err = r.Save(user)
	assert.NoError(t, err)

	found, err := r.GetById(id.Hex())
	assert.NoError(t, err)
	assert.NotNil(t, found)

	err = r.Delete(id.Hex())
	assert.NoError(t, err)

	found, err = r.GetById(id.Hex())
	assert.Error(t, err)
	assert.EqualError(t, mgo.ErrNotFound, err.Error())
	assert.Nil(t, found)
}
