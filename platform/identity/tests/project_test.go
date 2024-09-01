package tests

import (
	"encoding/json"
	"net/http"
	"pulselog/identity/models"
	"testing"
)

func createProject(t *testing.T, accessToken string) uint {
	headers := map[string]string{
		"Authorization": "Bearer " + accessToken,
	}

	createProjectPayload := map[string]string{
		"name": "Test Project",
	}

	resp := makeRequest(t, "POST", baseURL+"/projects", createProjectPayload, headers)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
	}

	var createProjectResponse struct {
		Message string `json:"message"`
		Data    struct {
			ID uint `json:"id"`
		} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&createProjectResponse); err != nil {
		t.Fatalf("Failed to decode create project response: %v", err)
	}

	return createProjectResponse.Data.ID
}

func getAllProjects(t *testing.T, accessToken string) []models.Project {
	headers := map[string]string{
		"Authorization": "Bearer " + accessToken,
	}

	resp := makeRequest(t, "GET", baseURL+"/projects/all", nil, headers)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var getAllProjectsResponse struct {
		Message string `json:"message"`
		Data    []models.Project
	}

	if err := json.NewDecoder(resp.Body).Decode(&getAllProjectsResponse); err != nil {
		t.Fatalf("Failed to decode get all projects response: %v", err)
	}

	return getAllProjectsResponse.Data
}

func TestCreateProject(t *testing.T) {
	// Initial Sign Up
	accessToken, _ := signUpUser(t, "testuser@example.com", "password123")

	// Create Project
	headers := map[string]string{
		"Authorization": "Bearer " + accessToken,
	}

	createProjectPayload := map[string]string{
		"name": "Test Project",
	}

	resp := makeRequest(t, "POST", baseURL+"/projects", createProjectPayload, headers)

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
	}

	// Cleanup
	deleteUser(t, accessToken)
}

func TestGetAllProjects(t *testing.T) {
	// Initial Sign Up
	accessToken, _ := signUpUser(t, "testuser@example.com", "password123")

	// Create Project
	_ = createProject(t, accessToken)

	// Get All Projects
	projects := getAllProjects(t, accessToken)

	// Check response
	if len(projects) != 1 {
		t.Errorf("Expected at least one project, got %d", len(projects))
	}

	// Cleanup
	deleteUser(t, accessToken)
}
