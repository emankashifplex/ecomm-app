package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// Represents a user in the system
type User struct {
	ID       string `bson:"_id,omitempty"`
	Username string `bson:"username"`
	Password string `bson:"password"`
}

// Hashes the user's password using bcrypt
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// Adds a new user to the database
func RegisterUser(db *mongo.Database, user *User) error {
	usersCollection := db.Collection("users")
	_, err := usersCollection.InsertOne(context.TODO(), user)
	return err
}

// Retrieves a user from the database by their username
func FindUserByUsername(db *mongo.Database, username string) (*User, error) {
	usersCollection := db.Collection("users")
	var user User
	err := usersCollection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Compares a hashed password with a plaintext password
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// CheckUserExistsInDatabase checks if a user with the given userID exists in the database.
func DoesUserExist(db *mongo.Database, userID int) (bool, error) {
	usersCollection := db.Collection("users")

	// Create a filter to find a user by their ID
	filter := bson.M{"_id": userID}

	// Use the CountDocuments method to check if a user with the provided ID exists
	count, err := usersCollection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return false, err
	}

	// If count is greater than 0, a user with the provided ID exists
	return count > 0, nil
}
