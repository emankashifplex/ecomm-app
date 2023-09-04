package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	model "ecomm-app/shopping-cart/models"

	"github.com/go-redis/redis/v8"
)

// Handles cart-related operations
type CartController struct {
	RedisClient *redis.Client
}

// Creates a new CartController instance
func NewCartController(redisClient *redis.Client) *CartController {
	return &CartController{
		RedisClient: redisClient,
	}
}

// Handles adding an item to the cart.
func (cc *CartController) AddItemToCart(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Parses request body to get product ID, quantity, and price
	var cartItem model.CartItem
	err := json.NewDecoder(r.Body).Decode(&cartItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Gets user ID from request context or headers
	userIDStr := r.Header.Get("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Construct the Redis key for the user's cart
	cartKey := GetCartKey(userID)

	// Get the current cart from Redis
	cart, err := GetCartFromRedis(ctx, cc.RedisClient, cartKey)
	if err != nil {
		http.Error(w, "Failed to retrieve cart", http.StatusInternalServerError)
		return
	}

	// Add the item to the cart
	cart.AddItem(cartItem.ProductID, cartItem.Quantity, cartItem.Price)

	// Save the updated cart to Redis
	err = SaveCartToRedis(ctx, cc.RedisClient, cartKey, cart)
	if err != nil {
		http.Error(w, "Failed to update cart", http.StatusInternalServerError)
		return
	}

	// Respond with the updated cart
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cart)
}

// RemoveItemFromCart handles removing an item from the cart.
func (cc *CartController) RemoveItemFromCart(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Parse product ID from request body
	var productID int
	err := json.NewDecoder(r.Body).Decode(&productID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get user ID from request context or headers
	userIDStr := r.Header.Get("user_id") // Assuming you pass user ID as a header
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Construct the Redis key for the user's cart
	cartKey := GetCartKey(userID)

	// Get the current cart from Redis
	cart, err := GetCartFromRedis(ctx, cc.RedisClient, cartKey)
	if err != nil {
		http.Error(w, "Failed to retrieve cart", http.StatusInternalServerError)
		return
	}

	// Remove the item from the cart
	cart.RemoveItem(productID)

	// Save the updated cart to Redis
	err = SaveCartToRedis(ctx, cc.RedisClient, cartKey, cart)
	if err != nil {
		http.Error(w, "Failed to update cart", http.StatusInternalServerError)
		return
	}

	// Respond with the updated cart
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cart)
}

// UpdateCartItemQuantity handles updating the quantity of a cart item.
func (cc *CartController) UpdateCartItemQuantity(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Parse product ID and new quantity from request body
	var updateData struct {
		ProductID int `json:"product_id"`
		Quantity  int `json:"quantity"`
	}
	err := json.NewDecoder(r.Body).Decode(&updateData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get user ID from request context or headers
	userIDStr := r.Header.Get("user_id") // Assuming you pass user ID as a header
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Construct the Redis key for the user's cart
	cartKey := GetCartKey(userID)

	// Get the current cart from Redis
	cart, err := GetCartFromRedis(ctx, cc.RedisClient, cartKey)
	if err != nil {
		http.Error(w, "Failed to retrieve cart", http.StatusInternalServerError)
		return
	}

	// Update the quantity of the cart item
	cart.UpdateQuantity(updateData.ProductID, updateData.Quantity)

	// Save the updated cart to Redis
	err = SaveCartToRedis(ctx, cc.RedisClient, cartKey, cart)
	if err != nil {
		http.Error(w, "Failed to update cart", http.StatusInternalServerError)
		return
	}

	// Respond with the updated cart
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cart)
}

// GetCart retrieves the user's cart.
func (cc *CartController) GetCart(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Get user ID from query parameter
	userIDStr := r.URL.Query().Get("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Construct the Redis key for the user's cart
	cartKey := GetCartKey(userID)

	// Get the current cart from Redis
	cart, err := GetCartFromRedis(ctx, cc.RedisClient, cartKey)
	if err != nil {
		http.Error(w, "Failed to retrieve cart", http.StatusInternalServerError)
		return
	}

	// Respond with the user's cart
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cart)
}

// getCartKey constructs the Redis key for a user's cart.
func GetCartKey(userID int) string {
	return "cart:" + strconv.Itoa(userID)
}

// getCartFromRedis retrieves a user's cart from Redis.
func GetCartFromRedis(ctx context.Context, client *redis.Client, key string) (*model.Cart, error) {
	cartJSON, err := client.Get(ctx, key).Result()
	if err == redis.Nil {
		return model.NewCart(), nil
	} else if err != nil {
		return nil, err
	}

	var cart model.Cart
	err = json.Unmarshal([]byte(cartJSON), &cart)
	return &cart, err
}

// saveCartToRedis saves a user's cart to Redis.
func SaveCartToRedis(ctx context.Context, client *redis.Client, key string, cart *model.Cart) error {
	cartJSON, err := json.Marshal(cart)
	if err != nil {
		return err
	}

	return client.Set(ctx, key, cartJSON, 0).Err()
}
