package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/pranay999000/smart-inventory/api-gateway/domain"
	"github.com/pranay999000/smart-inventory/api-gateway/pkg"
	vendorproto "github.com/pranay999000/smart-inventory/product-service/proto/vendor"
)

type VendorHandler struct {
	VendorClient	*vendorproto.VendorServiceClient
}

func NewVendorHandler(client *vendorproto.VendorServiceClient) *VendorHandler {
	return &VendorHandler{
		VendorClient: client,
	}
}

func (h *VendorHandler) CreateVendorFunc(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		pkg.WriteJSONResponse(w, &domain.VendorResponse{
			Success: false,
			ErrMessage: "method not allowed",
		}, http.StatusMethodNotAllowed)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 3)
	defer cancel()

	if err := json.NewDecoder(r.Body).Decode(&domain.Vendor); err != nil {
		pkg.WriteJSONResponse(w, &domain.VendorResponse{
			Success: false,
			ErrMessage: "unable to decode json request",
		}, http.StatusInternalServerError)
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

	res, err := (*h.VendorClient).CreateVendor(ctx, &vendorproto.CreateVendorRequest{
		Name: domain.Vendor.Name,
		Location: domain.Vendor.Location,
		UserId: claims.UserId,
	})

	if err != nil {
		log.Printf("Create Vandor: %v\n", err)
		pkg.WriteJSONResponse(w, &domain.VendorResponse{
			Success: false,
			ErrMessage: "unable to create vendor",
		}, http.StatusInternalServerError)
		return
	}

	switch result := res.Result.(type) {
	case *vendorproto.CreateVendorResponse_ErrMessage:
		pkg.WriteJSONResponse(w, domain.VendorResponse{
			Success: false,
			ErrMessage: result.ErrMessage,
		}, http.StatusInternalServerError)
	case *vendorproto.CreateVendorResponse_VendorId:
		pkg.WriteJSONResponse(w, domain.VendorResponse{
			Success: true,
			VendorId: uint(result.VendorId),
		}, http.StatusOK)
	}

}