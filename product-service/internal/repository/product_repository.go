package repository

import (
	"context"
	"log"

	"github.com/pranay999000/smart-inventory/product-service/internal/domain"
	"gorm.io/gorm"
)

type ProductRepo struct {
	ReadDB		*gorm.DB
	WriteDB		*gorm.DB
}

func NewProductRepo(readDB *gorm.DB, writeDB *gorm.DB) *ProductRepo {
	return &ProductRepo{
		ReadDB: readDB,
		WriteDB: writeDB,
	}
}

func (r *ProductRepo) CreateProduct(ctx context.Context, product *domain.Product) (uint, error) {

	if err := r.WriteDB.Create(product).Error; err != nil {
		log.Printf("Error creating product: %v\n", err)
		return 0, err
	}

	return product.ID, nil

}