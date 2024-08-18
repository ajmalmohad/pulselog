package repositories

import (
	"pulselog/identity/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	GenericRepository[models.User]
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		GenericRepository: NewGenericRepository[models.User](db),
		db:                db,
	}
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	result := r.db.Where("email = ?", email).First(&user)
	return &user, result.Error
}
