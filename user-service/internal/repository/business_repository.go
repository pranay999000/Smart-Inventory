package repository

import (
	"context"
	"log"

	"github.com/pranay999000/smart-inventory/user-service/internal/domain"
	"gorm.io/gorm"
)

type BusinessRepo struct {
	ReadDB	*gorm.DB
	WriteDB	*gorm.DB
}

func NewBusinessRepo(readDB *gorm.DB, writeDB *gorm.DB) *BusinessRepo {
	return &BusinessRepo{
		ReadDB: readDB,
		WriteDB: writeDB,
	}
}

func (r *BusinessRepo) CreateBusiness(ctx context.Context, business *domain.Business) (uint, error) {

	if err := r.WriteDB.Create(business).Error; err != nil {
		log.Printf("Failed to create business: %v", err)
		return 0, err
	}
	return business.ID, nil

}