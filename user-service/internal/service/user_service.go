package service

import (
	"context"
	"strings"

	"github.com/pranay999000/smart-inventory/user-service/internal/domain"
	"github.com/pranay999000/smart-inventory/user-service/internal/pkg"
	"github.com/pranay999000/smart-inventory/user-service/internal/repository"
	userpb "github.com/pranay999000/smart-inventory/user-service/proto/user"
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

func (s *UserServiceServer) SignIn(ctx context.Context, req *userpb.SignInRequest) (*userpb.SignInResponse, error) {

	if strings.TrimSpace(req.Email) == "" || strings.TrimSpace(req.Password) == "" {
		return &userpb.SignInResponse{
			Success: false,
			Result: &userpb.SignInResponse_ErrMessage{
				ErrMessage: &userpb.ErrorMessage{
					Message: "some fields are required",
				},
			},
		}, status.Error(codes.InvalidArgument, "bad request")
	}

	var user *domain.User

	doesExist, user, err := s.Repository.CheckEmail(ctx, req.Email)
	if err != nil {
		return &userpb.SignInResponse{
			Success: false,
			Result: &userpb.SignInResponse_ErrMessage{
				ErrMessage: &userpb.ErrorMessage{
					Message: "error in finding email",
				},
			},
		}, status.Errorf(codes.InvalidArgument, "unable to find email: %s", req.Email)
	}
	
	if !doesExist {
		return &userpb.SignInResponse{
			Success: false,
			Result: &userpb.SignInResponse_ErrMessage{
				ErrMessage: &userpb.ErrorMessage{
					Message: "email does not exist",
				},
			},
		}, status.Errorf(codes.InvalidArgument, "invalid email: %s", req.Email)
	}

	passwordMatch := pkg.CheckPasswordHash(req.Password, user.Password)

	if !passwordMatch {
		return &userpb.SignInResponse{
			Success: false,
			Result: &userpb.SignInResponse_ErrMessage{
				ErrMessage: &userpb.ErrorMessage{
					Message: "password does not match",
				},
			},
		}, status.Error(codes.InvalidArgument, "invalid password",)
	}

	tk, err := pkg.GenerateJWTToken(user.ID, user.FirstName, user.MiddleName, user.LastName, user.Avatar)

	if err != nil {
		return &userpb.SignInResponse{
			Success: false,
			Result: &userpb.SignInResponse_ErrMessage{
				ErrMessage: &userpb.ErrorMessage{
					Message: "unable to generate auth token",
				},
			},
		}, status.Errorf(codes.Internal, "unable to generate auth_tk: %v", err)
	}

	err = s.Repository.SetUserStatus(ctx, user.ID, domain.ACTIVE)
	if err != nil {
		return &userpb.SignInResponse{
			Success: false,
			Result: &userpb.SignInResponse_ErrMessage{
				ErrMessage: &userpb.ErrorMessage{
					Message: "unable to update user status",
				},
			},
		}, status.Errorf(codes.Internal, "unable to update user status: %v", err)
	}

	return &userpb.SignInResponse{
		Success: true,
		Result: &userpb.SignInResponse_AuthTk{
			AuthTk: &userpb.AuthTK{
				UserId: user.ID,
				Token: tk,
			},
		},
	}, nil
	
}

func (s *UserServiceServer) GetBusinessId(ctx context.Context, req *userpb.GetBusinessIdRequest) (*userpb.GetBusinessIdResponse, error) {

	if strings.TrimSpace(req.UserId) == "" {
		return &userpb.GetBusinessIdResponse{
			Success: false,
			Result: &userpb.GetBusinessIdResponse_ErrMessage{
				ErrMessage: "fields are required",
			},
		}, status.Error(codes.InvalidArgument, "bad request")
	}

	user, err := s.Repository.ValidateUserById(ctx, req.UserId)
	if err != nil {
		return &userpb.GetBusinessIdResponse{
			Success: false,
			Result: &userpb.GetBusinessIdResponse_ErrMessage{
				ErrMessage: "unable to find user",
			},
		}, status.Error(codes.Internal, "unable to find user")
	}

	return &userpb.GetBusinessIdResponse{
		Success: true,
		Result: &userpb.GetBusinessIdResponse_BusinessId{
			BusinessId: int32(user.BusinessId),
		},
	}, nil

}