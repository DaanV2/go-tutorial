package inventory

import (
	"github.com/DaanV2/go-tutorial/pkg/rand"
)

// Product represents a product in the inventory
type Product struct {
	ID       string  `json:"id"`       // ID is the unique identifier of the product
	Quantity int64   `json:"quantity"` // Quantity is the number of items in stock
	Note     string  `json:"note"`     // Note is a description of the product
	Price    float64 `json:"price"`    // Price is the cost of the product
	Tax      float64 `json:"tax"`      // Tax is the tax rate for the product
	Time     string  `json:"time"`     // Time is the time the product was added
}

// GenerateProduct generates a random Product value
func GenerateProduct() Product {
	return Product{
		ID:       GetRandomId(),
		Note:     rand.RandomHex(int(rand.Int64(20))),
		Price:    rand.Float64(100),
		Tax:      rand.Float64(10),
		Time:     rand.RandomTime(),
		Quantity: rand.Int64(100),
	}
}
