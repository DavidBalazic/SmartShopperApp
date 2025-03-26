// package tests

// import (
// 	"context"
// 	"errors"
// 	"testing"

// 	"github.com/DavidBalazic/SmartShopperApp/internal/repo"
// 	"github.com/DavidBalazic/SmartShopperApp/internal/models"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// type MockCollection struct {
// 	mock.Mock
// }

// func (m *MockCollection) FindOne(ctx context.Context, filter interface{}) *mongo.SingleResult {
// 	args := m.Called(ctx, filter)
// 	result := args.Get(0)
// 	if result == nil {
// 		return &mongo.SingleResult{}
// 	}
// 	return result.(*mongo.SingleResult)
// }

// func (m *MockCollection) Find(ctx context.Context, filter interface{}) (*mongo.Cursor, error) {
// 	args := m.Called(ctx, filter)
// 	return args.Get(0).(*mongo.Cursor), args.Error(1)
// }

// func newTestRepo(mockCollection *MockCollection) *repo.MongoProductRepository {
// 	return &repo.MongoProductRepository{Db: mockCollection}
// }

// // Test FindCheapestProduct
// func TestFindCheapestProduct(t *testing.T) {
// 	mockCollection := new(MockCollection)
// 	repo := newTestRepo(mockCollection)

// 	expectedFilter := bson.M{"$or": []bson.M{
// 		{"name": bson.M{"$regex": "apple", "$options": "i"}},
// 		{"name": bson.M{"$regex": "juice", "$options": "i"}},
// 	}}

// 	mockProduct := models.Product{
// 		ID:           "123",
// 		Name:         "Apple Juice",
// 		Price:        2.97,
// 		Quantity:     3.0,
// 		Unit:         "l",
// 		Store:        "lidl",
// 		PricePerUnit: 0.99,
// 	}

// 	// Simulate FindOne returning a valid result
// 	mockCollection.On("FindOne", mock.Anything, expectedFilter).
// 		Return(mongo.NewSingleResultFromDocument(mockProduct, nil, nil))

// 	result, err := repo.FindCheapestProduct("Apple Juice")

// 	// Assertions
// 	assert.NoError(t, err)
// 	assert.Equal(t, mockProduct, result)
// 	mockCollection.AssertExpectations(t)
// }

// // Test FindCheapestProductByStore
// func TestFindCheapestProductByStore(t *testing.T) {
// 	mockCollection := new(MockCollection)
// 	repo := newTestRepo(mockCollection)

// 	// Expected query conditions
// 	expectedFilter := bson.M{"$or": []bson.M{
// 		{"name": bson.M{"$regex": "milk", "$options": "i"}},
// 	}, "store": bson.M{"$regex": "walmart", "$options": "i"}}

// 	// Mock response
// 	mockProduct := models.Product{
// 		ID:           "456",
// 		Name:         "Milk",
// 		Quantity:     3.0,
// 		Unit:         "l",
// 		Price:        1.50,
// 		Store:        "spar",
// 		PricePerUnit: 0.50,
// 	}

// 	// Simulate FindOne returning a valid result
// 	mockCollection.On("FindOne", mock.Anything, expectedFilter).
// 		Return(mongo.NewSingleResultFromDocument(mockProduct, nil, nil))

// 	// Call the function
// 	result, err := repo.FindCheapestProductByStore("Milk", "Walmart")

// 	// Assertions
// 	assert.NoError(t, err)
// 	assert.Equal(t, mockProduct, result)
// 	mockCollection.AssertExpectations(t)
// }

// // Test FindAllProductPrices
// func TestFindAllProductPrices(t *testing.T) {
// 	mockCollection := new(MockCollection)
// 	repo := newTestRepo(mockCollection)

// 	// Expected query conditions
// 	expectedFilter := bson.M{"$or": []bson.M{
// 		{"name": bson.M{"$regex": "cereal", "$options": "i"}},
// 	}}

// 	// Mock response
// 	mockProducts := []models.Product{
// 		{ID: "789", Name: "Cereal A", Price: 3.99, Quantity: 500, Unit: "g", Store: "lidl", PricePerUnit: 2.00},
// 		{ID: "790", Name: "Cereal B", Price: 4.50, Quantity: 500, Unit: "g", Store: "spar", PricePerUnit: 2.25},
// 	}

// 	var mockProductsInterface []interface{}
// 	for _, p := range mockProducts {
// 		mockProductsInterface = append(mockProductsInterface, p)
// 	}

// 	// Mocking Find method
// 	cursor, err := mongo.NewCursorFromDocuments(mockProductsInterface, nil, nil)
// 	mockCollection.On("Find", mock.Anything, expectedFilter).
// 		Return(cursor, nil)

// 	// Call the function
// 	result, err := repo.FindAllProductPrices("Cereal")

// 	// Assertions
// 	assert.NoError(t, err)
// 	assert.Len(t, result, 2)
// 	assert.Equal(t, mockProducts, result)
// 	mockCollection.AssertExpectations(t)
// }

// // Test case when no products are found
// func TestFindCheapestProduct_NotFound(t *testing.T) {
// 	mockCollection := new(MockCollection)
// 	repo := newTestRepo(mockCollection)

// 	expectedFilter := bson.M{"$or": []bson.M{
// 		{"name": bson.M{"$regex": "unknown", "$options": "i"}},
// 	}}

// 	// Simulate FindOne returning no results
// 	mockCollection.On("FindOne", mock.Anything, expectedFilter).
// 		Return(&mongo.SingleResult{})

// 	// Call the function
// 	result, err := repo.FindCheapestProduct("Unknown Product")

// 	// Assertions
// 	assert.Error(t, err)
// 	assert.Equal(t, models.Product{}, result)
// 	mockCollection.AssertExpectations(t)
// }

// // Test case when an error occurs
// func TestFindCheapestProduct_Error(t *testing.T) {
// 	mockCollection := new(MockCollection)
// 	repo := newTestRepo(mockCollection)

// 	expectedFilter := bson.M{"$or": []bson.M{
// 		{"name": bson.M{"$regex": "errorcase", "$options": "i"}},
// 	}}

// 	// Simulate MongoDB returning an error
// 	mockCollection.On("FindOne", mock.Anything, expectedFilter).
// 		Return(nil, errors.New("database error"))

// 	// Call the function
// 	result, err := repo.FindCheapestProduct("ErrorCase")

// 	// Assertions
// 	assert.Error(t, err)
// 	assert.Equal(t, models.Product{}, result)
// 	mockCollection.AssertExpectations(t)
// }

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