package domain

import (
	"github.com/pranay999000/smart-inventory/inventory-service/internal/database"
	"gorm.io/gorm"
)

type Inventory struct {
	gorm.Model
	ID					uint				`json:"id" gorm:"primaryKey;autoIncrement"`
	VendorId			uint				`json:"vendor_id" gorm:"type:int;not null"`
	ProductId			uint				`json:"product_id" gorm:"type:int;not null"`
	BusinessId			uint				`json:"business_id" gorm:"type:int;not null"`
	Total				uint				`json:"quantity" gorm:"type:int;not null"`
	Sold				uint				`json:"sold" gorm:"type:int;not null"`
	Defected			uint				`json:"defected" gorm:"type:int;not null"`
}

func InventoryMigrate() {
	database.WriteDB.AutoMigrate(&Inventory{})
}