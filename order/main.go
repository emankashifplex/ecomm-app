package main

import (
	"database/sql"
	"ecomm-app/order/routes"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://eman:123@localhost:5432/product_data?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8083"
	}

	r := routes.SetOrderRoutes(db)

	log.Printf("Server started on :%s\n", port)
	http.ListenAndServe(":"+port, r)
}
