package models

import (
	"testing"
	"time"
)

func TestOrder(t *testing.T) {
	// Create a sample order
	order := Order{
		ID:        1,
		Product:   "Sample Product",
		Quantity:  5,
		Status:    string(StatusPending),
		CreatedAt: time.Now(),
	}

	// Check if the fields of the order match the expected values
	if order.ID != 1 {
		t.Errorf("Expected order ID to be 1, but got %d", order.ID)
	}
	if order.Product != "Sample Product" {
		t.Errorf("Expected product name to be 'Sample Product', but got '%s'", order.Product)
	}
	if order.Quantity != 5 {
		t.Errorf("Expected order quantity to be 5, but got %d", order.Quantity)
	}
	if order.Status != string(StatusPending) {
		t.Errorf("Expected order status to be 'Pending', but got '%s'", order.Status)
	}
	if order.CreatedAt.IsZero() {
		t.Error("Expected order creation time to be set, but it's zero")
	}
}

func TestOrderStatusConstants(t *testing.T) {
	// Check if the order status constants have the correct values
	if StatusPending != "Pending" {
		t.Errorf("Expected StatusPending to be 'Pending', but got '%s'", StatusPending)
	}
	if StatusProcessed != "Processed" {
		t.Errorf("Expected StatusProcessed to be 'Processed', but got '%s'", StatusProcessed)
	}
	if StatusFulfilled != "Fulfilled" {
		t.Errorf("Expected StatusFulfilled to be 'Fulfilled', but got '%s'", StatusFulfilled)
	}
}
