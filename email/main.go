package main

import (
	"ecomm-app/email/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Initialize routes
	routes.InitRoutes(r)

	// Start the Gin server
	r.Run(":8085")
}
