package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Calculate and display the shipping cost for an order.
func calculateShippingCost(orderID int) {
	data := map[string]interface{}{
		"ID":             orderID,
		"RecipientAddr":  "456 Elm St",
		"Weight":         5.0,
		"ShippingOption": "standard",
	}
	jsonData, _ := json.Marshal(data)
	resp, err := http.Post("http://localhost:8084/calculateshippingcost", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error calculating shipping cost:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Shipping cost calculated successfully.")

		// Parse and display the shipping cost from the response body.
		var result map[string]float64
		decoder := json.NewDecoder(resp.Body)
		if err := decoder.Decode(&result); err != nil {
			fmt.Println("Error parsing shipping cost:", err)
			return
		}

		shippingCost, exists := result["shippingCost"]
		if exists {
			fmt.Printf("Shipping Cost: $%.2f\n", shippingCost)
		} else {
			fmt.Println("Shipping cost not found in response.")
		}
	} else {
		fmt.Println("Shipping cost calculation failed.")
	}
}
