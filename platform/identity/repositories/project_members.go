package repositories

import (
	"pulselog/identity/models"

	"gorm.io/gorm"
)

type ProjectMemberRepository struct {
	GenericRepository[models.ProjectMember]
	db *gorm.DB
}

func NewProjectMemberRepository(db *gorm.DB) *ProjectMemberRepository {
	return &ProjectMemberRepository{
		GenericRepository: NewGenericRepository[models.ProjectMember](db),
		db:                db,
	}
}
