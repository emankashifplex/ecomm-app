package routes

import (
	"ecomm-app/product-catalog/controllers"
	"ecomm-app/product-catalog/models"

	"github.com/go-pg/pg/v10"
	"github.com/gorilla/mux"
)

func SetProductRoutes(router *mux.Router, db *pg.DB) {
	productController := controllers.NewProductController(models.NewProductService(db))

	router.HandleFunc("/products", productController.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id:[0-9]+}", productController.GetProductByID).Methods("GET")
	router.HandleFunc("/products/search", productController.SearchAndFilterProducts).Methods("GET")
	router.HandleFunc("/availability", productController.CheckProductAvailability).Methods("GET")
}
