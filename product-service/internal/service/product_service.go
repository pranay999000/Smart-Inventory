package service

import (
	"context"
	"strings"

	"github.com/pranay999000/smart-inventory/product-service/internal/domain"
	"github.com/pranay999000/smart-inventory/product-service/internal/repository"
	productpb "github.com/pranay999000/smart-inventory/product-service/proto/product"
	userproto "github.com/pranay999000/smart-inventory/user-service/proto/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductServiceServer struct {
	productpb.UnimplementedProductServiceServer
	ProductRepo		*repository.ProductRepo
	UserGRPCClient	*userproto.UserServiceClient
}

func NewProductService(productRepo *repository.ProductRepo, userClient *userproto.UserServiceClient) *ProductServiceServer {
	return &ProductServiceServer{
		ProductRepo: productRepo,
		UserGRPCClient: userClient,
	}
}

func (s *ProductServiceServer) CreateProduct(ctx context.Context, req *productpb.CreateProductRequest) (*productpb.CreateProductResponse, error) {

	if strings.TrimSpace(req.ProductName) == "" ||
		strings.TrimSpace(req.Image) == "" ||
		strings.TrimSpace(req.Category) == "" ||
		strings.TrimSpace(req.Description) == "" ||
		strings.TrimSpace(req.Price) == "" ||
		strings.TrimSpace(req.Specification) == "" ||
		strings.TrimSpace(req.CreatedBy) == "" {
		return &productpb.CreateProductResponse{
			Success: false,
			Result: &productpb.CreateProductResponse_ErrMessage{
				ErrMessage: "fields are required",
			},
		}, status.Error(codes.InvalidArgument, "bad request")
	}

	res, err := (*s.UserGRPCClient).GetBusinessId(ctx, &userproto.GetBusinessIdRequest{
		UserId: req.CreatedBy,
	})
	if err != nil {
		return &productpb.CreateProductResponse{
			Success: false,
			Result: &productpb.CreateProductResponse_ErrMessage{
				ErrMessage: "unable to get businessId",
			},
		}, status.Error(codes.Internal, "unable to find business Id")
	}

	var businessId uint

	switch result := res.Result.(type) {
	case *userproto.GetBusinessIdResponse_ErrMessage:
		return &productpb.CreateProductResponse{
			Success: false,
			Result: &productpb.CreateProductResponse_ErrMessage{
				ErrMessage: "unable to get businessId",
			},
		}, status.Error(codes.Internal, "unable to find business Id")
	case *userproto.GetBusinessIdResponse_BusinessId:
		businessId = uint(result.BusinessId)
	}

	product := &domain.Product{
		ProductName: req.ProductName,
		Price: req.Price,
		Image: req.Image,
		Category: req.Category,
		Description: req.Description,
		Specification: req.Specification,
		CreatedBy: req.CreatedBy,
		BusinessId: businessId,
	}

	productId, err := s.ProductRepo.CreateProduct(ctx, product)
	if err != nil {
		return &productpb.CreateProductResponse{
			Success: false,
			Result: &productpb.CreateProductResponse_ErrMessage{
				ErrMessage: "Unable to create product",
			},
		}, status.Errorf(codes.Internal, "Error while creating product: %v", err)
	}

	return &productpb.CreateProductResponse{
		Success: true,
		Result: &productpb.CreateProductResponse_ProductId{
			ProductId: int32(productId),
		},
	}, nil

}

func (s *ProductServiceServer) CheckProduct(ctx context.Context, req *productpb.CheckProductRequest) (*productpb.CheckProductResponse, error) {

	err := s.ProductRepo.CheckProduct(ctx, uint(req.ProductId))

	if err != nil {
		return nil, err
	}

	return &productpb.CheckProductResponse{
		Success: true,
	}, nil

}