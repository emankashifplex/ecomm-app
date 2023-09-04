package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockCartController is a mock implementation of the CartController.
type MockCartController struct {
	mock.Mock
}

func (m *MockCartController) AddItemToCart(w http.ResponseWriter, r *http.Request) {
	m.Called(w, r)
}

func (m *MockCartController) RemoveItemFromCart(w http.ResponseWriter, r *http.Request) {
	m.Called(w, r)
}

func (m *MockCartController) UpdateCartItemQuantity(w http.ResponseWriter, r *http.Request) {
	m.Called(w, r)
}

func (m *MockCartController) GetCart(w http.ResponseWriter, r *http.Request) {
	m.Called(w, r)
}

func TestSetupRoutes(t *testing.T) {
	// Create a new router
	router := mux.NewRouter()

	// Create a mock CartController
	mockCartController := new(MockCartController)

	// Set up routes using the mock CartController
	SetupRoutes(router, mockCartController)

	// Define the expected route paths
	expectedRoutes := []struct {
		path   string
		method string
	}{
		{"/add-to-cart", "POST"},
		{"/remove-from-cart", "POST"},
		{"/update-cart-item", "POST"},
		{"/get-cart", "GET"},
	}

	// Verify that the expected routes are registered
	for _, route := range expectedRoutes {
		match := false
		router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
			r := route.MatcherFunc(func(req *http.Request, rm *mux.RouteMatch) bool {
				return req.URL.Path == route.PathTemplate() && req.Method == route.GetMethods()[0]
			})

			if r.Match(httptest.NewRequest(route.method, route.path, nil), &mux.RouteMatch{}) {
				match = true
			}

			return nil
		})

		assert.True(t, match, "Expected route %s %s to be registered", route.method, route.path)
	}
}
