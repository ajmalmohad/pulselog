package config

import (
	"fmt"
	"log"
	"pulselog/auth/models"

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
	log.Println("Connected to the database successfully!")

	err = db.AutoMigrate(&models.User{}, &models.RefreshToken{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %v", err)
	}
	log.Println("Database migrated successfully!")

	return db, nil
}
