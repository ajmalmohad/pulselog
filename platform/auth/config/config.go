package config

import (
    "log"
    "os"
    "github.com/joho/godotenv"
)

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

var AppConfig Config

func LoadConfig() {
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found")
    }

    AppConfig = Config{
        JWTSecret: getEnv("JWT_SECRET"),
        DB: DatabaseConfig{
            DBHost:     getEnv("DB_HOST"),
            DBPort:     getEnv("DB_PORT"),
            DBUser:     getEnv("DB_USER"),
            DBPassword: getEnv("DB_PASSWORD"),
            DBName:     getEnv("DB_NAME"),
        },
    }
}

func getEnv(key string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
	log.Fatalf("environment variable not set: %s", key)
    return ""
}
