package cli

import (
	"fmt"

	"github.com/vmlellis/go-hexagonal/application/contract"
)

func Run(service contract.ProductServiceInterface, action string, productID string, productName string, price float64) (string, error) {
	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf(
			"Product ID %s with the name %s has been created with the the price %f and status %s.",
			product.GetId(), product.GetName(), product.GetPrice(), product.GetStatus(),
		)
	case "enable":
		product, err := service.Get(productID)
		if err != nil {
			return result, err
		}
		product, err = service.Enable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s has been enabled.", product.GetName())
	case "disable":
		product, err := service.Get(productID)
		if err != nil {
			return result, err
		}
		product, err = service.Disable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s has been disabled.", product.GetName())
	default:
		product, err := service.Get(productID)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf(
			"Product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
			product.GetId(), product.GetName(), product.GetPrice(), product.GetStatus(),
		)
	}

	return result, nil
}
