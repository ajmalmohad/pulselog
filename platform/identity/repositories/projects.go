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

func (r *ProjectRepository) FindByIDAndUser(project_id uint, user_id uint) (*models.Project, error) {
	var project models.Project
	err := r.db.
		Where("projects.id = ?", project_id).
		Where("projects.owner_id = ? OR project_members.user_id = ?", user_id, user_id).
		Joins("LEFT JOIN project_members ON project_members.project_id = projects.id").
		First(&project).Error
	if err != nil {
		return nil, err
	}
	return &project, nil
}
