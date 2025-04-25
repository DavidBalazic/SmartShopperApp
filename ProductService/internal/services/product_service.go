package services

import (
	"context"
	"errors"
	"log"

	"github.com/DavidBalazic/SmartShopperApp/internal/models"
	"github.com/DavidBalazic/SmartShopperApp/internal/repo"
	"github.com/DavidBalazic/SmartShopperApp/internal/rabbitmq"
)

type ProductService interface {
	GetProductById(ctx context.Context, id string) (models.Product, error)
	AddProduct(ctx context.Context, product models.Product) (models.Product, error)
	GetProductsByIds(ctx context.Context, ids []string) ([]models.Product, error)
	AddProducts(ctx context.Context, products []models.Product) ([]models.Product, error)
}

type productService struct {
	repo repo.ProductRepository
	publisher rabbitmq.Publisher
}

func NewProductService(repo repo.ProductRepository, publisher rabbitmq.Publisher) ProductService {
	return &productService{
		repo: repo,
		publisher: publisher,
	}
}

func (s *productService) GetProductById(ctx context.Context, id string) (models.Product, error) {
	product, err := s.repo.FindProductById(ctx, id)
	return product, err
}

func (s *productService) AddProduct(ctx context.Context, product models.Product) (models.Product, error) {
	if product.Name == "" {
		return models.Product{}, errors.New("product name is required")
	}
	if product.Price <= 0 {
		return models.Product{}, errors.New("product price must be greater than zero")
	}

	createdProduct, err := s.repo.AddProduct(ctx, product)
	if err != nil {
		log.Printf("failed to save product in repository: %v", err)
		return models.Product{}, err
	}

	message := map[string]interface{}{
		"id":            createdProduct.ID,
		"name":          createdProduct.Name,
		"description":   createdProduct.Description,
		"price":         createdProduct.Price,
		"quantity":      createdProduct.Quantity,
		"unit":          createdProduct.Unit,
		"store":         createdProduct.Store,
		"pricePerUnit":  createdProduct.PricePerUnit,
	}

	if err := s.publisher.PublishSingleProduct(message); err != nil {
		log.Printf("failed to publish product to RabbitMQ: %v", err)
		return models.Product{}, err
	}

	return createdProduct, nil
}

func (s *productService) GetProductsByIds(ctx context.Context, ids []string) ([]models.Product, error) {
	product, err := s.repo.FindProductsByIds(ctx, ids)
	return product, err
}

func (s *productService) AddProducts(ctx context.Context, products []models.Product) ([]models.Product, error) {
	for _, product := range products {
        if product.Name == "" {
            return nil, errors.New("each product must have a name")
        }
        if product.Price <= 0 {
            return nil, errors.New("each product must have a price greater than zero")
        }
    }

	createdProducts, err := s.repo.AddProducts(ctx, products)
	if err != nil {
		log.Printf("failed to save products in repository: %v", err)
		return nil, err
	}
	
	var productMessages []map[string]interface{}
	for _, createdProduct := range createdProducts {
		msg := map[string]interface{}{
			"id":            createdProduct.ID,
			"name":          createdProduct.Name,
			"description":   createdProduct.Description,
			"price":         createdProduct.Price,
			"quantity":      createdProduct.Quantity,
			"unit":          createdProduct.Unit,
			"store":         createdProduct.Store,
			"pricePerUnit":  createdProduct.PricePerUnit,
		}
		productMessages = append(productMessages, msg)
	}
	
	if err := s.publisher.PublishMultipleProducts(productMessages); err != nil {
		log.Printf("failed to publish products to RabbitMQ: %v", err)
		return nil, err
	}

	return createdProducts, nil
}