package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type CartItem struct {
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

// Add an item to the cart using the updated API format.
func addToCart(userID, productID, quantity int, price float64) {
	url := "http://localhost:8082/add"

	// Create the request body.
	requestBody := map[string]interface{}{
		"product_id": productID,
		"quantity":   quantity,
		"price":      price,
	}

	// Marshal the request body to JSON.
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error creating JSON request:", err)
		return
	}

	// Create the HTTP request.
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the User-Id header.
	req.Header.Set("User-Id", strconv.Itoa(userID))
	req.Header.Set("Content-Type", "application/json")

	// Send the HTTP request.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error adding item to cart:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Printf("Item with Product ID %d added to the cart for User ID %d.\n", productID, userID)
	} else {
		fmt.Println("Error adding item to cart.")
	}
}

// Remove an item from the cart using the updated API format.
func removeFromCart(userID, productID int) {
	url := "http://localhost:8082/remove"

	// Create the request body.
	requestBody := map[string]interface{}{
		"product_id": productID,
	}

	// Marshal the request body to JSON.
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error creating JSON request:", err)
		return
	}

	// Create the HTTP request.
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the User-Id header.
	req.Header.Set("User-Id", strconv.Itoa(userID))
	req.Header.Set("Content-Type", "application/json")

	// Send the HTTP request.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error removing item from cart:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Printf("Item with Product ID %d removed from the cart for User ID %d.\n", productID, userID)
	} else {
		fmt.Println("Error removing item from cart.")
	}
}

func viewCart(userID int) {
	url := fmt.Sprintf("http://localhost:8082/get?user_id=%d", userID)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error viewing cart:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Printf("Cart for User ID %d:\n", userID)
		// Parse and display the cart contents.
		body, _ := ioutil.ReadAll(resp.Body)
		var cartItems []CartItem
		if err := json.Unmarshal(body, &cartItems); err != nil {
			fmt.Println("Error parsing cart contents:", err)
			return
		}

		for _, item := range cartItems {
			fmt.Printf("Product ID: %d, Quantity: %d, Price: %.2f\n", item.ProductID, item.Quantity, item.Price)
		}
	} else {
		fmt.Println("Error viewing cart.")
	}
}
