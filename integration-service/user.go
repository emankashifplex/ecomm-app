package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// registerUser simulates user registration by sending a POST request to the registration API.
func registerUser(username, password string) int {
	// Prepare user registration data.
	data := map[string]string{"username": username, "password": password}
	jsonData, _ := json.Marshal(data)

	// Send the registration request.
	resp, err := http.Post("http://localhost:8080/register", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error registering user:", err)
		return 0
	}
	defer resp.Body.Close()

	// Check the response status code to determine if registration was successful.
	if resp.StatusCode == 201 {
		fmt.Println("User registered successfully.")
		// Extract and return the user ID from the response.
		body, _ := ioutil.ReadAll(resp.Body)
		var result map[string]int
		json.Unmarshal(body, &result)
		return result["userID"]
	}

	fmt.Println("User registration failed.")
	return 0
}

// loginUser simulates user login by sending a POST request to the login API.
func loginUser(username, password string) string {
	// Prepare login data.
	data := map[string]string{"username": username, "password": password}
	jsonData, _ := json.Marshal(data)

	// Send the login request.
	resp, err := http.Post("http://localhost:8080/login", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error logging in:", err)
		return ""
	}
	defer resp.Body.Close()

	// Check the response status code to determine if login was successful.
	if resp.StatusCode == http.StatusOK {
		fmt.Println("User logged in successfully.")
		// Extract and return the session token from the response.
		body, _ := ioutil.ReadAll(resp.Body)
		var result map[string]string
		json.Unmarshal(body, &result)
		sessionToken := result["sessionToken"]
		return sessionToken
	}

	fmt.Println("User login failed.")
	return ""
}
