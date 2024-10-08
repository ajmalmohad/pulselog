package config

import (
	"fmt"
	"pulselog/identity/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func buildDatabaseURI() string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		AppConfig.DB.DBUser,
		AppConfig.DB.DBPassword,
		AppConfig.DB.DBHost,
		AppConfig.DB.DBPort,
		AppConfig.DB.DBName,
	)
}

func InitDatabase() (*gorm.DB, error) {
	databaseURI := buildDatabaseURI()

	db, err := gorm.Open(postgres.Open(databaseURI), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.RefreshToken{},
		&models.Project{},
		&models.APIKey{},
		&models.ProjectMember{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %v", err)
	}

	return db, nil
}
