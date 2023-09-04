package routes

import (
	"ecomm-app/email/controllers"

	"github.com/gin-gonic/gin"
)

// InitRoutes initializes the routes for the application.
func InitRoutes(r *gin.Engine) {
	r.POST("/email", controllers.SendOrderConfirmationEmailController)
}
