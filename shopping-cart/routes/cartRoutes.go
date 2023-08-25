package routes

import (
	"ecomm-app/shopping-cart/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router, cartController *controllers.CartController) {
	r.HandleFunc("/add-to-cart", cartController.AddItemToCart).Methods("POST")
	r.HandleFunc("/remove-from-cart", cartController.RemoveItemFromCart).Methods("POST")
	r.HandleFunc("/update-cart-item", cartController.UpdateCartItemQuantity).Methods("POST")
	r.HandleFunc("/get-cart", cartController.GetCart).Methods("GET")
}
