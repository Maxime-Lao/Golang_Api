package product

// on définit notre classe
type Product struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Price     string `json:"price"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
