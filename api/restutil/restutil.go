package restutil

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/saravase/golang_grpc_mongo_microservice/security"
)

var (
	ErrBodyEmpty    = errors.New("body can't be empty")
	ErrUnauthorized = errors.New("unauthorized user")
)

type JError struct {
	Error string `json:"error"`
}

func WriteAsJson(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(data)
}

func WriteError(w http.ResponseWriter, statusCode int, err error) {
	e := "error"
	if err != nil {
		e = err.Error()
	}
	WriteAsJson(w, statusCode, JError{e})
}

func AuthRequestWithId(r *http.Request) (*security.TokenPayload, error) {
	tokenString, err := security.ExtractToken(r)
	if err != nil {
		return nil, err
	}

	payload, err := security.NewTokenPayload(tokenString)
	if err != nil {
		return nil, err
	}

	vars := mux.Vars(r)
	if payload.UserId != vars["id"] {
		return nil, ErrUnauthorized
	}

	return payload, nil
}
