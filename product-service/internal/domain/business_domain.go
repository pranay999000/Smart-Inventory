package domain

type Business struct {
	ID					int32			`json:"id"`
	BusinessName		string			`json:"business_name"`
	CreatedBy			string			`json:"created_by"`
	Industry			string			`json:"industry"`
	Description			string			`json:"description"`
	Logo				string			`json:"logo"`
	Inagurated			int				`json:"inagurated"`
	WebsiteUrl			string			`json:"website_url"`
}