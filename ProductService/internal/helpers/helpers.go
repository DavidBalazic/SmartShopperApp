package helpers

import (
	"context"
	"strings"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/metadata"

	"github.com/DavidBalazic/SmartShopperApp/internal/proto"
	"github.com/DavidBalazic/SmartShopperApp/internal/models"
)


func ToProductResponse(p models.Product) *proto.ProductResponse {
	return &proto.ProductResponse{
		Product: ToProtoProduct(p),
	}
}

func ToProtoProducts(products []models.Product) []*proto.Product {
	result := make([]*proto.Product, 0, len(products))
	for _, p := range products {
		result = append(result, ToProtoProduct(p))
	}
	return result
}

func ToProtoProduct(p models.Product) *proto.Product {
	return &proto.Product{
		Id:           p.ID,
		Name:         p.Name,
		Description:  p.Description,
		Price:        p.Price,
		Quantity:     p.Quantity,
		Unit:         p.Unit,
		Store:        p.Store,
		PricePerUnit: p.PricePerUnit,
		ImageUrl:     p.ImageUrl,
	}
}

func ExtractClientInfo(ctx context.Context) (ip, userAgent string) {
	// Get IP from peer info
	if p, ok := peer.FromContext(ctx); ok && p.Addr != nil {
		ip = p.Addr.String()
		// Remove port if present
		if colon := strings.LastIndex(ip, ":"); colon != -1 {
			ip = ip[:colon]
		}
	}

	// Get metadata (headers)
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		if ua := md.Get("user-agent"); len(ua) > 0 {
			userAgent = ua[0]
		}
	}

	return ip, userAgent
}