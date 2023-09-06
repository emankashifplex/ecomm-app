package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {
	// Create a sample product
	product := Product{
		ID:           1,
		Name:         "Sample Product",
		Description:  "This is a sample product",
		Price:        9.99,
		Availability: true,
	}

	// Check if the fields of the product match the expected values
	assert.Equal(t, 1, product.ID, "Expected product ID to be 1")
	assert.Equal(t, "Sample Product", product.Name, "Expected product name to be 'Sample Product'")
	assert.Equal(t, "This is a sample product", product.Description, "Expected product description to be 'This is a sample product'")
	assert.Equal(t, 9.99, product.Price, "Expected product price to be 9.99")
	assert.True(t, product.Availability, "Expected product availability to be true")
}
