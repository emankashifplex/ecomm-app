package routes

import (
	"ecomm-app/users-service/controllers"

	"github.com/gorilla/mux"
)

// Set up the user-related routes using the provided router
func SetUserRoutes(router *mux.Router) {

	router.HandleFunc("/register", controllers.RegisterUserHandler).Methods("POST")
	router.HandleFunc("/login", controllers.LoginHandler).Methods("POST")
	router.HandleFunc("/profile/{username}", controllers.UserProfileHandler).Methods("GET")
	router.HandleFunc("/get-product-info/{productID}", controllers.GetProductInfo).Methods("GET")
	router.HandleFunc("/get-cart", controllers.GetCartInfo).Methods("GET")
}
