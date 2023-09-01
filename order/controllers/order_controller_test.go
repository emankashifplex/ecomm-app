package controllers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"ecomm-app/order/models"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

var testDB *sql.DB // Declare a global database connection

func setup() {
	// Open the database connection only once and reuse it
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
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

func TestCreateOrder(t *testing.T) {
	// Initialize the router and controller
	router := mux.NewRouter()
	orderController := &OrderController{DB: testDB}
	router.HandleFunc("/orders", orderController.CreateOrder).Methods("POST")

	// Create a sample order to send in the request
	order := models.Order{
		ID:        1,
		Product:   "Sample Product",
		Quantity:  5,
		Status:    "Pending",
		CreatedAt: time.Now(),
	}

	// Convert the order to JSON
	orderJSON, err := json.Marshal(order)
	if err != nil {
		t.Fatal(err)
	}

	// Create a request
	req, err := http.NewRequest("POST", "/orders", bytes.NewBuffer(orderJSON))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Serve the request
	router.ServeHTTP(rr, req)

	// Check the response status code
	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestGetOrder(t *testing.T) {
	// Initialize the router and controller
	router := mux.NewRouter()
	orderController := &OrderController{DB: testDB}
	router.HandleFunc("/orders/{orderID}", orderController.GetOrder).Methods("GET")

	// Insert a sample order into the database
	_, err := testDB.Exec("INSERT INTO orders(id, product, quantity, status, created_at) VALUES($1, $2, $3, $4,$5)",
		1, "Sample Product", 5, string(models.StatusPending), time.Now())
	if err != nil {
		t.Fatal(err)
	}

	// Create a request
	req, err := http.NewRequest("GET", "/orders/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Serve the request
	router.ServeHTTP(rr, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, rr.Code)
}
