package service

import (
	"context"
	"strings"

	"github.com/pranay999000/smart-inventory/user-service/internal/domain"
	"github.com/pranay999000/smart-inventory/user-service/internal/repository"
	userpb "github.com/pranay999000/smart-inventory/user-service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServiceServer struct {
	Repository *repository.UserRepository
	userpb.UnimplementedUserServiceServer
}

func NewUserService(repo *repository.UserRepository) *UserServiceServer {
	return &UserServiceServer{
		Repository: repo,
	}
}

func (s *UserServiceServer) SignUp(ctx context.Context, req *userpb.SignUpUserRequest) (*userpb.SignUpUserResponse, error) {

	if strings.TrimSpace(req.User.FirstName) == "" ||
		strings.TrimSpace(req.User.MiddleName) == "" ||
		strings.TrimSpace(req.User.LastName) == "" ||
		strings.TrimSpace(req.User.Email) == "" ||
		strings.TrimSpace(req.User.PhoneNumber) == "" {
		return &userpb.SignUpUserResponse{
			Success: false,
			ErrMessage: &userpb.ErrorMessage{
				Message: "some fields are required",
			},
		}, status.Error(codes.InvalidArgument, "some fields are required")
	}

	newUser := domain.User{
		FirstName: req.User.FirstName,
		MiddleName: req.User.MiddleName,
		LastName: req.User.LastName,
		Email: req.User.Email,
		PhoneNumber: req.User.PhoneNumber,
		Status: domain.Status(req.User.Status),
		Role: domain.Role(req.User.Role),
		Avatar: req.User.Avatar,
		Location: req.User.Location,
		Password: req.User.Password,
	}

	err := s.Repository.CreateUser(ctx, &newUser)
	if err != nil {
		return &userpb.SignUpUserResponse{
			Success: false,
			ErrMessage: &userpb.ErrorMessage{
				Message: err.Error(),
			},
		}, status.Errorf(codes.InvalidArgument, "Error occured while creating user: %v", err)
	}

	return &userpb.SignUpUserResponse{
		Success: true,
		ErrMessage: nil,
	}, nil

}