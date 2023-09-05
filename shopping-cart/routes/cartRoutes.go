package routes

import (
	"ecomm-app/shopping-cart/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router, cartController *controllers.CartController) {
	r.HandleFunc("/add", cartController.AddItemToCart).Methods("POST")
	r.HandleFunc("/remove", cartController.RemoveItemFromCart).Methods("POST")
	r.HandleFunc("/update", cartController.UpdateCartItemQuantity).Methods("POST")
	r.HandleFunc("/get", cartController.GetCart).Methods("GET")
}
