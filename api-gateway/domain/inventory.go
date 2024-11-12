package domain

var InventoryRequest struct {
	BusinessId		uint		`json:"business_id"`
	VendorId		uint		`json:"vendor_id"`
	ProductId		uint		`json:"product_id"`
}

type InventoryResponse struct {
	Success			bool		`json:"success"`
	ErrMessage		interface{}	`json:"err_messsage"`
	InventoryId		uint		`json:"inventory_id"`
}