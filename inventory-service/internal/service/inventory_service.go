package service

import (
	"context"
	"strings"

	"github.com/pranay999000/smart-inventory/inventory-service/internal/domain"
	"github.com/pranay999000/smart-inventory/inventory-service/internal/repository"
	inventorypb "github.com/pranay999000/smart-inventory/inventory-service/proto/inventory"
	productproto "github.com/pranay999000/smart-inventory/product-service/proto/product"
	userproto "github.com/pranay999000/smart-inventory/user-service/proto/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type InventoryServiceServer struct {
	inventorypb.InventoryServiceServer
	InventoryRepo 	*repository.InventoryRepo
	UserGRPCClient	*userproto.UserServiceClient
	ProductGRPCClient	*productproto.ProductServiceClient
}

func NewInventoryServiceServer(repo *repository.InventoryRepo, userclient *userproto.UserServiceClient, productclient *productproto.ProductServiceClient) *InventoryServiceServer {
	return &InventoryServiceServer{
		InventoryRepo: repo,
		UserGRPCClient: userclient,
		ProductGRPCClient: productclient,
	}
}

func (s *InventoryServiceServer) CreateInventory(ctx context.Context, req *inventorypb.CreateInventoryRequest) (*inventorypb.CreateInventoryResponse, error) {

	// Validation
	if req.ProductId == 0 ||
		req.VendorId == 0 ||
		strings.TrimSpace(req.UserId) == "" {
		return &inventorypb.CreateInventoryResponse{
			Success: false,
			Result: &inventorypb.CreateInventoryResponse_ErrMessage{
				ErrMessage: "bad request",
			},
		}, status.Error(codes.InvalidArgument, "bad request")
	}

	// Get Business Id
	res, err := (*s.UserGRPCClient).GetBusinessId(ctx, &userproto.GetBusinessIdRequest{
		UserId: req.UserId,
	})
	if err != nil {
		return &inventorypb.CreateInventoryResponse{
			Success: false,
			Result: &inventorypb.CreateInventoryResponse_ErrMessage{
				ErrMessage: "unable to get businessId",
			},
		}, status.Error(codes.Internal, "unable to find business Id")
	}

	
	var businessId uint
	
	switch result := res.Result.(type) {
	case *userproto.GetBusinessIdResponse_ErrMessage:
		return &inventorypb.CreateInventoryResponse{
			Success: false,
			Result: &inventorypb.CreateInventoryResponse_ErrMessage{
				ErrMessage: "unable to get businessId",
			},
		}, status.Error(codes.Internal, "unable to find business Id")
	case *userproto.GetBusinessIdResponse_BusinessId:
		businessId = uint(result.BusinessId)
	}

	// Check Product
	checkProduct, err := (*s.ProductGRPCClient).CheckProduct(ctx, &productproto.CheckProductRequest{
		ProductId: req.ProductId,
	})
	if err != nil {
		return &inventorypb.CreateInventoryResponse{
			Success: false,
			Result: &inventorypb.CreateInventoryResponse_ErrMessage{
				ErrMessage: "unable to find product",
			},
		}, status.Error(codes.Internal, "unable to find product")
	}

	if !checkProduct.Success {
		return &inventorypb.CreateInventoryResponse{
			Success: false,
			Result: &inventorypb.CreateInventoryResponse_ErrMessage{
				ErrMessage: "unable to find product",
			},
		}, status.Error(codes.Internal, "unable to find product")
	}

	// Check if Product is already present in the vendor 
	err = s.InventoryRepo.CheckInventory(ctx, uint(req.ProductId), uint(req.VendorId))
	if err != nil {
		return &inventorypb.CreateInventoryResponse{
			Success: false,
			Result: &inventorypb.CreateInventoryResponse_ErrMessage{
				ErrMessage: "product already present",
			},
		}, status.Error(codes.InvalidArgument, "bad request")
	}

	// Create Inventory
	inventoryId, err := s.InventoryRepo.CreateInventory(ctx, &domain.Inventory{
		BusinessId: businessId,
		ProductId: uint(req.ProductId),
		VendorId: uint(req.VendorId),
	})
	if err != nil {
		return &inventorypb.CreateInventoryResponse{
			Success: false,
			Result: &inventorypb.CreateInventoryResponse_ErrMessage{
				ErrMessage: "unable to create inventory",
			},
		}, status.Error(codes.InvalidArgument, "error creating inventory")
	}

	return &inventorypb.CreateInventoryResponse{
		Success: true,
		Result: &inventorypb.CreateInventoryResponse_InventoryId{
			InventoryId: int32(inventoryId),
		},
	}, nil

}