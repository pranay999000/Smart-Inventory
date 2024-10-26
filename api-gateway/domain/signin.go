package domain

var SignInRequest struct {
	Email			string					`json:"email"`
	Password		string					`json:"password"`
}

type SignInResponse struct {
	Success			bool					`json:"success"`
	ErrMessage		interface{}				`json:"err_message"`
	AuthTK			interface{}				`json:"auth_tk"`
}