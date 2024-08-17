package repositories

import (
	"pulselog/auth/models"

	"gorm.io/gorm"
)

type RefreshTokenRepository struct {
	db *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{
		db: db,
	}
}

func (r *RefreshTokenRepository) Create(refreshToken *models.RefreshToken) (*models.RefreshToken, error) {
	result := r.db.Create(refreshToken)
	return refreshToken, result.Error
}

func (r *RefreshTokenRepository) Update(refreshToken *models.RefreshToken) (*models.RefreshToken, error) {
	result := r.db.Save(refreshToken)
	return refreshToken, result.Error
}

func (r *RefreshTokenRepository) Delete(refreshToken *models.RefreshToken) (*models.RefreshToken, error) {
	result := r.db.Delete(refreshToken)
	return refreshToken, result.Error
}

func (r *RefreshTokenRepository) FindByID(id uint) (*models.RefreshToken, error) {
	var refreshToken models.RefreshToken
	result := r.db.First(&refreshToken, id)
	return &refreshToken, result.Error
}

func (r *RefreshTokenRepository) FindByToken(token string) (*models.RefreshToken, error) {
	var refreshToken models.RefreshToken
	result := r.db.Where("token = ?", token).First(&refreshToken)
	return &refreshToken, result.Error
}

func (r *RefreshTokenRepository) FindAll() ([]*models.RefreshToken, error) {
	var refreshTokens []*models.RefreshToken
	result := r.db.Find(&refreshTokens)
	return refreshTokens, result.Error
}
