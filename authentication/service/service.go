package service

import (
	"context"
	"strings"
	"time"

	"github.com/saravase/golang_grpc_mongo_microservice/authentication/models"
	"github.com/saravase/golang_grpc_mongo_microservice/authentication/repository"
	"github.com/saravase/golang_grpc_mongo_microservice/authentication/validators"
	"github.com/saravase/golang_grpc_mongo_microservice/pb"
	"github.com/saravase/golang_grpc_mongo_microservice/security"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type AuthService interface {
	SignUp(ctx context.Context, req *pb.User) (*pb.User, error)
	SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error)
	GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error)
	ListUsers(req *pb.ListUserRequest, stream pb.AuthService_ListUsersServer) error
	UpdateUser(ctx context.Context, req *pb.User) (*pb.User, error)
	DeleteUser(ctx context.Context, req *pb.GetUserRequest) (*pb.DeleteUserResponse, error)
}

type authService struct {
	usersRepository repository.UsersRepository
}

func NewAuthService(usersRepository repository.UsersRepository) pb.AuthServiceServer {
	return &authService{
		usersRepository: usersRepository,
	}
}

func (s *authService) SignUp(ctx context.Context, req *pb.User) (*pb.User, error) {
	err := validators.ValidateSignUp(req)
	if err != nil {
		return nil, err
	}
	req.Name = strings.TrimSpace(req.Name)
	req.Email = validators.NormalizeEmail(req.Email)
	req.Password, err = security.EncryptPassword(req.Password)
	if err != nil {
		return nil, err
	}

	found, err := s.usersRepository.GetByEmail(req.Email)
	if err == mgo.ErrNotFound {
		user := new(models.User)
		user.FromProtoBuffer(req)
		err = s.usersRepository.Save(user)
		if err != nil {
			return nil, err
		}
		return user.ToProtoBuffer(), nil
	}

	if found == nil {
		return nil, err
	}

	return nil, validators.ErrEmailAlreadyExists
}

func (s *authService) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	email := validators.NormalizeEmail(req.Email)
	user, err := s.usersRepository.GetByEmail(email)
	if err != nil {
		return nil, validators.ErrSignInFailed
	}

	err = security.VerifyPassword(user.Password, req.Password)
	if err != nil {
		return nil, validators.ErrSignInFailed
	}

	token, err := security.NewToken(user.Id.Hex())
	if err != nil {
		return nil, validators.ErrSignInFailed
	}

	return &pb.SignInResponse{
		User:  user.ToProtoBuffer(),
		Token: token,
	}, nil
}

func (s *authService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	found, err := s.usersRepository.GetById(req.Id)
	if err != nil {
		return nil, err
	}

	return found.ToProtoBuffer(), nil
}

func (s *authService) ListUsers(req *pb.ListUserRequest, stream pb.AuthService_ListUsersServer) error {
	users, err := s.usersRepository.GetAll()
	if err != nil {
		return err
	}

	for _, user := range users {
		err := stream.Send(user.ToProtoBuffer())
		if err != nil {
			return err
		}
	}

	return nil
}
func (s *authService) UpdateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	if !bson.IsObjectIdHex(req.Id) {
		return nil, validators.ErrInvalidUserId
	}

	user, err := s.usersRepository.GetById(req.Id)
	if err != nil {
		return nil, err
	}

	if user.Name == req.Name {
		return user.ToProtoBuffer(), nil
	}

	user.Name = req.Name
	user.Updated = time.Now()
	err = s.usersRepository.Update(user)
	if err != nil {
		return nil, err
	}

	return user.ToProtoBuffer(), nil
}

func (s *authService) DeleteUser(ctx context.Context, req *pb.GetUserRequest) (*pb.DeleteUserResponse, error) {
	if !bson.IsObjectIdHex(req.Id) {
		return nil, validators.ErrInvalidUserId
	}

	err := s.usersRepository.Delete(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteUserResponse{Id: req.Id}, nil
}
