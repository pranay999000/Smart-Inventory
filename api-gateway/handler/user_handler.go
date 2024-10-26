package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/pranay999000/smart-inventory/api-gateway/domain"
	"github.com/pranay999000/smart-inventory/api-gateway/pkg"
	userproto "github.com/pranay999000/smart-inventory/user-service/proto"
)

type UserHandler struct {
	UserClient		*userproto.UserServiceClient
}

func NewUserHandler(client *userproto.UserServiceClient) *UserHandler {
	return &UserHandler{
		UserClient: client,
	}
}

func (h *UserHandler) SignUpHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		pkg.WriteJSONResponse(w, domain.SignUpResponse{
			Success: false,
			ErrMessage: "method not allowed",
		}, http.StatusMethodNotAllowed)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 3)
	defer cancel()

	if err := json.NewDecoder(r.Body).Decode(&domain.SignUpRequest); err != nil {
		log.Printf("error decoding json request: %v", err)
		pkg.WriteJSONResponse(w, domain.SignUpResponse{
			Success: false,
			ErrMessage: "error decoding json request",
		}, http.StatusInternalServerError)
		return
	}

	res, err := (*h.UserClient).SignUp(ctx, &userproto.SignUpUserRequest{
		User: &userproto.User{
			FirstName: domain.SignUpRequest.FirstName,
			MiddleName: domain.SignUpRequest.MiddleName,
			LastName: domain.SignUpRequest.LastName,
			Email: domain.SignUpRequest.Email,
			Password: domain.SignUpRequest.Password,
			PhoneNumber: domain.SignUpRequest.PhoneNumber,
			Avatar: domain.SignUpRequest.Avatar,
			Location: domain.SignUpRequest.Location,
			Status: domain.SignUpRequest.Status,
			Role: domain.SignUpRequest.Role,
		},
	})

	if err != nil {
		log.Printf("error signing up user: %v\n", err)
		pkg.WriteJSONResponse(w, domain.SignUpResponse{
			Success: false,
			ErrMessage: "signup error",
		}, http.StatusInternalServerError)
		return
	}

	pkg.WriteJSONResponse(w, domain.SignUpResponse{
		Success: res.Success,
		ErrMessage: res.ErrMessage,
	}, http.StatusOK)
}

func (h *UserHandler) SigninHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		pkg.WriteJSONResponse(w, domain.SignInResponse{
			Success: false,
			ErrMessage: "method not allowed",
		}, http.StatusMethodNotAllowed)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 3)
	defer cancel()

	if err := json.NewDecoder(r.Body).Decode(&domain.SignInRequest); err != nil {
		log.Printf("error decoding json request: %v", err)
		pkg.WriteJSONResponse(w, domain.SignInResponse{
			Success: false,
			ErrMessage: "error decoding json request",
		}, http.StatusInternalServerError)
		return
	}

	res, err := (*h.UserClient).SignIn(ctx, &userproto.SignInRequest{
		Email: domain.SignInRequest.Email,
		Password: domain.SignInRequest.Password,
	})
	if err != nil {
		log.Printf("error signing in user: %v\n", err)
		pkg.WriteJSONResponse(w, domain.SignInResponse{
			Success: false,
			ErrMessage: "unable to signin user",
		}, http.StatusInternalServerError)
		return
	}

	switch result := res.Result.(type) {
	case *userproto.SignInResponse_ErrMessage:
		log.Printf("An Error Occured: %v\n", err)
		pkg.WriteJSONResponse(w, domain.SignInResponse{
			Success: res.Success,
			ErrMessage: result.ErrMessage.Message,
		}, http.StatusInternalServerError)
	case *userproto.SignInResponse_AuthTk:
		pkg.WriteJSONResponse(w, domain.SignInResponse{
			Success: res.Success,
			AuthTK: result.AuthTk,
		}, http.StatusOK)
	}

}