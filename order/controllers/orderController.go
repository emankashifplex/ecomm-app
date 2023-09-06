package controllers

import (
	"database/sql"
	"ecomm-app/order/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// OrderController manages HTTP requests related to orders
type OrderController struct {
	DB *sql.DB // A reference to the database connection
}

var (
	userServiceEndpoint    = "http://localhost:8080"
	productServiceEndpoint = "http://localhost:8081"
)

// CreateOrder handles the creation of a new order
func (oc *OrderController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	// Decode incoming JSON data into an Order struct
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the user exists by making an HTTP request to the user service
	userExists, err := checkUserExists(order.UserID)
	if err != nil {
		http.Error(w, "Error checking user existence", http.StatusInternalServerError)
		return
	}
	if !userExists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Check if the product exists and the requested quantity is available by making an HTTP request to the product service
	productAvailable, err := checkProductAvailability(order.ProductID, order.Quantity)
	if err != nil {
		http.Error(w, "Error checking product availability", http.StatusInternalServerError)
		return
	}
	if !productAvailable {
		http.Error(w, "Product not available", http.StatusNotFound)
		return
	}

	// Set order status to "Pending" and record the creation time
	order.Status = string(models.StatusPending)
	order.CreatedAt = time.Now()

	// Insert the order into the database
	_, err = oc.DB.Exec("INSERT INTO orders(order_id, product_id, quantity, status, created_at) VALUES($1, $2, $3, $4, $5,$6)",
		order.ID, order.ProductID, order.UserID, order.Quantity, order.Status, order.CreatedAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetOrder retrieves details of a specific order.
func (oc *OrderController) GetOrder(w http.ResponseWriter, r *http.Request) {
	orderID := mux.Vars(r)["orderID"] // Extracts orderID from the URL

	var order models.Order
	// Queries the database for order details based on orderID
	err := oc.DB.QueryRow("SELECT id, product, user, quantity, status, created_at FROM orders WHERE id = $1", orderID).
		Scan(&order.ID, &order.ProductID, &order.UserID, &order.Quantity, &order.Status, &order.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Order not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert order data to JSON format
	orderJSON, err := json.Marshal(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(orderJSON)
}

// Function to check if the user exists by making an HTTP request to the user service
func checkUserExists(userID int) (bool, error) {
	// Create a GET request to the user service endpoint with the userID parameter
	resp, err := http.Get(fmt.Sprintf("%s/exists/%d", userServiceEndpoint, userID))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode == http.StatusOK {
		return true, nil
	} else if resp.StatusCode == http.StatusNotFound {
		return false, nil
	} else {
		return false, err
	}
}

// Function to check if the product exists and is available by making an HTTP request to the product service
func checkProductAvailability(productID int, quantity int) (bool, error) {
	// Create a GET request to the product service endpoint with productID and quantity parameters
	resp, err := http.Get(fmt.Sprintf("%s/availability?productID=%d&quantity=%d", productServiceEndpoint, productID, quantity))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode == http.StatusOK {
		return true, nil
	} else if resp.StatusCode == http.StatusNotFound {
		return false, nil
	} else {
		return false, err
	}
}
