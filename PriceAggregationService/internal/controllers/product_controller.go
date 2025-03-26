package controllers

import (
	"context"

	"github.com/DavidBalazic/SmartShopperApp/internal/models"
	"github.com/DavidBalazic/SmartShopperApp/internal/proto"
	"github.com/DavidBalazic/SmartShopperApp/internal/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductController struct {
	proto.UnimplementedProductServiceServer
	service services.ProductService
}

// NewProductServer creates a new gRPC server instance
func NewProductController(service services.ProductService) *ProductController {
	return &ProductController{service: service}
}

// GetCheapestProduct returns the cheapest product matching the name
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

// GetCheapestByStore returns the cheapest product in a specific store
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

// GetAllPrices returns all products matching the name
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