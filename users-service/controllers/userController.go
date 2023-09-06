package controllers

import (
	"context"
	"ecomm-app/users-service/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Handles user registration requests
func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Hash the user's password before storing it in the database
	err = user.HashPassword()
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	// Connect to the MongoDB database.
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.TODO())

	// Access the "ecommerce" database.
	db := client.Database("ecommerce")

	// Register the user in the database.
	err = models.RegisterUser(db, &user)
	if err != nil {
		http.Error(w, "Error registering user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// LoginHandler handles user login requests.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Connect to the MongoDB database.
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.TODO())

	// Access the "ecommerce" database.
	db := client.Database("ecommerce")

	// Find the user by username in the database.
	user, err := models.FindUserByUsername(db, creds.Username)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Verify the provided password against the stored password hash.
	err = models.VerifyPassword(user.Password, creds.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Return a success message if login is successful.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})
}

// UserProfileHandler retrieves and returns user profile information.
func UserProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the username from the request's URL path.
	vars := mux.Vars(r)
	username := vars["username"]

	// Connect to the MongoDB database.
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.TODO())

	// Access the "ecommerce" database.
	db := client.Database("ecommerce")

	// Find the user by username in the database.
	user, err := models.FindUserByUsername(db, username)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Return the user profile as JSON response.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Check if a user exists by userID
func CheckUserExistenceHandler(w http.ResponseWriter, r *http.Request) {
	// Extract userID from the request parameters
	userIDStr := r.URL.Query().Get("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid userID", http.StatusBadRequest)
		return
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	db := client.Database("ecommerce")
	// Check if the user exists
	exists, err := models.DoesUserExist(db, userID)

	// Respond with true if the user exists, false if not
	if exists {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "true")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "false")
	}
}
