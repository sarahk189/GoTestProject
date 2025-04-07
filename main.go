package main

import (
	"fmt"
	"net/http"
	"ikea-order-tracker/routes"
)

func main() {
	routes.RegisterRoutes()
	fmt.Println("🚀 Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
