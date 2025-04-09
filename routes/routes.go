package routes

import (
	"ikea-order-tracker/handlers"
	"net/http"
)

// RegisterRoutes sets up all HTTP routes for the application
func RegisterRoutes() {
	// Health check endpoint to verify the API is running
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Order Tracker API is healthy!"))
	})

	// Orders endpoint handles both GET and POST requests
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			// Handle GET request to retrieve all orders
			handlers.GetOrders(w, r)
		case "POST":
			// Handle POST request to create a new order
			handlers.CreateOrder(w, r)
		default:
			// Return 405 Method Not Allowed for unsupported HTTP methods
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}
