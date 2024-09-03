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

func (a *APIKeyRepository) GetAPIKeysByUserID(userID uint) ([]models.APIKey, error) {
	var apiKeys []models.APIKey
	err := a.db.
		Where("created_by = ?", userID).
		Find(&apiKeys).
		Error
	return apiKeys, err
}
