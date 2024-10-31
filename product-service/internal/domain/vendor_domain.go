package domain

import (
	"github.com/pranay999000/smart-inventory/product-service/internal/database"
	"gorm.io/gorm"
)

type Vendor struct {
	gorm.Model
	ID					uint				`json:"id" gorm:"primaryKey;autoIncrement"`
	BusinessId			uint				`json:"business_id" gorm:"type:int;not null"`
	Name				string				`json:"name" gorm:"type:varchar(40);not null"`
	Location			string				`json:"location" gorm:"type:varchar(40);not null"`
}

func VendorMigrate() {
	database.WriteDB.AutoMigrate(&Vendor{})
}