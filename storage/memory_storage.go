package storage

import "ikea-order-tracker/models"

var orders []models.Order

func AddOrder(order models.Order) {
	orders = append(orders, order)
}

func GetAllOrders() []models.Order {
	return orders
}
