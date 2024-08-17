package tests

import (
	"log"
	"net/http/httptest"
	"os"
	"pulselog/auth/models"
	"pulselog/auth/routes"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func mockDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.User{}, &models.RefreshToken{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func setup() (*httptest.Server, error) {
	db, err := mockDB()
	if err != nil {
		return nil, err
	}

	router := gin.Default()
	routes.SetupRoutes(router, db)

	server := httptest.NewServer(router)
	return server, nil
}

func TestMain(m *testing.M) {
	server, err := setup()
	if err != nil {
		log.Fatalf("Setup failed: %v", err)
	}
	defer server.Close()

	code := m.Run()
	os.Exit(code)
}

func TestExample(t *testing.T) {
	// Example test case
	if 1+1 != 2 {
		t.Error("Math is broken")
	}
}
