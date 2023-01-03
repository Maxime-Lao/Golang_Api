package payment

import (
	"time"
)

type Payment struct {
	ID        int       `json:"id"`
	ProductID int       `json:"product_id"`
	PricePaid float64   `json:"price_paid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
