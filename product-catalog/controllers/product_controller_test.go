package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"ecomm-app/product-catalog/models"

	"github.com/go-pg/pg/v10"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var testDB *pg.DB // Declare a global database connection

func setup() {
	// Connect to the database only once and reuse it
	options := &pg.Options{
		User:     "eman",
		Password: "123",
		Database: "test",
		Addr:     "localhost:5432",
	}
	db := pg.Connect(options)
	testDB = db
}

func teardown() {
	// Close the database connection after all tests are finished
	testDB.Close()
}

func TestMain(m *testing.M) {
	setup()
	defer teardown()
	m.Run()
}

func TestCreateProduct(t *testing.T) {
	// Initialize the router and controller
	router := mux.NewRouter()
	productController := NewProductController(models.NewProductService(testDB))
	router.HandleFunc("/products", productController.CreateProduct).Methods("POST")

	// Create a sample product to send in the request
	product := models.Product{
		Name:         "Sample Product",
		Description:  "Sample Description",
		Price:        19.99,
		Availability: true,
	}

	// Convert the product to JSON
	productJSON, err := json.Marshal(product)
	if err != nil {
		t.Fatal(err)
	}

	// Create a request
	req, err := http.NewRequest("POST", "/products", bytes.NewBuffer(productJSON))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Serve the request
	router.ServeHTTP(rr, req)

	// Check the response status code
	assert.Equal(t, http.StatusCreated, rr.Code)

	// You can add further assertions to check the response body, database state, etc.
}
