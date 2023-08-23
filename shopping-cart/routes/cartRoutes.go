package routes

import (
	"ecomm-app/shopping-cart/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/add-to-cart", controllers.AddToCart).Methods("POST")
	router.HandleFunc("/remove-from-cart/{productID}", controllers.RemoveFromCart).Methods("DELETE")
	router.HandleFunc("/update-quantity/{productID}", controllers.UpdateCartItemQuantity).Methods("PUT")
	router.HandleFunc("/calculate-total-price", controllers.CalculateTotalPrice).Methods("GET")
}
