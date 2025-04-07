package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/google/uuid"
	"ikea-order-tracker/models"
	"ikea-order-tracker/storage"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Invalid order data", http.StatusBadRequest)
		return
	}
	order.ID = uuid.New().String()
	order.Status = "pending"
	storage.AddOrder(order)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(storage.GetAllOrders())
}
