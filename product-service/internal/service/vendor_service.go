package service

import (
	"context"
	"strings"

	"github.com/pranay999000/smart-inventory/product-service/internal/domain"
	"github.com/pranay999000/smart-inventory/product-service/internal/repository"
	vendorpb "github.com/pranay999000/smart-inventory/product-service/proto/vendor"
	userproto "github.com/pranay999000/smart-inventory/user-service/proto/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type VendorServiceServer struct {
	vendorpb.UnimplementedVendorServiceServer
	VendorRepo		*repository.VendorRepo
	UserGRPCClient	*userproto.UserServiceClient
}

func NewVendorServiceServer(repo *repository.VendorRepo, userclient *userproto.UserServiceClient) *VendorServiceServer {
	return &VendorServiceServer{
		VendorRepo: repo,
		UserGRPCClient: userclient,
	}
}

func (s *VendorServiceServer) CreateVendor(ctx context.Context, req *vendorpb.CreateVendorRequest) (*vendorpb.CreateVendorResponse, error) {

	if strings.TrimSpace(req.Name) == "" ||
		strings.TrimSpace(req.Location) == "" ||
		strings.TrimSpace(req.UserId) == "" {
		return &vendorpb.CreateVendorResponse{
			Success: false,
			Result: &vendorpb.CreateVendorResponse_ErrMessage{
				ErrMessage: "fields are requied",
			},
		}, status.Error(codes.InvalidArgument, "bad request")
	}

	res, err := (*s.UserGRPCClient).GetBusinessId(ctx, &userproto.GetBusinessIdRequest{
		UserId: req.UserId,
	})
	if err != nil {
		return &vendorpb.CreateVendorResponse{
			Success: false,
			Result: &vendorpb.CreateVendorResponse_ErrMessage{
				ErrMessage: "unable to get businessId",
			},
		}, status.Error(codes.Internal, "unable to find business Id")
	}

	var businessId uint

	switch result := res.Result.(type) {
	case *userproto.GetBusinessIdResponse_ErrMessage:
		return &vendorpb.CreateVendorResponse{
			Success: false,
			Result: &vendorpb.CreateVendorResponse_ErrMessage{
				ErrMessage: "unable to get businessId",
			},
		}, status.Error(codes.Internal, "unable to find business Id")
	case *userproto.GetBusinessIdResponse_BusinessId:
		businessId = uint(result.BusinessId)
	}

	vendor := &domain.Vendor{
		Name: req.Name,
		Location: req.Location,
		BusinessId: businessId,
	}

	vendor_id, err := s.VendorRepo.CreateVendor(ctx, vendor)
	if err != nil {
		return &vendorpb.CreateVendorResponse{
			Success: false,
			Result: &vendorpb.CreateVendorResponse_ErrMessage{
				ErrMessage: "error creating vendor",
			},
		}, status.Errorf(codes.Internal, "error: %v", err)
	}

	return &vendorpb.CreateVendorResponse{
		Success: true,
		Result: &vendorpb.CreateVendorResponse_VendorId{
			VendorId: int32(vendor_id),
		},
	}, nil

}