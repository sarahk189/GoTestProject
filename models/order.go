package models

type Order struct {
	ID       string `json:"id"`
	Product  string `json:"product"`
	Quantity int    `json:"quantity"`
	Status   string `json:"status"` // e.g., "pending", "shipped", "cancelled"
}
