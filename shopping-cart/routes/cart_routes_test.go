package routes

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"ecomm-app/shopping-cart/controllers"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetCartRoute(t *testing.T) {
	// new router instance
	r := mux.NewRouter()

	// mock CartController instance
	mockCartController := &controllers.CartController{}

	// Add the route to the router
	SetupRoutes(r, mockCartController)

	// Create a request to the /get route
	req, err := http.NewRequest("GET", "/get?user_id=123", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Serve the request using the router
	r.ServeHTTP(rr, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, rr.Code)

}

func TestRemoveItemFromCartRoute(t *testing.T) {
	// New router instance
	r := mux.NewRouter()

	// Mock CartController instance
	mockCartController := &controllers.CartController{}

	// Add the route to the router
	SetupRoutes(r, mockCartController)

	// Create a sample request body for removing an item
	reqBody := `1` // Assuming the product ID to remove is 1

	// Create a request to the /remove route
	req, err := http.NewRequest("POST", "/remove", strings.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("user_id", "123")

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Serve the request using the router
	r.ServeHTTP(rr, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, rr.Code)

}

func TestUpdateCartItemQuantityRoute(t *testing.T) {
	// New router instance
	r := mux.NewRouter()

	// Mock CartController instance
	mockCartController := &controllers.CartController{}

	// Add the route to the router
	SetupRoutes(r, mockCartController)

	// Create a sample request body for updating the quantity
	reqBody := `{"product_id": 1, "quantity": 5}`

	// Create a request to the /update route
	req, err := http.NewRequest("POST", "/update", strings.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("user_id", "123")

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Serve the request using the router
	r.ServeHTTP(rr, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestAddItemToCartRoute(t *testing.T) {
	// new router instance
	r := mux.NewRouter()

	// mock CartController instance
	mockCartController := &controllers.CartController{}

	// Add the route to the router
	SetupRoutes(r, mockCartController)

	// Sample request body
	reqBody := `{"product_id": 1, "quantity": 2, "price": 10.99}`

	// Create a request to the /add route
	req, err := http.NewRequest("POST", "/add", strings.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("user_id", "123")

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Serve the request using the router
	r.ServeHTTP(rr, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, rr.Code)

}
