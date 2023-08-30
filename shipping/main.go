package main

import (
	"ecomm-app/shipping/routes"
	"fmt"
	"net/http"
)

func main() {
	routes.SetupRoutes()

	port := 8084
	fmt.Printf("Server is running on port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
