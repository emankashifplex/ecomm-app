package main

import (
	"fmt"
	"net/http"
)

// Send a confirmation email for an order.
func sendConfirmationEmail(userID, orderID int) {
	url := fmt.Sprintf("http://localhost:8085/email?user_id=%d&order_id=%d", userID, orderID)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error sending confirmation email:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Confirmation email sent successfully.")
	} else {
		fmt.Println("Sending confirmation email failed.")
	}
}
