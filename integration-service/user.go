package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func registerUser(username, password string) int {
	data := map[string]string{"username": username, "password": password}
	jsonData, _ := json.Marshal(data)
	resp, err := http.Post("http://localhost:8080/register", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error registering user:", err)
		return 0
	}
	defer resp.Body.Close()

	if resp.StatusCode == 201 {
		fmt.Println("User registered successfully.")
		// Extract and return the user ID.
		body, _ := ioutil.ReadAll(resp.Body)
		var result map[string]int
		json.Unmarshal(body, &result)
		return result["userID"]
	}

	fmt.Println("User registration failed.")
	return 0
}

// Simulate user login and return the session token.
func loginUser(username, password string) string {
	data := map[string]string{"username": username, "password": password}
	jsonData, _ := json.Marshal(data)
	resp, err := http.Post("http://localhost:8080/login", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error logging in:", err)
		return ""
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("User logged in successfully.")
		// Extract and return the session token.
		body, _ := ioutil.ReadAll(resp.Body)
		var result map[string]string
		json.Unmarshal(body, &result)
		sessionToken := result["sessionToken"]
		return sessionToken
	}

	fmt.Println("User login failed.")
	return ""
}
