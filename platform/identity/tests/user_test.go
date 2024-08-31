package tests

import (
	"net/http"
	"testing"
)

func deleteUser(t *testing.T, accessToken string) {
	headers := map[string]string{
		"Authorization": "Bearer " + accessToken,
	}

	resp := makeRequest(t, "DELETE", baseURL+"/users", nil, headers)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

func TestDeleteUser(t *testing.T) {
	accessToken, _ := signUpUser(t, "testuser@example.com", "password123")

	// Delete User
	deleteUser(t, accessToken)
}
