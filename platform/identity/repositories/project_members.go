package repositories

import (
	"pulselog/identity/models"

	"gorm.io/gorm"
)

type ProjectMemberRepository struct {
	GenericRepository[models.ProjectMember]
	DB *gorm.DB
}

func NewProjectMemberRepository(db *gorm.DB) *ProjectMemberRepository {
	return &ProjectMemberRepository{
		GenericRepository: NewGenericRepository[models.ProjectMember](db),
		DB:                db,
	}
}

func (r *ProjectMemberRepository) FindAllByProjectID(projectID uint) ([]models.ProjectMember, error) {
	var projectMembers []models.ProjectMember
	err := r.DB.
		Where("project_id = ?", projectID).
		Find(&projectMembers).
		Error
	if err != nil {
		return nil, err
	}
	return projectMembers, nil
}
