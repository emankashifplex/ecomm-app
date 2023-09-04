package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Mock function to send an order confirmation email
func sendOrderConfirmationEmail(userEmail, orderID string) error {
	fmt.Printf("Sent order confirmation email to %s for order ID %s\n", userEmail, orderID)
	return nil
}

// SendOrderConfirmationEmailController sends an order confirmation email
func SendOrderConfirmationEmailController(c *gin.Context) {
	// Extract the User ID from the query parameter
	userID := c.Query("user_id")
	orderID := c.Query("order_id")

	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is missing"})
		return
	}

	userEmail := getUserEmailByID(userID)

	if userEmail == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Mock sending an order confirmation email
	err := sendOrderConfirmationEmail(userEmail, orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order confirmation email sent"})
}

// Mock function to retrieve user's email by User ID
func getUserEmailByID(userID string) string {
	return "user@example.com"
}
