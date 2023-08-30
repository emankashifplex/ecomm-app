package main

import (
	"ecomm-app/product-catalog/models"
	"os"
	"testing"

	"github.com/go-pg/pg/v10"
	"github.com/stretchr/testify/require"
)

func TestProductService_CreateProduct(t *testing.T) {
	// Environment variable containing database conn fetched
	databaseURL := os.Getenv("TEST_DATABASE_URL")
	require.NotEmpty(t, databaseURL, "TEST_DATABASE_URL environment variable is required")

	// The PostgreSQL connection options are parsed from the databaseURL
	opt, err := pg.ParseURL(databaseURL)
	require.NoError(t, err, "Error parsing test database URL")

	// Connection to the test database is established
	db := pg.Connect(opt)
	defer db.Close()

	// Create a new ProductService instance using the test database connection
	service := models.NewProductService(db)

	// Clean up existing test data before inserting new data
	_, err = db.Exec("DELETE FROM products")
	require.NoError(t, err, "Error cleaning up test data")

	// Create a sample product for testing
	product := &models.Product{
		Name:         "Test Product",
		Description:  "Test description",
		Price:        29.99,
		Availability: true,
	}

	// Call the CreateProduct method
	err = service.CreateProduct(product)
	require.NoError(t, err, "Creating product should not return an error")

	// Fetch the product from the database and ensure it matches the created product
	fetchedProduct, err := service.GetProductByID(product.ID)
	require.NoError(t, err, "Fetching product should not return an error")
	require.Equal(t, product.Name, fetchedProduct.Name, "Product name should match")
	require.Equal(t, product.Description, fetchedProduct.Description, "Product description should match")
	require.Equal(t, product.Price, fetchedProduct.Price, "Product price should match")
	require.Equal(t, product.Availability, fetchedProduct.Availability, "Product availability should match")
}
