package product

type InputProduct struct {
	Name  string `json:"name" binding:"required"`
	Price string `json:"price" binding:"required"`
}
