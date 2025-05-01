package tests

import (
	"context"
	"testing"

	"github.com/DavidBalazic/SmartShopperApp/internal/models"
	"github.com/DavidBalazic/SmartShopperApp/internal/services"
	"github.com/DavidBalazic/SmartShopperApp/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	gomock "go.uber.org/mock/gomock"
)

func TestProductService_FindProductByID_OneProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepository(ctrl)
	mockRabbitmq := mocks.NewMockPublisher(ctrl)
	productService := services.NewProductService(mockRepo, mockRabbitmq)

	ctx := context.Background()

	t.Run("successfully find product by ID", func(t *testing.T) {
		objectID := primitive.NewObjectID()
		expectedProduct := models.Product{
			ID:    objectID.Hex(),
			Name:  "Milk",
			Price: 1.49,
			Store: "Aldi",
		}

		mockRepo.EXPECT().
			FindProductById(gomock.Any(), objectID.Hex()).
			Return(expectedProduct, nil)

		result, err := productService.GetProductById(ctx, objectID.Hex())

		assert.NoError(t, err)
		assert.Equal(t, expectedProduct.ID, result.ID)
		assert.Equal(t, expectedProduct.Name, result.Name)
		assert.Equal(t, expectedProduct.Price, result.Price)
		assert.Equal(t, expectedProduct.Store, result.Store)
	})
}

func TestProductService_FindProductsByIDs_OneProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepository(ctrl)
	mockRabbitmq := mocks.NewMockPublisher(ctrl)
	productService := services.NewProductService(mockRepo, mockRabbitmq)

	ctx := context.Background()

	t.Run("successfully find product by ID", func(t *testing.T) {
		objectID := primitive.NewObjectID()
		expectedProduct := models.Product{
			ID:    objectID.Hex(),
			Name:  "Milk",
			Price: 1.49,
			Store: "Aldi",
		}

		mockRepo.EXPECT().
			FindProductsByIds(gomock.Any(), []string{objectID.Hex()}).
			Return([]models.Product{expectedProduct}, nil)

		result, err := productService.GetProductsByIds(ctx, []string{objectID.Hex()})

		assert.NoError(t, err)
		assert.Len(t, result, 1)
		assert.Equal(t, expectedProduct.ID, result[0].ID)
		assert.Equal(t, expectedProduct.Name, result[0].Name)
		assert.Equal(t, expectedProduct.Price, result[0].Price)
		assert.Equal(t, expectedProduct.Store, result[0].Store)
	})
}

func TestProductService_FindProductsByIDs_MultipleProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepository(ctrl)
	mockRabbitmq := mocks.NewMockPublisher(ctrl)
	productService := services.NewProductService(mockRepo, mockRabbitmq)

	ctx := context.Background()

	t.Run("successfully find multiple products by IDs", func(t *testing.T) {
		id1 := primitive.NewObjectID().Hex()
		id2 := primitive.NewObjectID().Hex()

		expectedProducts := []models.Product{
			{
				ID:    id1,
				Name:  "Milk",
				Price: 1.49,
				Store: "Aldi",
			},
			{
				ID:    id2,
				Name:  "Bread",
				Price: 0.99,
				Store: "Lidl",
			},
		}

		mockRepo.EXPECT().
			FindProductsByIds(gomock.Any(), []string{id1, id2}).
			Return(expectedProducts, nil)

		result, err := productService.GetProductsByIds(ctx, []string{id1, id2})

		assert.NoError(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, expectedProducts[0].ID, result[0].ID)
		assert.Equal(t, expectedProducts[1].ID, result[1].ID)
		assert.Equal(t, expectedProducts[0].Name, result[0].Name)
		assert.Equal(t, expectedProducts[1].Name, result[1].Name)
	})
}

func TestProductService_FindProductsByIDs_SomeInvalidIds(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepository(ctrl)
	mockRabbitmq := mocks.NewMockPublisher(ctrl)
	productService := services.NewProductService(mockRepo, mockRabbitmq)

	ctx := context.Background()

	t.Run("returns only found products when some IDs are invalid", func(t *testing.T) {
		validID1 := primitive.NewObjectID().Hex()
		validID2 := primitive.NewObjectID().Hex()
		invalidID := "invalidid123"

		expectedProducts := []models.Product{
			{
				ID:    validID1,
				Name:  "Milk",
				Price: 1.49,
				Store: "Aldi",
			},
			{
				ID:    validID2,
				Name:  "Eggs",
				Price: 2.49,
				Store: "Spar",
			},
		}

		mockRepo.EXPECT().
			FindProductsByIds(gomock.Any(), []string{validID1, validID2, invalidID}).
			Return(expectedProducts, nil)

		result, err := productService.GetProductsByIds(ctx, []string{validID1, validID2, invalidID})

		assert.NoError(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, expectedProducts[0], result[0])
		assert.Equal(t, expectedProducts[1], result[1])
	})
}

func TestProductService_AddProduct_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepository(ctrl)
	mockRabbitmq := mocks.NewMockPublisher(ctrl)
	productService := services.NewProductService(mockRepo, mockRabbitmq)

	ctx := context.Background()

	product := models.Product{
		Name:         "Milk",
		Description:  "Fresh dairy product",
		Price:        1.49,
		Quantity:     100,
		Unit:         "Litre",
		Store:        "Aldi",
		PricePerUnit: 1.49,
	}

	createdProduct := product
	createdProduct.ID = primitive.NewObjectID().Hex()

	mockRepo.EXPECT().
		AddProduct(gomock.Any(), product).
		Return(createdProduct, nil)

	mockRabbitmq.EXPECT().
		PublishSingleProduct(gomock.Any()).Times(1)

	result, err := productService.AddProduct(ctx, product)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, createdProduct.ID, result.ID)
	assert.Equal(t, product.Name, result.Name)
	assert.Equal(t, product.Price, result.Price)
}

func TestProductService_AddProducts_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepository(ctrl)
	mockRabbitmq := mocks.NewMockPublisher(ctrl)
	productService := services.NewProductService(mockRepo, mockRabbitmq)

	ctx := context.Background()

	products := []models.Product{
		{
			Name:         "Milk",
			Description:  "Fresh dairy product",
			Price:        1.49,
			Quantity:     100,
			Unit:         "Litre",
			Store:        "Aldi",
			PricePerUnit: 1.49,
		},
		{
			Name:         "Bread",
			Description:  "Whole wheat bread",
			Price:        2.99,
			Quantity:     50,
			Unit:         "Loaf",
			Store:        "Spar",
			PricePerUnit: 2.99,
		},
		{
			Name:         "Eggs",
			Description:  "Free-range eggs",
			Price:        3.99,
			Quantity:     30,
			Unit:         "Carton",
			Store:        "Tesco",
			PricePerUnit: 3.99,
		},
	}

	createdProducts := []models.Product{
		{
			ID:           primitive.NewObjectID().Hex(),
			Name:         products[0].Name,
			Description:  products[0].Description,
			Price:        products[0].Price,
			Quantity:     products[0].Quantity,
			Unit:         products[0].Unit,
			Store:        products[0].Store,
			PricePerUnit: products[0].PricePerUnit,
		},
		{
			ID:           primitive.NewObjectID().Hex(),
			Name:         products[1].Name,
			Description:  products[1].Description,
			Price:        products[1].Price,
			Quantity:     products[1].Quantity,
			Unit:         products[1].Unit,
			Store:        products[1].Store,
			PricePerUnit: products[1].PricePerUnit,
		},
		{
			ID:           primitive.NewObjectID().Hex(),
			Name:         products[2].Name,
			Description:  products[2].Description,
			Price:        products[2].Price,
			Quantity:     products[2].Quantity,
			Unit:         products[2].Unit,
			Store:        products[2].Store,
			PricePerUnit: products[2].PricePerUnit,
		},
	}

	mockRepo.EXPECT().
		AddProducts(gomock.Any(), products).
		Return(createdProducts, nil).
		Times(1)

	mockRabbitmq.EXPECT().
		PublishMultipleProducts(gomock.Any()).
		Return(nil).
		Times(1) 

	result, err := productService.AddProducts(ctx, products)

	assert.NoError(t, err)
	assert.Len(t, result, 3) 
	assert.Equal(t, createdProducts[0].ID, result[0].ID)
	assert.Equal(t, createdProducts[1].ID, result[1].ID)
	assert.Equal(t, createdProducts[2].ID, result[2].ID)
}