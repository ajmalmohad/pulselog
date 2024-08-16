package config

import (
    "fmt"
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "pulselog/auth/models"
)

var DB *gorm.DB

func ConnectDatabase() {
    databaseURI := buildDatabaseURI()
    
    var err error
    DB, err = gorm.Open(postgres.Open(databaseURI), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to the database: %v", err)
    }
    log.Println("Connected to the database successfully!")

    err = DB.AutoMigrate(&models.User{}, &models.RefreshToken{})
    if err != nil {
        log.Fatalf("failed to migrate database: %v", err)
    }
    log.Println("Database migrated successfully!")
}

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