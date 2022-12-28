package product

import "gorm.io/gorm"

type Repository interface {
	Insert(product Product) (Product, error)
	ListAll() ([]Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Insert(product Product) (Product, error) {
	err := r.db.Create(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) ListAll() ([]Product, error) {
	var products []Product
	err := r.db.Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}
