package models

import (
	"time"
)

// Represents an order placed for a product
type Order struct {
	ID        int
	ProductID int
	UserID    int
	Quantity  int
	Status    string
	CreatedAt time.Time
}

// OrderStatus defines the possible statuses for an order.
type OrderStatus string

const (
	StatusPending   OrderStatus = "Pending"
	StatusProcessed OrderStatus = "Processed"
	StatusFulfilled OrderStatus = "Fulfilled"
)
