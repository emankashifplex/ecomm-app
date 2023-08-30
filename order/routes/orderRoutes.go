package routes

import (
	"database/sql"
	"ecomm-app/order/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

// SetOrderRoutes configures and returns an HTTP router with order-related routes
func SetOrderRoutes(db *sql.DB) http.Handler {
	// Create a new router instance using Gorilla Mux
	r := mux.NewRouter()

	// Create an instance of the OrderController with the provided DB connection
	orderController := controllers.OrderController{DB: db}

	r.HandleFunc("/orders", orderController.CreateOrder).Methods("POST")
	r.HandleFunc("/orders/{orderID}", orderController.GetOrder).Methods("GET")

	return r
}
