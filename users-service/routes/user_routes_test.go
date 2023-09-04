package routes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"ecomm-app/users-service/models"

	"github.com/gorilla/mux"
)

func TestUserRoutes(t *testing.T) {
	// Create a new router and set up the user routes
	router := mux.NewRouter()
	SetUserRoutes(router)

	t.Run("TestRegisterUser", func(t *testing.T) {
		// Prepare user registration data
		userData := models.User{
			Username: "testuser",
			Password: "testpassword",
		}
		userDataBytes, _ := json.Marshal(userData)

		// Create a new HTTP request for user registration
		req, err := http.NewRequest("POST", "/register", bytes.NewReader(userDataBytes))
		if err != nil {
			t.Fatal(err)
		}

		// Create a recorder to capture the response
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		// Check the response status code
		if recorder.Code != http.StatusCreated {
			t.Errorf("Expected status %d, but got %d", http.StatusCreated, recorder.Code)
		}
	})

	t.Run("TestLoginUser", func(t *testing.T) {
		// Prepare user login data
		loginData := struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}{
			Username: "testuser",
			Password: "testpassword",
		}
		loginDataBytes, _ := json.Marshal(loginData)

		// Create a new HTTP request for user login
		req, err := http.NewRequest("POST", "/login", bytes.NewReader(loginDataBytes))
		if err != nil {
			t.Fatal(err)
		}

		// Create a recorder to capture the response
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		// Check the response status code
		if recorder.Code != http.StatusOK {
			t.Errorf("Expected status %d, but got %d", http.StatusOK, recorder.Code)
		}

		// Check the response message
		var response map[string]string
		if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
			t.Fatal(err)
		}

		if message, ok := response["message"]; !ok || message != "Login successful" {
			t.Errorf("Expected message 'Login successful', but got %s", message)
		}
	})

}
