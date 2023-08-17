package main

import (
	"context"
	"log"
	"net/http"

	"ecomm-app/users-service/routes"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Connect to the MongoDB database
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer client.Disconnect(context.TODO())

	// Create a new router using Gorilla Mux
	router := mux.NewRouter()

	// Set up user-related routes using the SetUserRoutes function from routes package
	routes.SetUserRoutes(router)

	// Handle all incoming requests using the router
	http.Handle("/", router)

	// Start the HTTP server on port 8080
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
