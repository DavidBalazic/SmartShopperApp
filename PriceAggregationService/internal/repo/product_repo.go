package repo

import (
	"context"
	"github.com/DavidBalazic/SmartShopperApp/internal/models"
)

type ProductRepository interface {
	FindCheapestProduct(ctx context.Context, name string) (models.Product, error)
	FindCheapestProductByStore(ctx context.Context, name, store string) (models.Product, error)
	FindAllProductPrices(ctx context.Context, name string) ([]models.Product, error)
}