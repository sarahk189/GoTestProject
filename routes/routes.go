package routes

import (
	"net/http"
	"ikea-order-tracker/handlers"
)

func RegisterRoutes() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Order Tracker API is healthy!"))
	})
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			handlers.GetOrders(w, r)
		case "POST":
			handlers.CreateOrder(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}
