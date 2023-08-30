package routes

import (
	"ecomm-app/shipping/controllers"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/calculate_shipping_cost", controllers.CalculateShippingCostHandler)
}
