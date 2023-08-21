package models

import (
	"github.com/go-pg/pg/v10"
)

// Product represents a product's information
type Product struct {
	ID           int     `pg:"id,pk"`
	Name         string  `pg:"name"`
	Description  string  `pg:"description"`
	Price        float64 `pg:"price"`
	Availability bool    `pg:"availability"`
}

// ProductService provides methods to interact with products in the database
type ProductService struct {
	DB *pg.DB
}

// NewProductService creates a new ProductService instance
func NewProductService(db *pg.DB) *ProductService {
	return &ProductService{
		DB: db,
	}
}

// CreateProduct inserts a new product record into the database
func (s *ProductService) CreateProduct(product *Product) error {
	_, err := s.DB.Model(product).Insert()
	return err
}

// GetProductByID retrieves a product by its ID from the database
func (s *ProductService) GetProductByID(id int) (*Product, error) {
	// Create a new empty Product instance to hold the retrieved data.
	product := new(Product)

	// SELECT query to retrieve a product based on the given ID.
	err := s.DB.Model(product).
		Where("id = ?", id).
		Select()

	if err != nil {
		return nil, err
	}

	return product, nil
}

// SearchProducts searches for products with names containing the given query.
func (s *ProductService) SearchProducts(query string) ([]*Product, error) {
	var products []*Product
	err := s.DB.Model(&products).Where("name ILIKE ?", "%"+query+"%").Select()
	if err != nil {
		return nil, err
	}
	return products, nil
}

// FilterProducts filters products based on price range and availability
func (s *ProductService) FilterProducts(minPrice, maxPrice float64, availability bool) ([]*Product, error) {
	var products []*Product
	q := s.DB.Model(&products).
		Where("price >= ?", minPrice).
		Where("price <= ?", maxPrice)
	if availability {
		q = q.Where("availability = ?", true)
	}
	err := q.Select()
	if err != nil {
		return nil, err
	}
	return products, nil
}
