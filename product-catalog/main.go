package main

import (
	"ecomm-app/product-catalog/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/go-pg/pg/v10"
	"github.com/gorilla/mux"
)

func main() {
	// PostgreSQL database connection
	db := pg.Connect(&pg.Options{
		User:     "eman",
		Password: "123",
		Database: "product_data",
		Addr:     "localhost:5432",
	})

	defer db.Close()

	// Create a new router instance
	router := mux.NewRouter()

	// Set up routes for the product service
	routes.SetProductRoutes(router, db)

	// Start the HTTP server
	serverAddr := "localhost:8081"
	fmt.Printf("Starting server on %s...\n", serverAddr)
	err := http.ListenAndServe(serverAddr, router)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
