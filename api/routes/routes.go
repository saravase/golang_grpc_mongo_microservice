package routes

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/saravase/golang_grpc_mongo_microservice/api/middlewares"
)

type Route struct {
	Method       string
	Path         string
	Handler      http.HandlerFunc
	AuthRequired bool
}

func Install(router *mux.Router, routeList []*Route) {
	for _, route := range routeList {
		if route.AuthRequired {
			router.HandleFunc(route.Path,
				middlewares.LogRequests(middlewares.Authenticate(route.Handler))).Methods(route.Method)
		} else {
			router.HandleFunc(route.Path, middlewares.LogRequests(route.Handler)).Methods(route.Method)
		}
	}
}

func WithCORS(router *mux.Router) http.Handler {
	headers := handlers.AllowedHeaders([]string{
		"X-Requested-with",
		"Content-Type",
		"Accept",
		"Authorization",
	})
	origins := handlers.AllowedOrigins([]string{
		"*",
	})
	methods := handlers.AllowedMethods([]string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodDelete,
	})

	return handlers.CORS(headers, origins, methods)(router)
}
