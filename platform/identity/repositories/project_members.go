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

func (r *ProjectMemberRepository) FindAllByProjectID(projectID uint) ([]models.ProjectMember, error) {
	var projectMembers []models.ProjectMember
	err := r.db.
		Where("project_id = ?", projectID).
		Find(&projectMembers).
		Error
	if err != nil {
		return nil, err
	}
	return projectMembers, nil
}
