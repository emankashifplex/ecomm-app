package routes

import (
	"ecomm-app/shipping/controllers"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/calculateshippingcost", controllers.CalculateShippingCostHandler)
}
