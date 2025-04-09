package handlers

import (
    "encoding/json"
    "net/http"
    "ikea-order-tracker/models"
    "ikea-order-tracker/services"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
    var order models.Order
    err := json.NewDecoder(r.Body).Decode(&order)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    services.CreateOrder(&order)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(order)
}
