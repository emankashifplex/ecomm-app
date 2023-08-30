package controllers

import (
	"ecomm-app/shipping/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func CalculateShippingCostHandler(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&order); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	shippingCost := models.CalculateShippingCost(order)
	fmt.Fprintf(w, "Shipping cost: $%.2f", shippingCost)
}
