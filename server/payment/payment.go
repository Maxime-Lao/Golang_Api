package payment

import (
	"go-api/server/product"
	"time"
)

type Payment struct {
	ID        int `json:"id"`
	ProductID int
	Product   product.Product
	PricePaid float64   `json:"price_paid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
