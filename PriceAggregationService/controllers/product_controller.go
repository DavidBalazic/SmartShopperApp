package controllers

import (
	"context"
	"github.com/DavidBalazic/SmartShopperApp/proto"
	"github.com/DavidBalazic/SmartShopperApp/services"
)

type ProductController struct {
	proto.UnimplementedProductServiceServer
	service *services.ProductService
}

func NewProductController(service *services.ProductService) *ProductController {
	return &ProductController{service: service}
}

func (c *ProductController) GetCheapestProduct(ctx context.Context, req *proto.ProductRequest) (*proto.ProductResponse, error) {
	product, err := c.service.GetCheapestProduct(req.Name)
	if err != nil {
		return nil, err
	}
	return &proto.ProductResponse{Product: &proto.Product{
		Id:           product.ID,
		Name:         product.Name,
		Description:  product.Description,
		Price:        product.Price,
		Quantity:     product.Quantity,
		Unit:     	  product.Unit,
		Store:        product.Store,
		PricePerUnit: product.PricePerUnit,
	}}, nil
}

func (c *ProductController) GetCheapestByStore(ctx context.Context, req *proto.StoreRequest) (*proto.ProductResponse, error) {
	product, err := c.service.GetCheapestByStore(req.Name, req.Store)
	if err != nil {
		return nil, err
	}

	return &proto.ProductResponse{
		Product: &proto.Product{
			Id:           product.ID,
			Name:         product.Name,
			Description:  product.Description,
			Price:        product.Price,
			Quantity:     product.Quantity,
			Unit:         product.Unit,
			Store:        product.Store,
			PricePerUnit: product.PricePerUnit,
		},
	}, nil
}

func (c *ProductController) GetAllPrices(ctx context.Context, req *proto.ProductRequest) (*proto.ProductList, error) {
	products, err := c.service.GetAllPrices(req.Name)
	if err != nil {
		return nil, err
	}

	var productList []*proto.Product
	for _, product := range products {
		productList = append(productList, &proto.Product{
			Id:           product.ID,
			Name:         product.Name,
			Description:  product.Description,
			Price:        product.Price,
			Quantity:     product.Quantity,
			Unit:         product.Unit,
			Store:        product.Store,
			PricePerUnit: product.PricePerUnit,
		})
	}

	return &proto.ProductList{Products: productList}, nil
}