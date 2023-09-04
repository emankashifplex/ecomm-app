package routes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"ecomm-app/product-catalog/models"

	"github.com/gorilla/mux"
)

func TestProductRoutes(t *testing.T) {
	// Create a new router and set up the product routes
	router := mux.NewRouter()
	SetProductRoutes(router)

	t.Run("TestCreateProduct", func(t *testing.T) {
		// product creation data
		productData := models.Product{
			Name:         "Test Product",
			Description:  "A test product",
			Price:        19.99,
			Availability: true,
		}
		productDataBytes, _ := json.Marshal(productData)

		// Create a new HTTP request for creating a product
		req, err := http.NewRequest("POST", "/products", bytes.NewReader(productDataBytes))
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

	t.Run("TestGetProductByID", func(t *testing.T) {
		// Prepare a product ID for retrieval
		productID := "1"

		// Create a new HTTP request for product retrieval by ID
		req, err := http.NewRequest("GET", "/products/"+productID, nil)
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

		var product models.Product
		if err := json.Unmarshal(recorder.Body.Bytes(), &product); err != nil {
			t.Fatal(err)
		}

		expectedName := "Test Product"
		if product.Name != expectedName {
			t.Errorf("Expected product name %s, but got %s", expectedName, product.Name)
		}
	})

}
