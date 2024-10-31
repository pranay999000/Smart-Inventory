package domain

var Vendor struct {
	ID			uint		`json:"id"`
	BusinessId	uint		`json:"business_id"`
	Name		string		`json:"name"`
	Location	string		`json:"location"`
}

type VendorResponse struct {
	Success		bool		`json:"success"`
	ErrMessage	interface{}	`json:"err_message"`
	VendorId	uint		`json:"vendor_id"`
}