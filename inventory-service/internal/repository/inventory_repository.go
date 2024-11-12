package repository

import (
	"context"
	"errors"
	"log"

	"github.com/pranay999000/smart-inventory/inventory-service/internal/domain"
	"gorm.io/gorm"
)

type InventoryRepo struct {
	ReadDB		*gorm.DB
	WriteDB		*gorm.DB
}

func NewInventoryRepo(readDB *gorm.DB, writeDB *gorm.DB) *InventoryRepo {
	return &InventoryRepo{
		ReadDB: readDB,
		WriteDB: writeDB,
	}
}

func (r *InventoryRepo) CreateInventory(ctx context.Context, inventory *domain.Inventory) (uint, error) {
	if err := r.WriteDB.Create(inventory).Error; err != nil {
		log.Printf("error creating inventory: %v\n", err)
		return 0, err
	}

	return inventory.ID, nil
}

func (r *InventoryRepo) CheckInventory(ctx context.Context, productId uint, vendorId uint) (error) {
	var inventory domain.Inventory

	if err := r.ReadDB.Where(map[string]interface{}{"product_id": productId, "vendor_id": vendorId}).First(&inventory).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		log.Printf("error finding inventory: %v\n", err)

		return err
	} else {
		return nil
	}

}