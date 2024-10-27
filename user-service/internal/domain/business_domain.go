package domain

import (
	"github.com/pranay999000/smart-inventory/user-service/internal/database"
	"gorm.io/gorm"
)

type Business struct {
	gorm.Model
	ID					uint			`json:"id" gorm:"primaryKey;autoIncrement"`
	BusinessName		string			`json:"business_name" gorm:"type:varchar(40);not null"`
	CreatedBy			string			`json:"created_by" gorm:"type:varchar(40);not null"`
	Industry			string			`json:"industry" gorm:"type:varchar(40)"`
	Description			string			`json:"description" gorm:"type:text"`
	Logo				string			`json:"logo" gorm:"type:text"`
	Inagurated			int				`json:"inagurated" gorm:"type:int"`
	WebsiteUrl			string			`json:"website_url" gorm:"type:text"`
}

func Migrate() {
	database.WriteDB.AutoMigrate(&Business{})
}