package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/pranay999000/smart-inventory/api-gateway/domain"
	"github.com/pranay999000/smart-inventory/api-gateway/pkg"
	inventoryproto "github.com/pranay999000/smart-inventory/inventory-service/proto/inventory"
)

type InventoryHandler struct {
	InventoryClient		*inventoryproto.InventoryServiceClient
}

func NewInventoryHandler(handler *inventoryproto.InventoryServiceClient) *InventoryHandler {
	return &InventoryHandler{
		InventoryClient: handler,
	}
}

func (h *InventoryHandler) CreateInventoryFunc(w http.ResponseWriter, r *http.Request) {
	
	if r.Method != http.MethodPost {
		pkg.WriteJSONResponse(w, &domain.InventoryResponse{
			Success: false,
			ErrMessage: "method not allowed",
		}, http.StatusMethodNotAllowed)
		return
	}

	claims, ok := pkg.GetClaims(r)
	if !ok {
		pkg.WriteJSONResponse(w, &domain.VendorResponse{
			Success: false,
			ErrMessage: "userId not found",
		}, http.StatusInternalServerError)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 3)
	defer cancel()

	if err := json.NewDecoder(r.Body).Decode(&domain.InventoryRequest); err != nil {
		pkg.WriteJSONResponse(w, &domain.InventoryResponse{
			Success: false,
			ErrMessage: "unable to decode json request",
		}, http.StatusInternalServerError)
		return
	}

	res, err := (*h.InventoryClient).CreateInventory(ctx, &inventoryproto.CreateInventoryRequest{
		VendorId: int32(domain.InventoryRequest.VendorId),
		ProductId: int32(domain.InventoryRequest.ProductId),
		UserId: claims.UserId,
	})

	if err != nil {
		log.Printf("Create Inventory: %v\n", err)
		pkg.WriteJSONResponse(w, &domain.InventoryResponse{
			Success: false,
			ErrMessage: "unable to create inventory",
		}, http.StatusInternalServerError)
		return
	}

	switch result := res.Result.(type) {
	case *inventoryproto.CreateInventoryResponse_ErrMessage:
		pkg.WriteJSONResponse(w, domain.InventoryResponse{
			Success: false,
			ErrMessage: result.ErrMessage,
		}, http.StatusInternalServerError)
	case *inventoryproto.CreateInventoryResponse_InventoryId:
		pkg.WriteJSONResponse(w, domain.InventoryResponse{
			Success: true,
			InventoryId: uint(result.InventoryId),
		}, http.StatusOK)
	}

}