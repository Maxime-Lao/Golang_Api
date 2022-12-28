package payment

type Payment struct {
	Id        int    `json:"id"`
	ProductID int    `json:"product_id"`
	PricePaid string `json:"price_paid"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
