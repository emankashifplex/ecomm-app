package controllers

import (
	"ecomm-app/product-catalog/models"
	"testing"

	"github.com/stretchr/testify/require"
)

// interface for database operations.
type MockDatabase interface {
	CreateProduct(product *models.Product) error
	GetProductByID(id int) (*models.Product, error)
}

// MockDB implements the MockDatabase interface for testing.
type MockDB struct {
	Products map[int]*models.Product
}

func (m *MockDB) CreateProduct(product *models.Product) error {
	m.Products[product.ID] = product
	return nil
}

func (m *MockDB) GetProductByID(id int) (*models.Product, error) {
	product, ok := m.Products[id]
	if !ok {
		return nil, Errors.new("Product not found")
	}
	return product, nil
}

func TestProductService_CreateProduct(t *testing.T) {
	// Create a new ProductService instance using the mock database
	mockDB := &MockDB{
		Products: make(map[int]*models.Product),
	}
	service := models.NewProductService(mockDB)

	// Create a sample product for testing
	product := &models.Product{
		ID:           1, // Set a unique ID for testing
		Name:         "Test Product",
		Description:  "Test description",
		Price:        29.99,
		Availability: true,
	}

	// Call the CreateProduct method
	err := service.CreateProduct(product)
	require.NoError(t, err, "Creating product should not return an error")

	// Fetch the product from the mock database and ensure it matches the created product
	fetchedProduct, err := service.GetProductByID(product.ID)
	require.NoError(t, err, "Fetching product should not return an error")
	require.Equal(t, product.Name, fetchedProduct.Name, "Product name should match")
	require.Equal(t, product.Description, fetchedProduct.Description, "Product description should match")
	require.Equal(t, product.Price, fetchedProduct.Price, "Product price should match")
	require.Equal(t, product.Availability, fetchedProduct.Availability, "Product availability should match")
}
