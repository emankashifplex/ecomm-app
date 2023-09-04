package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSendOrderConfirmationEmailController(t *testing.T) {
	// Create a new Gin router
	r := gin.New()

	// Define a sample request with user_id and order_id query parameters
	req, err := http.NewRequest(http.MethodGet, "/email?user_id=123&order_id=456", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder
	w := httptest.NewRecorder()

	// Handle the request with the SendOrderConfirmationEmailController
	r.GET("/email", SendOrderConfirmationEmailController)

	r.ServeHTTP(w, req)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

}
