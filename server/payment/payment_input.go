package payment

type InputPayment struct {
	ProductID int     `json:"product_id" binding:"required"`
	PricePaid float64 `json:"price_paid" binding:"required"`
}
