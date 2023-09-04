package controllers

import (
	"ecomm-app/shopping-cart/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddItemToCart(t *testing.T) {
	// Create a mock Redis client
	mockRedisClient := new(MockRedisClient)

	// Create a RedisClientInterface using the mockRedisClient
	redisClient := NewRedisClient(mockRedisClient)

	// Create a CartController with the mock Redis client
	cc := NewCartController(redisClient)

	// Define a sample request payload
	payload := `{"product_id": 1, "quantity": 2, "price": 10.0}`

	// Create a sample request
	req := httptest.NewRequest("POST", "/add-to-cart", strings.NewReader(payload))
	req.Header.Set("User-Id", "123")

	// Create a sample response recorder
	w := httptest.NewRecorder()

	// Mock the RedisClient's behavior
	mockRedisClient.On("Get", mock.Anything, mock.Anything).Return(`{"user_id": 123, "items": [], "total": 0.0}`, nil)
	mockRedisClient.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	// Call the controller method
	cc.AddItemToCart(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse the response body to check if it's a valid JSON
	var responseCart models.Cart
	err := json.NewDecoder(w.Body).Decode(&responseCart)
	assert.NoError(t, err)

}
