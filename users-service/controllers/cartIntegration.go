package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func GetCartInfo(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from the request query parameter
	userIDStr := r.URL.Query().Get("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Forward request to shopping cart service
	cartServiceURL := "http://localhost:8082/get-cart" // Replace with the actual URL
	cartServiceURL += "?user_id=" + strconv.Itoa(userID)

	resp, err := http.Get(cartServiceURL)
	if err != nil {
		http.Error(w, "Failed to fetch cart", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Decode the response from the shopping cart service
	var cartResponse map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&cartResponse)
	if err != nil {
		http.Error(w, "Failed to decode cart response", http.StatusInternalServerError)
		return
	}

	// Return response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	json.NewEncoder(w).Encode(cartResponse)
}
