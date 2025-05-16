package controllers

import (
	"context"


	"github.com/DavidBalazic/SmartShopperApp/internal/models"
	"github.com/DavidBalazic/SmartShopperApp/internal/proto"
	"github.com/DavidBalazic/SmartShopperApp/internal/services"
	"github.com/DavidBalazic/SmartShopperApp/internal/contextkeys"
	"github.com/DavidBalazic/SmartShopperApp/internal/helpers"

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

	return helpers.ToProductResponse(product), nil
}

func (s *ProductController) AddProduct(ctx context.Context, req *proto.AddProductRequest) (*proto.ProductResponse, error) {
	// if req.GetName() == "" || req.GetPrice() == 0 || req.GetQuantity() == 0 || req.GetUnit() == "" || req.GetStore() == ""  || req.GetPricePerUnit() == 0 { 
	// 	return nil, status.Error(codes.InvalidArgument, "product name, price, quantity, unit, store and price per unit are required")
	// }

	ip, userAgent := helpers.ExtractClientInfo(ctx)

	// Put metadata in context for downstream service
	ctx = context.WithValue(ctx, contextkeys.IPKey, ip)
	ctx = context.WithValue(ctx, contextkeys.UserAgentKey, userAgent)

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

	return helpers.ToProductResponse(createdProduct), nil
}

func (s *ProductController) GetProductsByIds(ctx context.Context, req *proto.ProductsIdsRequest) (*proto.ProductList, error) {
	if len(req.GetIds()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "list of IDs is required")
	}
	products, err := s.service.GetProductsByIds(ctx, req.GetIds())
	if err != nil {
		return nil, err
	}
	return &proto.ProductList{Products: helpers.ToProtoProducts(products)}, nil
}

func (s *ProductController) AddProducts(ctx context.Context, req *proto.AddProductsRequest) (*proto.ProductList, error) {
	if len(req.GetProducts()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "at least one product is required")
	}

	ip, userAgent := helpers.ExtractClientInfo(ctx)

	// Put metadata in context for downstream service
	ctx = context.WithValue(ctx, contextkeys.IPKey, ip)
	ctx = context.WithValue(ctx, contextkeys.UserAgentKey, userAgent)

	var productModels []models.Product

	for _, p := range req.GetProducts() {
		// if p.GetName() == "" || p.GetPrice() == 0 || p.GetQuantity() == 0 || p.GetUnit() == "" || p.GetStore() == "" || p.GetPricePerUnit() == 0 {
		// 	return nil, status.Error(codes.InvalidArgument, "all product fields are required")
		// }

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

	return &proto.ProductList{Products: helpers.ToProtoProducts(createdProducts)}, nil
}