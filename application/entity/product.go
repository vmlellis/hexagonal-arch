package entity

import (
	"errors"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"github.com/vmlellis/go-hexagonal/application/contract"
	"github.com/vmlellis/go-hexagonal/application/dto"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type product struct {
	ID     string  `valid:"uuidv4,optional"`
	Name   string  `valid:"required"`
	Price  float64 `valid:"float,optional"`
	Status string  `valid:"required"`
}

func NewProduct(name string, price float64) contract.ProductInterface {
	return &product{
		ID:     uuid.NewV4().String(),
		Status: DISABLED,
		Name:   name,
		Price:  price,
	}
}

func RegisterProduct(productDto dto.ProductDto) contract.ProductInterface {
	return &product{
		ID:     productDto.ID,
		Status: productDto.Status,
		Name:   productDto.Name,
		Price:  productDto.Price,
	}
}

func (p *product) IsValid() (bool, error) {
	if p.Status == "" {
		p.Status = DISABLED
	}

	if p.Status != ENABLED && p.Status != DISABLED {
		return false, errors.New("the status must be enabled or disabled")
	}

	if p.Price < 0 {
		return false, errors.New("the price must be greater or equal than zero")
	}

	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}
	return errors.New("the price must be greater than zero to enable the product")
}

func (p *product) Disable() error {
	if p.Price == 0 {
		p.Status = DISABLED
		return nil
	}
	return errors.New("the price must be zero in order to have the product disabled")
}

func (p *product) GetId() string {
	return p.ID
}

func (p *product) GetName() string {
	return p.Name
}

func (p *product) GetStatus() string {
	return p.Status
}

func (p *product) GetPrice() float64 {
	return p.Price
}
