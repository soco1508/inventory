package service

import (
	"backend/internal/db/models"
	"backend/internal/db/repository"
	"context"
)

type ProductService interface {
	GetPopularProducts(ctx context.Context) ([]*models.Product, error)
}

type productService struct {
	productRepo repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return &productService{productRepo: productRepo}
}

func (p *productService) GetPopularProducts(ctx context.Context) ([]*models.Product, error) {
	return p.productRepo.GetPopularProducts(ctx)
}
