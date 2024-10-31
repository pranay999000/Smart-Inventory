package service

import (
	"context"
	"strings"

	"github.com/pranay999000/smart-inventory/product-service/internal/domain"
	"github.com/pranay999000/smart-inventory/product-service/internal/repository"
	vendorpb "github.com/pranay999000/smart-inventory/product-service/proto/vendor"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type VendorServiceServer struct {
	vendorpb.UnimplementedVendorServiceServer
	VendorRepo		*repository.VendorRepo
}

func NewVendorServiceServer(repo *repository.VendorRepo) *VendorServiceServer {
	return &VendorServiceServer{
		VendorRepo: repo,
	}
}

func (s *VendorServiceServer) CreateVendor(ctx context.Context, req *vendorpb.CreateVendorRequest) (*vendorpb.CreateVendorResponse, error) {

	if strings.TrimSpace(req.Name) == "" ||
		strings.TrimSpace(req.Location) == "" ||
		req.BusinessId == 0 {
		return &vendorpb.CreateVendorResponse{
			Success: false,
			Result: &vendorpb.CreateVendorResponse_ErrMessage{
				ErrMessage: "fields are requied",
			},
		}, status.Error(codes.InvalidArgument, "bad request")
	}

	vendor := &domain.Vendor{
		Name: req.Name,
		Location: req.Location,
		BusinessId: uint(req.BusinessId),
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