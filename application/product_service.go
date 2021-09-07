package application

import "github.com/vmlellis/go-hexagonal/application/contract"

type productService struct {
	Persistence contract.ProductPersistenceInterface
}

func NewProductService(persistence contract.ProductPersistenceInterface) contract.ProductServiceInterface {
	return &productService{Persistence: persistence}
}

func (s *productService) Get(id string) (contract.ProductInterface, error) {
	product, err := s.Persistence.Get(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *productService) Create(name string, price float64) (contract.ProductInterface, error) {
	product := NewProduct(name, price)
	_, err := product.IsValid()
	if err != nil {
		return product, err
	}
	result, err := s.Persistence.Save(product)
	if err != nil {
		return product, err
	}
	return result, nil
}

func (s *productService) Enable(product contract.ProductInterface) (contract.ProductInterface, error) {
	err := product.Enable()
	if err != nil {
		return product, err
	}

	result, err := s.Persistence.Save(product)
	if err != nil {
		return product, err
	}

	return result, err
}

func (s *productService) Disable(product contract.ProductInterface) (contract.ProductInterface, error) {
	err := product.Disable()
	if err != nil {
		return product, err
	}

	result, err := s.Persistence.Save(product)
	if err != nil {
		return product, err
	}

	return result, err
}
