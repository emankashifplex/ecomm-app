package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"ecomm-app/shipping/models"
)

func TestCalculateShippingCostHandler(t *testing.T) {
	// Create a sample order for testing
	order := models.Order{
		ID:             1,
		RecipientAddr:  "123 Main St",
		Weight:         3.0,
		ShippingOption: "expedited",
	}

	// Marshal the order into JSON
	orderJSON, err := json.Marshal(order)
	assert.NoError(t, err, "Error marshaling order to JSON")

	// Create a request with the JSON payload
	req, err := http.NewRequest("POST", "/calculate-shipping", bytes.NewReader(orderJSON))
	assert.NoError(t, err, "Error creating request")

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the CalculateShippingCostHandler
	CalculateShippingCostHandler(rr, req)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, rr.Code, "Unexpected response status code")

	// Parse the response body and check if it contains the expected content
	expectedResponse := fmt.Sprintf("Shipping cost: $%.2f", 3.0*0.5*1.5)
	assert.Equal(t, expectedResponse, rr.Body.String(), "Response content mismatch")
}
