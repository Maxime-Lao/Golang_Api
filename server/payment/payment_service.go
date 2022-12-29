package payment

type Service interface {
	Create(input InputPayment) (Payment, error)
	GetAll() ([]Payment, error)
	GetById(id int) (Payment, error)
	Update(id int, inputPayment InputPayment) (Payment, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) Create(input InputPayment) (Payment, error) {
	var payment Payment
	payment.ProductID = input.ProductID
	payment.PricePaid = input.PricePaid

	newPayment, err := s.repository.Create(payment)
	if err != nil {
		return payment, err
	}

	return newPayment, nil
}

func (s *service) GetAll() ([]Payment, error) {
	payments, err := s.repository.GetAll()
	if err != nil {
		return payments, err
	}

	return payments, nil
}

func (s *service) GetById(id int) (Payment, error) {
	payment, err := s.repository.GetById(id)
	if err != nil {
		return payment, err
	}

	return payment, nil
}

func (s *service) Update(id int, input InputPayment) (Payment, error) {
	uPayment, err := s.repository.Update(id, input)
	if err != nil {
		return uPayment, err
	}

	return uPayment, nil
}

func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
