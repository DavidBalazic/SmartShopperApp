package services

import (
	"context"
	"errors"
	"log"
	"time"
	"fmt"

	"github.com/DavidBalazic/SmartShopperApp/internal/models"
	"github.com/DavidBalazic/SmartShopperApp/internal/repo"
	"github.com/DavidBalazic/SmartShopperApp/internal/rabbitmq"
	"github.com/DavidBalazic/SmartShopperApp/internal/kafka"
	"github.com/DavidBalazic/SmartShopperApp/internal/dtos"
	"github.com/DavidBalazic/SmartShopperApp/internal/contextkeys"
)

type ProductService interface {
	GetProductById(ctx context.Context, id string) (models.Product, error)
	AddProduct(ctx context.Context, product models.Product) (models.Product, error)
	GetProductsByIds(ctx context.Context, ids []string) ([]models.Product, error)
	AddProducts(ctx context.Context, products []models.Product) ([]models.Product, error)
}

type productService struct {
	repo      repo.ProductRepository
	publisher rabbitmq.Publisher
	auditLogger kafka.AuditLogger
}

func NewProductService(repo repo.ProductRepository, publisher rabbitmq.Publisher, auditLogger kafka.AuditLogger) ProductService {
	return &productService{
		repo:      repo,
		publisher: publisher,
		auditLogger: auditLogger,
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

	ip := ctx.Value(contextkeys.IPKey)
	userAgent := ctx.Value(contextkeys.UserAgentKey)

	createdProduct, err := s.repo.AddProduct(ctx, product)
	if err != nil {
		log.Printf("failed to save product in repository: %v", err)
		return models.Product{}, err
	}

	auditLog := dtos.AuditLog{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Actor: dtos.AuditActor{
			ID:        "admin", // TODO: retrieve from context
			IP:        fmt.Sprintf("%v", ip),
			UserAgent: fmt.Sprintf("%v", userAgent),
		},
		Action:   "add-product",
		Resource: "product",
		Service:  "ProductService",
		Details: map[string]interface{}{
			"name":         createdProduct.Name,
			"price":        createdProduct.Price,
			"quantity":     createdProduct.Quantity,
			"unit":         createdProduct.Unit,
			"store":        createdProduct.Store,
			"pricePerUnit": createdProduct.PricePerUnit,
			"imageUrl":     createdProduct.ImageUrl,
		},
	}

	if err := s.auditLogger.PublishAuditLog(ctx, auditLog); err != nil {
		log.Printf("failed to publish audit log: %v", err)
	}

	message := dtos.ProductMessage{
		ID:           createdProduct.ID,
		Name:         createdProduct.Name,
		Description:  createdProduct.Description,
		Price:        createdProduct.Price,
		Quantity:     createdProduct.Quantity,
		Unit:         createdProduct.Unit,
		Store:        createdProduct.Store,
		PricePerUnit: createdProduct.PricePerUnit,
		ImageUrl:     createdProduct.ImageUrl,
	}

	// Publish the product message to RabbitMQ
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

	ip := ctx.Value(contextkeys.IPKey)
	userAgent := ctx.Value(contextkeys.UserAgentKey)

	createdProducts, err := s.repo.AddProducts(ctx, products)
	if err != nil {
		log.Printf("failed to save products in repository: %v", err)
		return nil, err
	}
	
	var productMessages []dtos.ProductMessage
	for _, createdProduct := range createdProducts {
		msg := dtos.ProductMessage{
			ID:           createdProduct.ID,
			Name:         createdProduct.Name,
			Description:  createdProduct.Description,
			Price:        createdProduct.Price,
			Quantity:     createdProduct.Quantity,
			Unit:         createdProduct.Unit,
			Store:        createdProduct.Store,
			PricePerUnit: createdProduct.PricePerUnit,
			ImageUrl:     createdProduct.ImageUrl,
		}
		productMessages = append(productMessages, msg)
	}

	auditLog := dtos.AuditLog{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Actor: dtos.AuditActor{
			ID:        "admin", // TODO: retrieve from context
			IP:        fmt.Sprintf("%v", ip),
			UserAgent: fmt.Sprintf("%v", userAgent),
		},
		Action:   "add-products",
		Resource: "products",
		Service:  "ProductService",
		Details: map[string]interface{}{
			"products": productMessages,
		},
	}

	if err := s.auditLogger.PublishAuditLog(ctx, auditLog); err != nil {
		log.Printf("failed to publish audit log: %v", err)
	}

	// Publish each product message to RabbitMQ
	if err := s.publisher.PublishMultipleProducts(productMessages); err != nil {
		log.Printf("failed to publish products to RabbitMQ: %v", err)
		return nil, err
	}

	return createdProducts, nil
}