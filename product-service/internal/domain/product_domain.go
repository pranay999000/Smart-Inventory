package domain

import (
	"github.com/pranay999000/smart-inventory/product-service/internal/database"
	"gorm.io/gorm"
)


type Product struct {
	gorm.Model
	ID					uint			`json:"id" gorm:"primaryKey;autoIncrement"`
	ProductName			string			`json:"product_name" gorm:"type:varchar(40);not null"`
	Image				string			`json:"image" gorm:"type:text"`
	Price				string			`json:"price" gorm:"type:varchar(12);not null"`
	Description			string			`json:"description" gorm:"type:text"`
	Category			string			`json:"category" gorm:"type:varchar(24);not null"`
	Specification		string			`json:"specification" gorm:"type:text"`
	BusinessId			uint			`json:"business_id" gorm:"type:int;not null"`
	CreatedBy			string			`json:"created_by" gorm:"type:varchar(40);not null"`
}

func Migrate() {
	database.WriteDB.AutoMigrate(&Product{})
}