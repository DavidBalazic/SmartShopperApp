package services

import (
	"context"
	"github.com/DavidBalazic/SmartShopperApp/internal/repo"
	"github.com/DavidBalazic/SmartShopperApp/internal/models"
)

type ProductService interface {
	GetCheapestProduct(ctx context.Context, name string) (models.Product, error)
	GetCheapestByStore(ctx context.Context, name, store string) (models.Product, error)
	GetAllPrices(ctx context.Context, name string) ([]models.Product, error)
	GetProductById(ctx context.Context, id string) (models.Product, error)
	AddProduct(ctx context.Context, product models.Product) (models.Product, error)
}

type productService struct {
	repo repo.ProductRepository
}

func NewProductService(repo repo.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) GetCheapestProduct(ctx context.Context, name string) (models.Product, error) {
	product, err := s.repo.FindCheapestProduct(ctx, name)
	return product, err
}

func (s *productService) GetCheapestByStore(ctx context.Context, name, store string) (models.Product, error) {
	product, err := s.repo.FindCheapestProductByStore(ctx, name, store)
	return product, err
}

func (s *productService) GetAllPrices(ctx context.Context, name string) ([]models.Product, error) {
	products, err := s.repo.FindAllProductPrices(ctx, name)
	return products, err
}

func (s *productService) GetProductById(ctx context.Context, id string) (models.Product, error) {
	product, err := s.repo.FindProductById(ctx, id)
	return product, err
}

func (s *productService) AddProduct(ctx context.Context, product models.Product) (models.Product, error) {
	product, err := s.repo.AddProduct(ctx, product)
	return product, err
}