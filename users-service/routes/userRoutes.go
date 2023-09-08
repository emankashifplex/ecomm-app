package routes

import (
	"ecomm-app/users-service/controllers"

	"github.com/gorilla/mux"
)

// Set up the user-related routes using the provided router
func SetUserRoutes(router *mux.Router) {

	router.HandleFunc("/register", controllers.RegisterUserHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/login", controllers.LoginHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/profile/{username}", controllers.UserProfileHandler).Methods("GET")
	router.HandleFunc("/exists/{user_id}", controllers.CheckUserExistenceHandler).Methods("GET")
}
