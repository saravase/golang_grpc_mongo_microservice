package routes

import (
	"net/http"

	"github.com/saravase/golang_grpc_mongo_microservice/api/resthandlers"
)

func NewAuthRoutes(authHandlers resthandlers.AuthHandlers) []*Route {
	return []*Route{
		{
			Path:    "/signup",
			Method:  http.MethodPost,
			Handler: authHandlers.SignUp,
		},
		{
			Path:    "/users",
			Method:  http.MethodGet,
			Handler: authHandlers.GetUsers,
		},
		{
			Path:    "/users/{id}",
			Method:  http.MethodGet,
			Handler: authHandlers.GetUser,
		},
		{
			Path:    "/users/{id}",
			Method:  http.MethodPut,
			Handler: authHandlers.PutUser,
		},
		{
			Path:    "/users/{id}",
			Method:  http.MethodDelete,
			Handler: authHandlers.DeleteUser,
		},
	}
}
