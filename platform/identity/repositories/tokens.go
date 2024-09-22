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
	if result.Error != nil {
		return nil, result.Error
	}
	return &refreshToken, nil
}

func (r *RefreshTokenRepository) DeleteByTokenAndUserID(token string, userID uint) error {
	result := r.db.Where("token = ? AND user_id = ?", token, userID).Delete(&models.RefreshToken{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *RefreshTokenRepository) DeleteByUserID(userID uint) error {
	result := r.db.Where("user_id = ?", userID).Delete(&models.RefreshToken{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
