package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func createProduct() int {
	rand.Seed(time.Now().UnixNano())

	// Generate random product data.
	productData := map[string]interface{}{
		"name":         "Bottle",
		"description":  "Blue",
		"price":        8.00,
		"availability": true,
	}

	jsonData, _ := json.Marshal(productData)
	resp, err := http.Post("http://localhost:8081/products", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating product:", err)
		return 0
	}

	defer resp.Body.Close()

	if resp.StatusCode == 201 {
		fmt.Println("Random product created successfully.")
		// Extract and return the product ID.
		body, _ := ioutil.ReadAll(resp.Body)
		var result map[string]int
		json.Unmarshal(body, &result)
		return result["ID"]
	}

	fmt.Println("Random product creation failed.")
	return 0
}

// Retrieve a product by its product ID.
func getProduct(productID int) {
	url := fmt.Sprintf("http://localhost:8081/products/%d", productID)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error retrieving product:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 201 {
		fmt.Printf("Product details for Product ID %d:\n", productID)
		// Process and display the product details.
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	} else if resp.StatusCode == http.StatusNotFound {
		fmt.Printf("Product with Product ID %d not found.\n", productID)
	} else {
		fmt.Println("Error retrieving product.")
	}
}

// Search for products using the product search API with separate parameters.
func searchProducts(query string, minPrice, maxPrice float64, availability bool) {
	// Construct the URL with query parameters.
	queryParams := "?query=" + query +
		"&minPrice=" + strconv.FormatFloat(minPrice, 'f', 2, 64) +
		"&maxPrice=" + strconv.FormatFloat(maxPrice, 'f', 2, 64) +
		"&availability=" + strconv.FormatBool(availability)
	url := "http://localhost:8081/products/search" + queryParams

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error performing product search:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Printf("Product search results for query '%s':\n", query)
		// Process and display the search results.
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	} else {
		fmt.Println("Product search failed.")
	}
}
