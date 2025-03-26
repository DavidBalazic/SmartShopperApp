package tests

import (
	"context"
	"testing"

	"github.com/DavidBalazic/SmartShopperApp/internal/models"
	"github.com/DavidBalazic/SmartShopperApp/internal/services"
	"github.com/DavidBalazic/SmartShopperApp/mocks"
	gomock "go.uber.org/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestProductService_GetCheapestProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepository(ctrl)
	productService := services.NewProductService(mockRepo)

	ctx := context.Background()

	t.Run("successfully find cheapest product", func(t *testing.T) {
		expectedProduct := models.Product{
			Name:         "Apple Juice",
			Price:        2.97,
			Quantity:     3.0,
			Unit:         "l",
			Store:        "lidl",
			PricePerUnit: 0.99,
		}

		mockRepo.EXPECT().
			FindCheapestProduct(gomock.Any(), "juice").
			Return(expectedProduct, nil)

		result, err := productService.GetCheapestProduct(ctx, "juice")

		assert.NoError(t, err)
		assert.Equal(t, "Apple Juice", result.Name)
		assert.Equal(t, 2.97, result.Price)
		assert.Equal(t, "lidl", result.Store)
	})
}

func TestProductService_GetCheapestByStore(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepository(ctrl)
	productService := services.NewProductService(mockRepo)

	ctx := context.Background()

	t.Run("successfully find cheapest product by store", func(t *testing.T) {
		expectedProduct := models.Product{
			Name:  "Organic Eggs",
			Price: 3.99,
			Store: "Costco",
		}

		mockRepo.EXPECT().
			FindCheapestProductByStore(gomock.Any(), "eggs", "costco").
			Return(expectedProduct, nil)

		result, err := productService.GetCheapestByStore(ctx, "eggs", "costco")

		assert.NoError(t, err)
		assert.Equal(t, "Organic Eggs", result.Name)
		assert.Equal(t, 3.99, result.Price)
		assert.Equal(t, "Costco", result.Store)
	})
}

func TestProductService_GetAllPrices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepository(ctrl)
	productService := services.NewProductService(mockRepo)

	ctx := context.Background()

	t.Run("successfully get all prices for product", func(t *testing.T) {
		expectedProducts := []models.Product{
			{
				Name:  "Bread",
				Price: 2.99,
				Store: "Walmart",
			},
			{
				Name:  "Bread",
				Price: 3.49,
				Store: "Costco",
			},
		}

		mockRepo.EXPECT().
			FindAllProductPrices(gomock.Any(), "bread").
			Return(expectedProducts, nil)

		results, err := productService.GetAllPrices(ctx, "bread")

		assert.NoError(t, err)
		assert.Len(t, results, 2)
		assert.Equal(t, "Walmart", results[0].Store)
		assert.Equal(t, 2.99, results[0].Price)
		assert.Equal(t, "Costco", results[1].Store)
		assert.Equal(t, 3.49, results[1].Price)
	})
}