package tests

// TODO: Refactor this to a better test suite
// Right now this is fine to test the API endpoints

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

func TestMain(m *testing.M) {
	server, err := setup("4000")
	if err != nil {
		log.Fatalf("Setup failed: %v", err)
	}
	defer server.Close()

	code := m.Run()
	os.Exit(code)
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

func signUpUser(t *testing.T, email, password string) (string, string) {
	signUpPayload := map[string]string{
		"name":     "testuser",
		"password": password,
		"email":    email,
	}

	resp := makeRequest(t, "POST", "http://localhost:4000/auth/signup", signUpPayload, nil)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
	}

	var signUpResponse struct {
		Message string `json:"message"`
		Data    struct {
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token"`
		} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&signUpResponse); err != nil {
		t.Fatalf("Failed to decode sign-up response: %v", err)
	}

	return signUpResponse.Data.AccessToken, signUpResponse.Data.RefreshToken
}

func logInUser(t *testing.T, email, password string) (string, string) {
	loginPayload := map[string]string{
		"password": password,
		"email":    email,
	}

	resp := makeRequest(t, "POST", "http://localhost:4000/auth/login", loginPayload, nil)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var loginResponse struct {
		Message string `json:"message"`
		Data    struct {
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token"`
		} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&loginResponse); err != nil {
		t.Fatalf("Failed to decode login response: %v", err)
	}

	return loginResponse.Data.AccessToken, loginResponse.Data.RefreshToken
}

func reauthenticateUser(t *testing.T, refreshToken string) (string, string) {
	reauthPayload := map[string]string{
		"refresh_token": refreshToken,
	}

	resp := makeRequest(t, "POST", "http://localhost:4000/auth/reauthenticate", reauthPayload, nil)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var reauthResponse struct {
		Message string `json:"message"`
		Data    struct {
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token"`
		} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&reauthResponse); err != nil {
		t.Fatalf("Failed to decode reauth response: %v", err)
	}

	return reauthResponse.Data.AccessToken, reauthResponse.Data.RefreshToken
}

func deleteUser(t *testing.T, accessToken string) {
	headers := map[string]string{
		"Authorization": "Bearer " + accessToken,
	}

	resp := makeRequest(t, "DELETE", "http://localhost:4000/users", nil, headers)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

func TestSignUp(t *testing.T) {
	// Initial Sign-Up
	accessToken, _ := signUpUser(t, "testuser@example.com", "password123")

	// Duplicate Sign-Up
	signUpPayload := map[string]string{
		"name":     "testuser",
		"password": "password123",
		"email":    "testuser@example.com",
	}

	resp := makeRequest(t, "POST", "http://localhost:4000/auth/signup", signUpPayload, nil)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, resp.StatusCode)
	}

	// Clean-Up
	deleteUser(t, accessToken)
}

func TestLogIn(t *testing.T) {
	// Initial Sign-Up
	accessToken, _ := signUpUser(t, "testuser@example.com", "password123")

	// Initial Log-In
	_, _ = logInUser(t, "testuser@example.com", "password123")

	// Incorrect Log-In
	loginPayload := map[string]string{
		"password": "wrongpassword",
		"email":    "testuser@example.com",
	}

	resp := makeRequest(t, "POST", "http://localhost:4000/auth/login", loginPayload, nil)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, resp.StatusCode)
	}

	// Clean-Up
	deleteUser(t, accessToken)
}

func TestReauthenticate(t *testing.T) {
	// Initial Sign-Up
	accessToken, refreshToken := signUpUser(t, "testuser@example.com", "password123")

	// Reauthenticate
	reauthenticateUser(t, refreshToken)

	// Incorrect Reauthenticate
	reauthPayload := map[string]string{
		"refresh_token": "wrongtoken",
	}

	resp := makeRequest(t, "POST", "http://localhost:4000/auth/reauthenticate", reauthPayload, nil)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, resp.StatusCode)
	}

	// Clean-Up
	deleteUser(t, accessToken)
}

func TestDeleteUser(t *testing.T) {
	accessToken, _ := signUpUser(t, "testuser@example.com", "password123")

	// Delete User
	deleteUser(t, accessToken)
}
