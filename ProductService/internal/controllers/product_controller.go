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
	service  services.ProductService
}

func NewProductController(service services.ProductService) *ProductController {
	return &ProductController{
		service:  service,
	}
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

	return toProductResponse(createdProduct), nil
}

func (s *ProductController) GetProductsByIds(ctx context.Context, req *proto.ProductsIdsRequest) (*proto.ProductList, error) {
	if len(req.GetIds()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "list of IDs is required")
	}
	products, err := s.service.GetProductsByIds(ctx, req.GetIds())
	if err != nil {
		return nil, err
	}
	return &proto.ProductList{Products: toProtoProducts(products)}, nil
}

func (s *ProductController) AddProducts(ctx context.Context, req *proto.AddProductsRequest) (*proto.ProductList, error) {
	if len(req.GetProducts()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "at least one product is required")
	}

	var productModels []models.Product

	for _, p := range req.GetProducts() {
		if p.GetName() == "" || p.GetPrice() == 0 || p.GetQuantity() == 0 || p.GetUnit() == "" || p.GetStore() == "" || p.GetPricePerUnit() == 0 {
			return nil, status.Error(codes.InvalidArgument, "all product fields are required")
		}

		product := models.Product{
			Name:         p.GetName(),
			Description:  p.GetDescription(),
			Price:        p.GetPrice(),
			Quantity:     p.GetQuantity(),
			Unit:         p.GetUnit(),
			Store:        p.GetStore(),
			PricePerUnit: p.GetPricePerUnit(),
		}

		productModels = append(productModels, product)
	}

	createdProducts, err := s.service.AddProducts(ctx, productModels)
	if err != nil {
		return nil, err
	}

	return &proto.ProductList{Products: toProtoProducts(createdProducts)}, nil
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