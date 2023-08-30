package controllers

import (
	"ecomm-app/product-catalog/models"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

// GetProductInfo handles the request for fetching product information
func GetProductInfo(w http.ResponseWriter, r *http.Request) {
	// Extract the product ID from the URL path
	vars := mux.Vars(r)
	productID := vars["productID"]

	// Check if productID is empty
	if productID == "" {
		http.Error(w, "Missing productID parameter", http.StatusBadRequest)
		return
	}

	// Communicate with the product catalog service to fetch product information
	productInfo, err := fetchProductInfoFromCatalog(productID)
	if err != nil {
		http.Error(w, "Error fetching product information", http.StatusInternalServerError)
		return
	}

	// Return the product information as JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(productInfo)
}

func fetchProductInfoFromCatalog(productID string) (*models.Product, error) {
	resp, err := http.Get("http://localhost:8081/products/" + productID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var productInfo models.Product
		if err := json.NewDecoder(resp.Body).Decode(&productInfo); err != nil {
			return nil, err
		}
		return &productInfo, nil
	}

	return nil, errors.New("Product information not found")
}
