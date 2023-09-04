package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCart(t *testing.T) {
	// Create a new cart
	cart := NewCart()

	// Add an item to the cart
	productID := 1
	quantity := 2
	price := 10.99
	cart.AddItem(productID, quantity, price)

	// Check if the cart has the added item
	assert.Len(t, cart.Items, 1, "Cart should have one item")
	assert.Equal(t, productID, cart.Items[0].ProductID, "Product IDs do not match")
	assert.Equal(t, quantity, cart.Items[0].Quantity, "Quantities do not match")
	assert.Equal(t, price, cart.Items[0].Price, "Prices do not match")
	assert.Equal(t, price*float64(quantity), cart.Total, "Total cost is incorrect")

	// Update the quantity of the added item
	newQuantity := 3
	cart.UpdateQuantity(productID, newQuantity)
	assert.Equal(t, newQuantity, cart.Items[0].Quantity, "Updated quantities do not match")
	assert.Equal(t, price*float64(newQuantity), cart.Total, "Total cost after update is incorrect")

	// Remove the item from the cart
	cart.RemoveItem(productID)
	assert.Empty(t, cart.Items, "Cart should be empty after removing item")
	assert.Equal(t, 0.0, cart.Total, "Total cost should be zero after removing all items")
}
