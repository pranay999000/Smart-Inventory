package domain

import userproto "github.com/pranay999000/smart-inventory/user-service/proto"

var SignUpRequest struct {
	Email			string					`json:"email"`
	Password		string					`json:"password"`
	PhoneNumber		string					`json:"phone_number"`
	FirstName		string					`json:"first_name"`
	MiddleName		string					`json:"middle_name"`
	LastName		string					`json:"last_name"`
	Status			userproto.UserStatus	`json:"status"`
	Role			userproto.UserRole		`json:"role"`
	Avatar			string					`json:"avatar"`
	Location		string					`json:"location"`
}

type SignUpResponse struct {
	Success			bool					`json:"success"`
	ErrMessage		interface{}				`json:"err_message"`
}