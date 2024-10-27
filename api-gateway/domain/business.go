package domain

var Business struct {
	ID					uint		`json:"id"`
	BusinessName		string		`json:"business_name"`
	CreatedBy			string		`json:"created_by"`
	Industry			string		`json:"industry"`
	Description			string		`json:"description"`
	Logo				string		`json:"logo"`
	Inagurated			int			`json:"inagurated"`
	WebsiteUrl			string		`json:"website_url"`
}

type BusinessResponse struct {
	Success				bool		`json:"success"`
	ErrMessage			interface{}	`json:"err_message"`
}