package product

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Create(product Product) (Product, error)
	GetAll() ([]Product, error)
	GetById(id int) (Product, error)
	Update(id int, inputProduct InputProduct) (Product, error)
	Delete(id int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(product Product) (Product, error) {

    uncite := r.db.Where(&Product{Name: product.Name}).First(&product).Error
    if uncite != nil {
        err := r.db.Create(&product).Error
        if err != nil {
            return product, err
        }
    } else {
        return product, errors.New("The Product name already exist")
    }

	return product, nil
}

func (r *repository) GetAll() ([]Product, error) {
	var products []Product
	err := r.db.Find(&products).Error
	if err != nil {
		return products, err
	}

	return products, nil
}

func (r *repository) GetById(id int) (Product, error) {
	var product Product

	err := r.db.Where(&Product{ID: id}).First(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) Update(id int, inputProduct InputProduct) (Product, error) {


	product, err := r.GetById(id)
	if err != nil {
		return product, err
	}

	product.Name = inputProduct.Name
	product.Price = inputProduct.Price

	uncite := r.db.Where("name = ? AND price = ?", product.Name, product.Price).First(&product).Error
    if uncite != nil {
		err = r.db.Save(&product).Error
		if err != nil {
			return product, err
		}
	} else {
		return product, errors.New("The Product name already exist")
	}

	return product, nil
}

func (r *repository) Delete(id int) error {
	product := &Product{ID: id}
	tx := r.db.Delete(product)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("Product not found")
	}

	return nil
}
