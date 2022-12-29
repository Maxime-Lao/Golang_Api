package payment

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Create(payment Payment) (Payment, error)
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

func (r *repository) Create(payment Payment) (Payment, error) {
	err := r.db.Create(&payment).Error
	if err != nil {
		return payment, err
	}

	return payment, nil
}

func (r *repository) GetAll() ([]Payment, error) {
	var payments []Payment
	err := r.db.Find(&payments).Error
	if err != nil {
		return payments, err
	}

	return payments, nil
}

func (r *repository) GetById(id int) (Payment, error) {
	var payment Payment

	err := r.db.Where(&Payment{ID: id}).First(&payment).Error
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

	payment.ProductID = inputPayment.ProductID
	payment.PricePaid = inputPayment.PricePaid

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
