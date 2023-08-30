package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"ecomm-app/product-catalog/models"

	"github.com/gorilla/mux"
)

// ProductController handles HTTP requests related to products
type ProductController struct {
	ProductService *models.ProductService
}

// Constructor NewProductController creates a new ProductController instance
func NewProductController(productService *models.ProductService) *ProductController {
	return &ProductController{
		ProductService: productService,
	}
}

// CreateProduct handles the creation of a new product
func (c *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	// Decode the incoming JSON request into a Product struct.
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Call the ProductService to create the product in the database
	err = c.ProductService.CreateProduct(&product)
	if err != nil {
		http.Error(w, "Error creating product", http.StatusInternalServerError)
		return
	}

	// Respond with a status code indicating successful creation
	w.WriteHeader(http.StatusCreated)
}

// GetProductByID retrieves a product by its ID
func (c *ProductController) GetProductByID(w http.ResponseWriter, r *http.Request) {
	// Extract the product ID from the URL parameters
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// Call the ProductService to retrieve the product by ID from the database
	product, err := c.ProductService.GetProductByID(id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	// Respond with the retrieved product as a JSON-encoded response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// SearchProducts searches for products based on a query string
func (c *ProductController) SearchProducts(w http.ResponseWriter, r *http.Request) {
	// Extract the search query from the URL query parameters
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "Missing query parameter", http.StatusBadRequest)
		return
	}

	// Call the ProductService to search for products using the given query
	products, err := c.ProductService.SearchProducts(query)
	if err != nil {
		http.Error(w, "Error searching products", http.StatusInternalServerError)
		return
	}

	// Respond with the search results as a JSON-encoded response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// Filters products based on given criteria
func (c *ProductController) FilterProducts(w http.ResponseWriter, r *http.Request) {
	// Extract filter criteria from the URL query parameters
	minPriceStr := r.URL.Query().Get("minPrice")
	maxPriceStr := r.URL.Query().Get("maxPrice")
	availabilityStr := r.URL.Query().Get("availability")

	// Convert filter criteria to appropriate data types
	minPrice, _ := strconv.ParseFloat(minPriceStr, 64)
	maxPrice, _ := strconv.ParseFloat(maxPriceStr, 64)
	availability, _ := strconv.ParseBool(availabilityStr)

	// Call the ProductService to filter products based on the provided criteria
	products, err := c.ProductService.FilterProducts(minPrice, maxPrice, availability)
	if err != nil {
		http.Error(w, "Error filtering products", http.StatusInternalServerError)
		return
	}

	// Respond with the filtered products as a JSON-encoded response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
