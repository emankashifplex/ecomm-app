package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateShippingCost(t *testing.T) {
	// Create an example order
	order := Order{
		ID:             1,
		RecipientAddr:  "123 Main St",
		Weight:         3.0,
		ShippingOption: "expedited",
	}

	// Calculate the expected shipping cost
	expectedCost := 3.0 * 0.5 * 1.5

	// Calculate the actual shipping cost using the function
	actualCost := CalculateShippingCost(order)

	// Use the assertion library to compare expected and actual values
	assert.Equal(t, expectedCost, actualCost, "Shipping costs do not match for expedited option")
}

// Other test cases for standard and overnight shipping options can be written similarly.
