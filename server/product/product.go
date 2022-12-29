package product

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
