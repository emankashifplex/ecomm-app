package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Create an order.
func createOrder(userID int, sessionToken string, productID, quantity int) int {
	data := map[string]interface{}{
		"Product":  "Product Name",
		"Quantity": quantity,
	}
	jsonData, _ := json.Marshal(data)

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8083/orders", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating create order request:", err)
		return 0
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Id", fmt.Sprintf("%d", userID))
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", sessionToken))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error creating order:", err)
		return 0
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Order created successfully.")
		// Extract and return the order ID.
		body, _ := ioutil.ReadAll(resp.Body)
		var result map[string]int
		json.Unmarshal(body, &result)
		return result["orderID"]
	}

	fmt.Println("Order creation failed.")
	return 0
}

// Get an order by its ID using the getOrder API.
func getOrder(orderID int) {
	url := fmt.Sprintf("http://localhost:8083/orders/%d", orderID)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error retrieving order:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Printf("Order details for Order ID %d:\n", orderID)
		// Parse and display the order details.
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	} else {
		fmt.Println("Error retrieving order.")
	}
}
