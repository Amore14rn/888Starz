package service

import (
	"context"
	"github.com/Amore14rn/888Starz/internal/domain/product/model"
)

type repository interface {
	Create(ctx context.Context, req model.CreateProduct) error
	All(ctx context.Context) ([]model.Product, error)
	GetProductByID(ctx context.Context, id string) (model.Product, error)
}

type ProductService struct {
	repository repository
}

func NewProductService(repository repository) *ProductService {
	return &ProductService{repository: repository}
}
