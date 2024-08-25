package tests

import (
	"bytes"
	"encoding/json"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"pulselog/identity/config"
	"pulselog/identity/models"
	"pulselog/identity/routes"
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

	err = db.AutoMigrate(
		&models.User{},
		&models.RefreshToken{},
		&models.Project{},
		&models.APIKey{},
		&models.ProjectMember{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func mockEnvs() {
	os.Setenv("IDENTITY_DB_USER", "test_user")
	os.Setenv("IDENTITY_DB_PASSWORD", "test_password")
	os.Setenv("IDENTITY_DB_NAME", "test_db")
	os.Setenv("IDENTITY_DB_PORT", "5432")
	os.Setenv("IDENTITY_DB_HOST", "localhost")
	os.Setenv("JWT_SECRET", "super_secret_jwt_key")
}

func setup(port string) (*httptest.Server, error) {
	mockEnvs()

	config.LoadEnvironmentVars()

	db, err := mockDB()
	if err != nil {
		return nil, err
	}

	router := gin.Default()
	routes.SetupAuthRoutes(router, db)
	routes.SetupUserRoutes(router, db)

	server := httptest.NewUnstartedServer(router)

	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return nil, err
	}
	server.Listener = l
	server.Start()

	return server, nil
}

func makeRequest(t *testing.T, method, url string, payload interface{}, headers map[string]string) *http.Response {
	var payloadBytes []byte
	var err error

	if payload != nil {
		payloadBytes, err = json.Marshal(payload)
		if err != nil {
			t.Fatalf("Failed to marshal payload: %v", err)
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}

	return resp
}

func TestMain(m *testing.M) {
	server, err := setup("4000")
	if err != nil {
		log.Fatalf("Setup failed: %v", err)
	}
	defer server.Close()

	code := m.Run()
	os.Exit(code)
}
