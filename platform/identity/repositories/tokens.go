package repositories

import (
	"pulselog/identity/models"

	"gorm.io/gorm"
)

type RefreshTokenRepository struct {
	GenericRepository[models.RefreshToken]
	db *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{
		GenericRepository: NewGenericRepository[models.RefreshToken](db),
		db:                db,
	}
}

func (r *RefreshTokenRepository) FindByToken(token string) (*models.RefreshToken, error) {
	var refreshToken models.RefreshToken
	result := r.db.Where("token = ?", token).First(&refreshToken)
	return &refreshToken, result.Error
}
