package validators

import (
	"errors"
	"strings"

	"github.com/saravase/golang_grpc_mongo_microservice/pb"
	"gopkg.in/mgo.v2/bson"
)

var (
	ErrInvalidUserId      = errors.New("invalid user id")
	ErrEmptyName          = errors.New("name can't be empty")
	ErrEmptyEmail         = errors.New("email can't be empty")
	ErrEmptyPassword      = errors.New("password can't be empty")
	ErrEmailAlreadyExists = errors.New("email already exists")
)

func ValidateSignUp(u *pb.User) error {
	if !bson.IsObjectIdHex(u.Id) {
		return ErrInvalidUserId
	}
	if u.Name != "" {
		return ErrEmptyName
	}
	if u.Email != "" {
		return ErrEmptyEmail
	}
	if u.Password != "" {
		return ErrEmptyPassword
	}
	return nil
}

func NormalizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}
