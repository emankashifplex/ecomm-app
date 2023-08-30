package controllers

import (
	"database/sql"
	"ecomm-app/order/models"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// OrderController manages HTTP requests related to orders
type OrderController struct {
	DB *sql.DB // A reference to the database connection
}

// CreateOrder handles the creation of a new order
func (oc *OrderController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	// Decode incoming JSON data into an Order struct
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set order status to "Pending" and record the creation time
	order.Status = string(models.StatusPending)
	order.CreatedAt = time.Now()

	// Insert the order into the database
	_, err := oc.DB.Exec("INSERT INTO orders(product, quantity, status, created_at) VALUES($1, $2, $3, $4)",
		order.Product, order.Quantity, order.Status, order.CreatedAt)
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
	err := oc.DB.QueryRow("SELECT id, product, quantity, status, created_at FROM orders WHERE id = $1", orderID).
		Scan(&order.ID, &order.Product, &order.Quantity, &order.Status, &order.CreatedAt)
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
