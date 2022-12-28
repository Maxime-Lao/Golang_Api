package product

type InputTask struct {
	Name  string `json:"name" binding:"required"`
	Price string `json:"price" binding:"required"`
}
