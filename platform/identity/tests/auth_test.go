package tests

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"testing"
)

func signUpUser(t *testing.T, email, password string) (string, string) {
	signUpPayload := map[string]string{
		"name":     "testuser",
		"password": password,
		"email":    email,
	}

	resp := makeRequest(t, "POST", baseURL+"/auth/signup", signUpPayload, nil)
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

	resp := makeRequest(t, "POST", baseURL+"/auth/login", loginPayload, nil)
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

	resp := makeRequest(t, "POST", baseURL+"/auth/reauthenticate", reauthPayload, nil)
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

func TestMain(m *testing.M) {
	server, err := setup(port)
	if err != nil {
		log.Fatalf("Setup failed: %v", err)
	}
	defer server.Close()

	code := m.Run()
	os.Exit(code)
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

	resp := makeRequest(t, "POST", baseURL+"/auth/signup", signUpPayload, nil)
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

	resp := makeRequest(t, "POST", baseURL+"/auth/login", loginPayload, nil)
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

	resp := makeRequest(t, "POST", baseURL+"/auth/reauthenticate", reauthPayload, nil)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, resp.StatusCode)
	}

	// Clean-Up
	deleteUser(t, accessToken)
}
