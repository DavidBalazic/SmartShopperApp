package repo

import (
	"github.com/DavidBalazic/SmartShopperApp/models"
)

type ProductRepository interface {
	FindCheapestProduct(name string) (models.Product, error)
	FindCheapestProductByStore(name, store string) (models.Product, error)
	FindAllProductPrices(name string) ([]models.Product, error)
}