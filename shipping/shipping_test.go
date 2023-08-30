package main

import (
	"ecomm-app/shipping/models"
	"testing"
)

func TestCalculateShippingCost(t *testing.T) {
	// Define a set of test cases with different shipping options and expected costs
	testCases := []struct {
		name         string       // Name of the test case
		order        models.Order // Order configuration for testing
		expectedCost float64      // Expected shipping cost
	}{
		// Test case 1
		{
			name: "Standard Shipping",
			order: models.Order{
				Weight:         5.0,
				ShippingOption: "standard",
			},
			expectedCost: 2.5, // Expected cost: 5.0 * 0.5
		},
		// Test case 2
		{
			name: "Expedited Shipping",
			order: models.Order{
				Weight:         5.0,
				ShippingOption: "expedited",
			},
			expectedCost: 3.75, // Expected cost: 5.0 * 0.5 * 1.5
		},
	}

	// Iterate through each test case
	for _, tc := range testCases {
		// Run a subtest with the test case name
		t.Run(tc.name, func(t *testing.T) {
			// Calculate the actual shipping cost using the tested function
			actualCost := models.CalculateShippingCost(tc.order)

			// Compare the actual and expected costs
			if actualCost != tc.expectedCost {
				// If they don't match, generate an error message
				t.Errorf("Expected cost: %.2f, but got: %.2f", tc.expectedCost, actualCost)
			}
		})
	}
}
