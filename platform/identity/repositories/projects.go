package repositories

import (
	"pulselog/identity/models"

	"gorm.io/gorm"
)

type ProjectRepository struct {
	GenericRepository[models.Project]
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{
		GenericRepository: NewGenericRepository[models.Project](db),
		db:                db,
	}
}
