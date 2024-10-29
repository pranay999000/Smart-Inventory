package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/pranay999000/smart-inventory/api-gateway/domain"
	"github.com/pranay999000/smart-inventory/api-gateway/pkg"
	productproto "github.com/pranay999000/smart-inventory/product-service/proto/product"
)

type ProductHandler struct {
	ProductClient *productproto.ProductServiceClient
}

func NewProductHandler(client *productproto.ProductServiceClient) *ProductHandler {
	return &ProductHandler{
		ProductClient: client,
	}
}

func (h *ProductHandler) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	
	if r.Method != http.MethodPost {
		pkg.WriteJSONResponse(w, &domain.ProductResponse{
			Success: false,
			ErrMessage: "method not allowed",
		}, http.StatusMethodNotAllowed)
		return
	}

	claims, ok := pkg.GetClaims(r)
	if !ok {
		pkg.WriteJSONResponse(w, &domain.ProductResponse{
			Success: false,
			ErrMessage: "userId not found",
		}, http.StatusInternalServerError)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&domain.Product); err != nil {
		pkg.WriteJSONResponse(w, &domain.ProductResponse{
			Success: false,
			ErrMessage: "unable to decode json",
		}, http.StatusInternalServerError)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 3)
	defer cancel()

	res, err := (*h.ProductClient).CreateProduct(ctx, &productproto.CreateProductRequest{
		ProductName: domain.Product.ProductName,
		Image: domain.Product.Image,
		Price: domain.Product.Price,
		Description: domain.Product.Description,
		Specification: domain.Product.Specification,
		Category: domain.Product.Category,
		CreatedBy: claims.UserId,
	})

	if err != nil {
		log.Printf("error creating product: %v\n", err)
		pkg.WriteJSONResponse(w, &domain.ProductResponse{
			Success: false,
			ErrMessage: "error creating product",
		}, http.StatusInternalServerError)
		return
	}

	switch result := res.Result.(type) {
	case *productproto.CreateProductResponse_ErrMessage:
		pkg.WriteJSONResponse(w, &domain.ProductResponse{
			Success: res.Success,
			ErrMessage: result.ErrMessage,
		}, http.StatusInternalServerError)
	case *productproto.CreateProductResponse_ProductId:
		pkg.WriteJSONResponse(w, &domain.ProductResponse{
			Success: res.Success,
			ProductId: result.ProductId,
		}, http.StatusOK)
	}
}