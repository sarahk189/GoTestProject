package storage

import "ikea-order-tracker/models"

// In-memory storage for orders
// Note: This is a simple implementation that will be cleared when the server restarts
var orders []models.Order

// AddOrder adds a new order to the in-memory storage
func AddOrder(order models.Order) {
	orders = append(orders, order)
}

// GetAllOrders returns all orders currently stored in memory
func GetAllOrders() []models.Order {
	return orders
}
