package services

import (
	"github.com/DavidBalazic/SmartShopperApp/repo"
	"github.com/DavidBalazic/SmartShopperApp/models"
)

type ProductService struct {
	repo repo.ProductRepository
}

func NewProductService(repo repo.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetCheapestProduct(name string) (models.Product, error) {
	product, err := s.repo.FindCheapestProduct(name)
	return product, err
}

func (s *ProductService) GetCheapestByStore(name, store string) (models.Product, error) {
	product, err := s.repo.FindCheapestProductByStore(name, store)
	return product, err
}

func (s *ProductService) GetAllPrices(name string) ([]models.Product, error) {
	products, err := s.repo.FindAllProductPrices(name)
	return products, err
}