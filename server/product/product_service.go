package product

type Service interface {
	Store(input InputProduct) (Product, error)
	ListAll() ([]Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Store(input InputProduct) (Product, error) {
	product := Product{}
	product.Name = input.Name
	product.Price = input.Price

	newProduct, err := s.repository.Insert(product)
	if err != nil {
		return newProduct, err
	}

	return newProduct, nil
}

func (s *service) ListAll() ([]Product, error) {
	products, err := s.repository.ListAll()
	if err != nil {
		return products, err
	}

	return products, nil
}
