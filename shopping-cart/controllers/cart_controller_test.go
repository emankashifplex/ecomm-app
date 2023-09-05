package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

// NewRedisClient initializes a Redis client and returns it.
func NewRedisClient(addr string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: addr,
	})
}

func TestAddItemToCartWithMiniredis(t *testing.T) {
	// Miniredis instance
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatal(err)
	}
	defer mr.Close()

	// Replace RedisClient usage with a miniredis client
	redisClient := NewRedisClient(mr.Addr())

	// CartController instance with the miniredis client
	controller := NewCartController(redisClient)

	// Sample HTTP request for adding an item
	reqBody := []byte(`{"product_id": 1, "quantity": 2, "price": 10.99}`)
	req, err := http.NewRequest("POST", "/add", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("user_id", "123")

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the AddItemToCart method to handle the request
	http.HandlerFunc(controller.AddItemToCart).ServeHTTP(rr, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, rr.Code)

}
