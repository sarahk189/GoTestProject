package models

// Order represents an IKEA order in the system
// The struct tags define how the fields are serialized to/from JSON
type Order struct {
	ID       string `json:"id"`       // Unique identifier for the order
	Product  string `json:"product"`  // Name or identifier of the IKEA product
	Quantity int    `json:"quantity"` // Number of items ordered
	Status   string `json:"status"`   // Current status of the order (pending, shipped, cancelled)
}
