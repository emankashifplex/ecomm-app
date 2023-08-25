package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	model "ecomm-app/shopping-cart/models"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

type CartController struct {
	RedisClient *redis.Client
}

func NewCartController(redisClient *redis.Client) *CartController {
	return &CartController{
		RedisClient: redisClient,
	}
}

func (cc *CartController) AddItemToCart(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Parse request body to get product ID, quantity, and price
	var cartItem model.CartItem
	err := json.NewDecoder(r.Body).Decode(&cartItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get user ID from request context or headers
	userIDStr := r.Header.Get("User-Id") // Assuming you pass user ID as a header
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	cartKey := getCartKey(userID)
	cart, err := getCartFromRedis(ctx, cc.RedisClient, cartKey)
	if err != nil {
		http.Error(w, "Failed to retrieve cart", http.StatusInternalServerError)
		return
	}

	cart.AddItem(cartItem.ProductID, cartItem.Quantity, cartItem.Price)

	err = saveCartToRedis(ctx, cc.RedisClient, cartKey, cart)
	if err != nil {
		http.Error(w, "Failed to update cart", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cart)
}

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
	userIDStr := r.Header.Get("User-Id") // Assuming you pass user ID as a header
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	cartKey := getCartKey(userID)
	cart, err := getCartFromRedis(ctx, cc.RedisClient, cartKey)
	if err != nil {
		http.Error(w, "Failed to retrieve cart", http.StatusInternalServerError)
		return
	}

	cart.RemoveItem(productID)

	err = saveCartToRedis(ctx, cc.RedisClient, cartKey, cart)
	if err != nil {
		http.Error(w, "Failed to update cart", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cart)
}

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
	userIDStr := r.Header.Get("User-Id") // Assuming you pass user ID as a header
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	cartKey := getCartKey(userID)
	cart, err := getCartFromRedis(ctx, cc.RedisClient, cartKey)
	if err != nil {
		http.Error(w, "Failed to retrieve cart", http.StatusInternalServerError)
		return
	}

	cart.UpdateQuantity(updateData.ProductID, updateData.Quantity)

	err = saveCartToRedis(ctx, cc.RedisClient, cartKey, cart)
	if err != nil {
		http.Error(w, "Failed to update cart", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cart)
}

func (cc *CartController) GetCart(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	userIDStr := r.URL.Query().Get("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	cartKey := getCartKey(userID)
	cart, err := getCartFromRedis(ctx, cc.RedisClient, cartKey)
	if err != nil {
		http.Error(w, "Failed to retrieve cart", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cart)
}

func getCartKey(userID int) string {
	return "cart:" + strconv.Itoa(userID)
}

func getCartFromRedis(ctx context.Context, client *redis.Client, key string) (*model.Cart, error) {
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

func saveCartToRedis(ctx context.Context, client *redis.Client, key string, cart *model.Cart) error {
	cartJSON, err := json.Marshal(cart)
	if err != nil {
		return err
	}

	return client.Set(ctx, key, cartJSON, 0).Err()
}
