package service

import (
	"context"
	"strings"

	"github.com/pranay999000/smart-inventory/user-service/internal/domain"
	"github.com/pranay999000/smart-inventory/user-service/internal/repository"
	businesspb "github.com/pranay999000/smart-inventory/user-service/proto/business"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BusinessServiceServer struct {
	businesspb.UnimplementedBusinessServiceServer
	BusinessRepo	*repository.BusinessRepo
	UserRepo		*repository.UserRepository
}

func NewBusinessServiceServer(businessRepo *repository.BusinessRepo, userRepo *repository.UserRepository) *BusinessServiceServer {
	return &BusinessServiceServer{
		BusinessRepo: businessRepo,
		UserRepo: userRepo,
	}
}

func (s *BusinessServiceServer) CreateBusiness(ctx context.Context, req *businesspb.CreateBusinessRequest) (*businesspb.CreateBusinessResponse, error) {

	if strings.TrimSpace(req.Business.CreatedBy) == "" ||
		strings.TrimSpace(req.Business.BusinessName) == "" ||
		strings.TrimSpace(req.Business.Industry) == "" {
			return nil, status.Error(codes.InvalidArgument, "bad request")
		}

	_, err := s.UserRepo.ValidateUserById(ctx, req.Business.CreatedBy)
	if err != nil {
		return nil, err
	}

	business := &domain.Business{
		BusinessName: req.Business.BusinessName,
		CreatedBy: req.Business.CreatedBy,
		Industry: req.Business.Industry,
		Logo: req.Business.Logo,
		WebsiteUrl: req.Business.WebsiteUrl,
		Description: req.Business.Description,
		Inagurated: int(req.Business.Inagurated),
	}

	businessId, err := s.BusinessRepo.CreateBusiness(ctx, business)
	if err != nil {
		return nil, err
	}

	if err := s.UserRepo.SetBusinessIdAndStatus(ctx, businessId, req.Business.CreatedBy); err != nil {
		return &businesspb.CreateBusinessResponse{
			Success: false,
			Result: &businesspb.CreateBusinessResponse_ErrMessage{
				ErrMessage: &businesspb.ErrorMessage{
					Message: "failed to update user's business id and status",
				},
			},
		}, err
	}

	return &businesspb.CreateBusinessResponse{
		Success: true,
		Result: &businesspb.CreateBusinessResponse_Message{
			Message: "successfully created business",
		},
	}, nil

}