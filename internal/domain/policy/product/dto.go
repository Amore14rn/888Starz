package product

import (
	"github.com/Amore14rn/888Starz/internal/domain/product/model"
)

type CreateProductInput struct {
	ID          string
	Description string
	Tags        []string
	Quantity    int
	History     model.ProductHistory
}

func NewProductInput(description string, tags []string, quantity int, history model.ProductHistory) CreateProductInput {
	return CreateProductInput{
		Description: description,
		Tags:        tags,
		Quantity:    quantity,
		History:     history,
	}
}

type CreateProductOutput struct {
	Product model.Product
}
