package payment

type InputPayment struct {
	ProductName string `json:"productName" binding:"required"`
}
