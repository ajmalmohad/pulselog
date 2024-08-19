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

func (r *RefreshTokenRepository) DeleteByTokenAndUserID(token string, userID uint) error {
	result := r.db.Where("token = ? AND user_id = ?", token, userID).Delete(&models.RefreshToken{})
	return result.Error
}

func (r *RefreshTokenRepository) DeleteByUserID(userID uint) error {
	result := r.db.Where("user_id = ?", userID).Delete(&models.RefreshToken{})
	return result.Error
}
