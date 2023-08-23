package main

import (
	"context"
	"ecomm-app/shopping-cart/routes"
	"fmt"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

var (
	redisClient *redis.Client
)

func main() {
	// Initialize Redis client
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Your Redis server address
		Password: "",               // No password by default
		DB:       0,                // Default DB
	})

	// Ping Redis to ensure the connection is established
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
		return
	}

	// Initialize router
	router := mux.NewRouter()

	// Set up routes
	routes.SetupRoutes(router)

	// Start the HTTP server
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", router)
}
