package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/pranay999000/smart-inventory/api-gateway/domain"
	"github.com/pranay999000/smart-inventory/api-gateway/pkg"
	businessproto "github.com/pranay999000/smart-inventory/user-service/proto/business"
)

type BusinessHandler struct {
	BusinessClient		*businessproto.BusinessServiceClient
}

func NewBusinessHandler(client *businessproto.BusinessServiceClient) *BusinessHandler {
	return &BusinessHandler{
		BusinessClient: client,
	}
}

func (h *BusinessHandler) CreateBusinessHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		pkg.WriteJSONResponse(w, domain.BusinessResponse{
			Success: false,
			ErrMessage: "method not allowed",
		}, http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 3)
	defer cancel()

	if err := json.NewDecoder(r.Body).Decode(&domain.Business); err != nil {
		log.Printf("error decoding json request: %v", err)
		pkg.WriteJSONResponse(w, domain.SignUpResponse{
			Success: false,
			ErrMessage: "error decoding json request",
		}, http.StatusInternalServerError)
		return
	}

	res, err := (*h.BusinessClient).CreateBusiness(ctx, &businessproto.CreateBusinessRequest{
		Business: &businessproto.Business{
			BusinessName: domain.Business.BusinessName,
			CreatedBy: domain.Business.CreatedBy,
			Industry: domain.Business.Industry,
			Inagurated: int32(domain.Business.Inagurated),
			Logo: domain.Business.Logo,
			WebsiteUrl: domain.Business.WebsiteUrl,
			Description: domain.Business.Description,
		},
	})

	if err != nil {
		log.Printf("error creating business: %v\n", err)
		pkg.WriteJSONResponse(w, domain.BusinessResponse{
			Success: res.Success,
			ErrMessage: "error creating business",
		}, http.StatusInternalServerError)
		return
	}

	switch result := res.Result.(type) {
	case *businessproto.CreateBusinessResponse_ErrMessage:
		log.Printf("An Error Occured: %v\n", err)
		pkg.WriteJSONResponse(w, domain.BusinessResponse{
			Success: res.Success,
			ErrMessage: result.ErrMessage.Message,
		}, http.StatusInternalServerError)
	case *businessproto.CreateBusinessResponse_Business:
		pkg.WriteJSONResponse(w, domain.BusinessResponse{
			Success: res.Success,
			Business: result.Business,
		}, http.StatusOK)
	}

}