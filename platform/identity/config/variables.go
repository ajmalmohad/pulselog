package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

var AppConfig *Config

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	log.Fatalf("environment variable not set: %s", key)
	return ""
}

type Config struct {
	DB        DatabaseConfig
	JWTSecret string
}

type DatabaseConfig struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func LoadEnvironmentVars(relativePath string) error {
	var err error

	if relativePath == "" {
		err = godotenv.Load()
	} else {
		var absPath string
		absPath, err = filepath.Abs(relativePath)
		if err != nil {
			return err
		}

		err = godotenv.Load(absPath)
	}

	if err != nil {
		return err
	}

	config := Config{
		JWTSecret: getEnv("JWT_SECRET"),
		DB: DatabaseConfig{
			DBHost:     getEnv("IDENTITY_DB_HOST"),
			DBPort:     getEnv("IDENTITY_DB_PORT"),
			DBUser:     getEnv("IDENTITY_DB_USER"),
			DBPassword: getEnv("IDENTITY_DB_PASSWORD"),
			DBName:     getEnv("IDENTITY_DB_NAME"),
		},
	}

	AppConfig = &config
	return nil
}
