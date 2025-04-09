package services

import (
    "github.com/google/uuid"
    "ikea-order-tracker/models"
    "ikea-order-tracker/storage"
)

func CreateOrder(order *models.Order) {
    order.ID = uuid.New().String()
    order.Status = "pending"
    storage.AddOrder(*order)
}
