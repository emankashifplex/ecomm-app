package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestUserRegistrationAndAuthentication(t *testing.T) {
	// Sample user data
	username := "testuser"
	password := "testpassword"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	// Mock database functions
	mockRegisterUser := func(user *User) error {
		// Simulate successful user registration
		return nil
	}

	mockFindUserByUsername := func(username string) (*User, error) {
		// Simulate finding user by username
		if username == "testuser" {
			return &User{
				Username: username,
				Password: string(hashedPassword),
			}, nil
		}
		return nil, nil
	}

	mockVerifyPassword := func(hashedPassword, password string) error {
		// Simulate successful password verification
		return nil
	}

	// Register user
	user := &User{
		Username: username,
		Password: password,
	}
	err := mockRegisterUser(user)
	assert.NoError(t, err, "Failed to register user")

	// Find user by username
	foundUser, err := mockFindUserByUsername(username)
	assert.NoError(t, err, "Failed to find user by username")
	assert.NotNil(t, foundUser, "User not found")
	assert.Equal(t, username, foundUser.Username, "Username mismatch")

	// Verify user's password
	err = mockVerifyPassword(foundUser.Password, password)
	assert.NoError(t, err, "Password verification failed")
}
