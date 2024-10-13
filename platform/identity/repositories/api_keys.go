package repositories

import (
	"pulselog/identity/models"

	"gorm.io/gorm"
)

type APIKeyRepository struct {
	GenericRepository[models.APIKey]
	DB *gorm.DB
}

func NewAPIKeyRepository(db *gorm.DB) *APIKeyRepository {
	return &APIKeyRepository{
		GenericRepository: NewGenericRepository[models.APIKey](db),
		DB:                db,
	}
}

func (a *APIKeyRepository) GetAPIKeysByUserID(userID uint) ([]models.APIKey, error) {
	var apiKeys []models.APIKey
	err := a.DB.
		Where("created_by = ?", userID).
		Find(&apiKeys).
		Error
	if err != nil {
		return nil, err
	}
	return apiKeys, nil
}
