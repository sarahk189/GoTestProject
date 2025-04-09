package main

import (
	"fmt"
	"ikea-order-tracker/routes"
	"net/http"
)

func main() {
	// Register all HTTP routes for the application
	routes.RegisterRoutes()
	
	// Print a startup message indicating the server is running
	fmt.Println("🚀 Server running at http://localhost:8080")
	
	// Start the HTTP server on port 8080
	// This will block until the server is stopped
	http.ListenAndServe(":8080", nil)
}
