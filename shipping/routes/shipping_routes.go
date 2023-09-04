package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetupRoutes(t *testing.T) {
	// Create a new test server
	server := httptest.NewServer(nil)
	defer server.Close()

	// Call the SetupRoutes function
	SetupRoutes()

	// Make a request to the /calculate_shipping_cost route
	resp, err := http.Get(server.URL + "/calculateshippingcost")
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	// Assert the response status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
	}

}
