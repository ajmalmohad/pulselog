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

func (r *ProjectRepository) FindByIDUserAndRoles(projectID uint, userID uint, allowedRoles []models.Role) (*models.Project, error) {
	var project models.Project
	err := r.db.
		Where("projects.id = ?", projectID).
		Where("projects.owner_id = ? OR project_members.user_id = ?", userID, userID).
		Where("project_members.role IN ?", allowedRoles).
		Joins("LEFT JOIN project_members ON project_members.project_id = projects.id").
		First(&project).Error
	if err != nil {
		return nil, err
	}
	return &project, nil
}

func (r *ProjectRepository) FindAllByUserID(userID uint) ([]models.Project, error) {
	var projects []models.Project
	err := r.db.
		Where("projects.owner_id = ? OR project_members.user_id = ?", userID, userID).
		Joins("LEFT JOIN project_members ON project_members.project_id = projects.id").
		Find(&projects).Error
	if err != nil {
		return nil, err
	}
	return projects, nil
}
