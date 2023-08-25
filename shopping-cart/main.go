package main

import (
	"log"
	"net/http"

	"ecomm-app/shopping-cart/controllers"
	"ecomm-app/shopping-cart/routes"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Initialize Redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Change this to your Redis server address
	})

	// Create an instance of your CartController with the Redis client
	cartController := controllers.NewCartController(redisClient)

	routes.SetupRoutes(r, cartController)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
