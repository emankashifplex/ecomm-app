package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"ecomm-app/users-service/controllers"
	"ecomm-app/users-service/models"
)

func TestRegisterUserHandler(t *testing.T) {
	// Creates a new user to register
	user := models.User{
		Username: "testuser",
		Password: "testpassword",
	}

	// Converts user to JSON
	userJSON, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}

	// Creates a new HTTP request with the JSON user data
	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(userJSON))
	if err != nil {
		t.Fatal(err)
	}

	// Creates a response recorder to capture the response
	recorder := httptest.NewRecorder()

	// Calls the RegisterUserHandler function to handle the request
	controllers.RegisterUserHandler(recorder, req)

	// Checks the response status code
	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, recorder.Code)
	}

}

func TestLoginHandler(t *testing.T) {
	// Define test user credentials
	testUser := models.User{
		Username: "testuser",
		Password: "testpassword",
	}

	// Convert credentials to JSON
	creds := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{
		Username: testUser.Username,
		Password: "testpassword",
	}

	credsJSON, err := json.Marshal(creds)
	if err != nil {
		t.Fatal(err)
	}

	// Creates a new HTTP request with the JSON credentials
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(credsJSON))
	if err != nil {
		t.Fatal(err)
	}

	// Creates a response recorder to capture the response
	recorder := httptest.NewRecorder()

	// Calls the LoginHandler function to handle the request
	controllers.LoginHandler(recorder, req)

	// Check the response status code
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}
}
