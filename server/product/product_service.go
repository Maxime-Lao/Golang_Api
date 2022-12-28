package product

type Service interface {
	Create(input InputProduct) (Product, error)
	GetAll() ([]Product, error)
	GetById(id int) (Product, error)
	Update(id int, inputProduct InputProduct) (Product, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) Create(input InputProduct) (Product, error) {
	var product Product
	product.Name = input.Name
	product.Price = input.Price

	newProduct, err := s.repository.Create(product)
	if err != nil {
		return product, err
	}

	return newProduct, nil
}

func (s *service) GetAll() ([]Product, error) {
	products, err := s.repository.GetAll()
	if err != nil {
		return products, err
	}

	return products, nil
}

func (s *service) GetById(id int) (Product, error) {
	product, err := s.repository.GetById(id)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *service) Update(id int, input InputProduct) (Product, error) {
	uProduct, err := s.repository.Update(id, input)
	if err != nil {
		return uProduct, err
	}

	return uProduct, nil
}

func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
