package payment

import (
	"errors"
	"go-api/server/product"

	"gorm.io/gorm"
)

type Repository interface {
	Create(productName string) (Payment, error)
	GetAll() ([]Payment, error)
	GetById(id int) (Payment, error)
	Update(id int, inputPayment InputPayment) (Payment, error)
	Delete(id int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(productName string) (Payment, error) {
	var product product.Product
	r.db.Where("name = ?", productName).First(&product)

	var payment Payment
	payment.PricePaid = product.Price
	payment.ProductID = product.ID
	payment.Product = product

	err := r.db.Create(&payment).Error
	if err != nil {
		return payment, err
	}

	return payment, nil
}

func (r *repository) GetAll() ([]Payment, error) {
	var payments []Payment
	err := r.db.Find(&payments).Error

	for index, payment := range payments {
		r.db.Where(&product.Product{ID: payment.ProductID}).First(&payments[index].Product)
	}

	if err != nil {
		return payments, err
	}

	return payments, nil
}

func (r *repository) GetById(id int) (Payment, error) {
	var payment Payment

	err := r.db.Where(&Payment{ID: id}).First(&payment).Error

	r.db.Where(&product.Product{ID: payment.ProductID}).First(&payment.Product)
	if err != nil {
		return payment, err
	}

	return payment, nil
}

func (r *repository) Update(id int, inputPayment InputPayment) (Payment, error) {
	payment, err := r.GetById(id)
	if err != nil {
		return payment, err
	}

	var product product.Product
	r.db.Where("name = ?", inputPayment.ProductName).First(&product)

	if payment.ProductID == product.ID {
		return payment, nil
	}

	payment.PricePaid = product.Price
	payment.Product = product

	err = r.db.Save(&payment).Error
	if err != nil {
		return payment, err
	}

	return payment, nil
}

func (r *repository) Delete(id int) error {
	payment := &Payment{ID: id}
	tx := r.db.Delete(payment)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("Payment not found")
	}

	return nil
}
