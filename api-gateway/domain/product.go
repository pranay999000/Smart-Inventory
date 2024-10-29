package domain

var Product struct {
	ID					uint			`json:"id"`
	ProductName			string			`json:"product_name"`
	Image				string			`json:"image"`
	Price				string			`json:"price"`
	Description			string			`json:"description"`
	Category			string			`json:"category"`
	Specification		string			`json:"specification"`
	BusinessId			uint			`json:"business_id"`
}

type ProductResponse struct {
	Success				bool		`json:"success"`
	ErrMessage			interface{}	`json:"err_message"`
	ProductId			interface{}	`json:"product_id"`
}
