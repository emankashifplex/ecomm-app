package routes

import (
	"ecomm-app/email/controllers"

	"github.com/gin-gonic/gin"
)

// Initializes the routes for the application
func InitRoutes(r *gin.Engine) {
	r.POST("/email", controllers.SendOrderConfirmationEmailController)
}
