package domain

import (
	"github.com/vmlellis/go-hexagonal/application/contract"
	"github.com/vmlellis/go-hexagonal/application/dto"
	"github.com/vmlellis/go-hexagonal/application/entity"
)

type ProductRequest struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Status string  `json:"status"`
}

func (req *ProductRequest) Bind() (contract.ProductInterface, error) {
	productDto := dto.ProductDto{
		ID:     req.ID,
		Name:   req.Name,
		Price:  req.Price,
		Status: req.Status,
	}

	product := entity.RegisterProduct(productDto)
	_, err := product.IsValid()
	if err != nil {
		return nil, err
	}
	return product, nil
}
