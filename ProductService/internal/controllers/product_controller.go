package controllers

import (
	"context"

	"github.com/DavidBalazic/SmartShopperApp/internal/models"
	"github.com/DavidBalazic/SmartShopperApp/internal/proto"
	"github.com/DavidBalazic/SmartShopperApp/internal/services"
	"github.com/DavidBalazic/SmartShopperApp/internal/rabbitmq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


type ProductController struct {
	proto.UnimplementedProductServiceServer
	service  services.ProductService
	publisher *rabbitmq.Publisher
}

func NewProductController(service services.ProductService, publisher *rabbitmq.Publisher) *ProductController {
	return &ProductController{
		service:  service,
		publisher: publisher,
	}
}

func (s *ProductController) GetCheapestProduct(ctx context.Context, req *proto.ProductRequest) (*proto.ProductResponse, error) {
	if req.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "product name is required")
	}

	product, err := s.service.GetCheapestProduct(ctx, req.GetName()) // Added context propagation
	if err != nil {
		return nil, err
	}

	return toProductResponse(product), nil
}

func (s *ProductController) GetCheapestByStore(ctx context.Context, req *proto.StoreRequest) (*proto.ProductResponse, error) {
	if req.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "product name is required")
	}
	if req.GetStore() == "" {
		return nil, status.Error(codes.InvalidArgument, "store name is required")
	}

	product, err := s.service.GetCheapestByStore(ctx, req.GetName(), req.GetStore())
	if err != nil {
		return nil, err
	}

	return toProductResponse(product), nil
}

func (s *ProductController) GetAllPrices(ctx context.Context, req *proto.ProductRequest) (*proto.ProductList, error) {
	if req.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "product name is required")
	}

	products, err := s.service.GetAllPrices(ctx, req.GetName())
	if err != nil {
		return nil, err
	}

	return &proto.ProductList{
		Products: toProtoProducts(products),
	}, nil
}

func (s *ProductController) GetProductById(ctx context.Context, req *proto.ProductIdRequest) (*proto.ProductResponse, error) {
	if req.GetId() == "" {
		return nil, status.Error(codes.InvalidArgument, "product ID is required")
	}

	product, err := s.service.GetProductById(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return toProductResponse(product), nil
}

func (s *ProductController) AddProduct(ctx context.Context, req *proto.AddProductRequest) (*proto.ProductResponse, error) {
	if req.GetName() == "" || req.GetPrice() == 0 || req.GetQuantity() == 0 || req.GetUnit() == "" || req.GetStore() == ""  || req.GetPricePerUnit() == 0 { 
		return nil, status.Error(codes.InvalidArgument, "product name, price, quantity, unit, store and price per unit are required")
	}

	product := models.Product{
		Name:         req.GetName(),
		Description:  req.GetDescription(),
		Price:        req.GetPrice(),
		Quantity:     req.GetQuantity(),
		Unit:         req.GetUnit(),
		Store:        req.GetStore(),
		PricePerUnit: req.GetPricePerUnit(),
	}

	createdProduct, err := s.service.AddProduct(ctx, product)
	if err != nil {
		return nil, err
	}

	productMessage := map[string]interface{}{
		"id":            createdProduct.ID,
		"name":          createdProduct.Name,
		"description":   createdProduct.Description,
		"price":         createdProduct.Price,
		"quantity":      createdProduct.Quantity,
		"unit":          createdProduct.Unit,
		"store":         createdProduct.Store,
		"pricePerUnit":  createdProduct.PricePerUnit,
	}
	err = s.publisher.Publish(productMessage)
	if err != nil {
		return nil, err
	}

	return toProductResponse(product), nil
}

// Helper functions for conversion
func toProductResponse(p models.Product) *proto.ProductResponse {
	return &proto.ProductResponse{
		Product: toProtoProduct(p),
	}
}

func toProtoProducts(products []models.Product) []*proto.Product {
	result := make([]*proto.Product, 0, len(products))
	for _, p := range products {
		result = append(result, toProtoProduct(p))
	}
	return result
}

func toProtoProduct(p models.Product) *proto.Product {
	return &proto.Product{
		Id:           p.ID,
		Name:         p.Name,
		Description:  p.Description,
		Price:        p.Price,
		Quantity:     p.Quantity,
		Unit:         p.Unit,
		Store:        p.Store,
		PricePerUnit: p.PricePerUnit,
	}
}