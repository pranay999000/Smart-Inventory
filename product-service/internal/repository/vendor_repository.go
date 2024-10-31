package repository

import (
	"context"
	"log"

	"github.com/pranay999000/smart-inventory/product-service/internal/domain"
	"gorm.io/gorm"
)

type VendorRepo struct {
	ReadDB		*gorm.DB
	WriteDB		*gorm.DB
}

func NewVendorRepo(readDB *gorm.DB, writeDB *gorm.DB) *VendorRepo {
	return &VendorRepo{
		ReadDB: readDB,
		WriteDB: writeDB,
	}
}

func (r *VendorRepo) CreateVendor(ctx context.Context, vendor *domain.Vendor) (uint, error) {

	if err := r.WriteDB.Create(vendor).Error; err != nil {
		log.Printf("Error creating product: %v\n", err)
		return 0, err
	}

	return vendor.ID, nil

}