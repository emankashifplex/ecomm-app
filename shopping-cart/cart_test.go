package main

import (
	"bytes"
	"context"
	"ecomm-app/shopping-cart/controllers"
	"ecomm-app/shopping-cart/models"
	"ecomm-app/shopping-cart/routes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestAddItemToCartIntegration(t *testing.T) {
	// Initialized a Redis client for testing
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Cart controller instance is created using the Redis client
	cartController := controllers.NewCartController(redisClient)

	// Clean up the test environment
	defer cleanupTestEnvironment(redisClient)

	// Create a new cart instance with some initial state
	cart := models.NewCart()
	cart.AddItem(111, 2, 10.0)

	// Simulate the incoming cart item data
	cartItem := models.CartItem{
		ProductID: 2,
		Quantity:  3,
		Price:     15.0,
	}

	// Simulate the Redis key for the user's cart
	cartKey := controllers.GetCartKey(111)

	// Save the initial cart state to Redis for testing
	controllers.SaveCartToRedis(context.Background(), redisClient, cartKey, cart)

	// Create a new HTTP request and response recorder
	requestBody, _ := json.Marshal(cartItem)
	req, _ := http.NewRequest("POST", "/add-to-cart", bytes.NewBuffer(requestBody))
	req.Header.Set("User-Id", "111")
	rr := httptest.NewRecorder()

	// Serve the request using your router and controller
	r := mux.NewRouter()
	routes.SetupRoutes(r, cartController)
	r.ServeHTTP(rr, req)

	// Verify the HTTP status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Get the updated cart from Redis
	updatedCart, _ := controllers.GetCartFromRedis(context.Background(), redisClient, cartKey)

	// Verify the changes
	assert.Equal(t, 2, len(updatedCart.Items)) // New item should be added
	assert.Equal(t, 2, updatedCart.Items[0].Quantity)
	assert.Equal(t, 3, updatedCart.Items[1].Quantity)
	assert.Equal(t, 55.0, updatedCart.Total) // Total should be updated

	// Clean up the test environment
	cleanupTestEnvironment(redisClient)
}

// Helper function to clean up the test environment
func cleanupTestEnvironment(client *redis.Client) {
	userID := 111
	cartKey := controllers.GetCartKey(userID)
	client.Del(context.Background(), cartKey)
}
