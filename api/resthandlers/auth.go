package resthandlers

import "github.com/saravase/golang_grpc_mongo_microservice/pb"

type AuthHandlers interface {
}

type authHandlers struct {
	authSvcClient pb.AuthServiceClient
}

func NewAuthHandlers(authSvcClient pb.AuthServiceClient) AuthHandlers {
	return &authHandlers{
		authSvcClient: authSvcClient,
	}
}
