package main

import (
	"context"
	"log"
	"net/http"

	"ecomm-app/users-service/routes"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
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
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5500/frontend/index.html"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	// Use the CORS middleware to wrap router
	handler := c.Handler(router)

	// Start the HTTP server on port 8080
	err = http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
