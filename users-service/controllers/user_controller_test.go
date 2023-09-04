package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterUserHandler_Success(t *testing.T) {
	// Prepare the request body
	reqBody := map[string]interface{}{
		"username": "testuser",
		"password": "testpassword",
	}
	body, _ := json.Marshal(reqBody)

	// Create a request
	req, err := http.NewRequest("POST", "/register", bytes.NewReader(body))
	assert.NoError(t, err)

	// Recorder to capture response
	rr := httptest.NewRecorder()

	// Call the handler
	RegisterUserHandler(rr, req)

	// Assert the response
	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestLoginHandler_Success(t *testing.T) {
	// Prepare the request body
	reqBody := map[string]interface{}{
		"username": "testuser",
		"password": "testpassword",
	}
	body, _ := json.Marshal(reqBody)

	// Create the request
	req, err := http.NewRequest("POST", "/login", bytes.NewReader(body))
	assert.NoError(t, err)

	// Recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the handler
	LoginHandler(rr, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, rr.Code)

}
