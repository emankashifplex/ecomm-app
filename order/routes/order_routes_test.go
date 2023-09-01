package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetOrderRoutes(t *testing.T) {
	// Create a test router using SetOrderRoutes with a nil database connection
	router := SetOrderRoutes(nil)

	// Create a test request for creating an order
	createOrderRequest, _ := http.NewRequest("POST", "/orders", nil)
	createOrderResponse := httptest.NewRecorder()

	// Create a test request for getting an order
	getOrderRequest, _ := http.NewRequest("GET", "/orders/1", nil)
	getOrderResponse := httptest.NewRecorder()

	// Serve the create order request
	router.ServeHTTP(createOrderResponse, createOrderRequest)

	// Serve the get order request
	router.ServeHTTP(getOrderResponse, getOrderRequest)

	// Check if the status codes are as expected
	if createOrderResponse.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, but got %d for create order request", http.StatusNotFound, createOrderResponse.Code)
	}

	if getOrderResponse.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, but got %d for get order request", http.StatusNotFound, getOrderResponse.Code)
	}

}
