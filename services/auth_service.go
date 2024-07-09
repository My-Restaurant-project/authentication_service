package services

import (
	"context"
	"fmt"

	pb "authentication_service/genproto/authentication_service"
	"authentication_service/repositories"
)

type AuthService struct {
	pb.UnimplementedAuthenticationServiceServer
	repo *repositories.UserRepository
}

func NewAuthService(db *repositories.UserRepository) *AuthService {
	return &AuthService{repo: db}
}

func (s *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	fmt.Println(req.Password)
	user, err := s.repo.Login(req)

	if err != nil {
		return nil, err
	}
	return user, nil

}

func (s *AuthService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user, err := s.repo.Register(req)
	if err != nil {
		return nil, err
	}
	return user, nil

}

func (s *AuthService) GetProfileById(ctx context.Context, UserId *pb.UserIdRequest) (*pb.UserIdResponse, error) {
	user, err := s.repo.GetProfileById(UserId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
