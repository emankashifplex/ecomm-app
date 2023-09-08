package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestProductRoutes(t *testing.T) {
	// Create a new router
	router := mux.NewRouter()

	// Set up the product routes
	SetProductRoutes(router, nil)

	// Define test cases for each route
	testCases := []struct {
		Method     string
		Path       string
		ExpectCode int
	}{
		{"GET", "/products", http.StatusMethodNotAllowed}, // Expect a 405 Method Not Allowed for POST
		{"GET", "/products/123", http.StatusNotFound},     // Expect a 404 Not Found for GET
		{"GET", "/products/search", http.StatusOK},        // Expect a 200 OK for GE
	}

	for _, tc := range testCases {
		t.Run(tc.Method+" "+tc.Path, func(t *testing.T) {
			// Create a request with the specified method and path
			req, err := http.NewRequest(tc.Method, tc.Path, nil)
			if err != nil {
				t.Fatal(err)
			}

			// Create a response recorder to capture the response
			rr := httptest.NewRecorder()

			// Serve the request using the router
			router.ServeHTTP(rr, req)

			// Check the response status code
			assert.Equal(t, tc.ExpectCode, rr.Code)
		})
	}
}
