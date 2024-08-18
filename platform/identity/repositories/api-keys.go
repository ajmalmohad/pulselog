package repositories

import (
	"pulselog/identity/models"

	"gorm.io/gorm"
)

type APIKeyRepository struct {
	GenericRepository[models.APIKey]
	db *gorm.DB
}

func NewAPIKeyRepository(db *gorm.DB) *APIKeyRepository {
	return &APIKeyRepository{
		GenericRepository: NewGenericRepository[models.APIKey](db),
		db:                db,
	}
}
