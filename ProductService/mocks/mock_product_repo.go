// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repo/product_repo.go
//
// Generated by this command:
//
//	mockgen -source=internal/repo/product_repo.go -destination=mocks/mock_product_repo.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	models "github.com/DavidBalazic/SmartShopperApp/internal/models"
	gomock "go.uber.org/mock/gomock"
)

// MockProductRepository is a mock of ProductRepository interface.
type MockProductRepository struct {
	ctrl     *gomock.Controller
	recorder *MockProductRepositoryMockRecorder
}

// MockProductRepositoryMockRecorder is the mock recorder for MockProductRepository.
type MockProductRepositoryMockRecorder struct {
	mock *MockProductRepository
}

// NewMockProductRepository creates a new mock instance.
func NewMockProductRepository(ctrl *gomock.Controller) *MockProductRepository {
	mock := &MockProductRepository{ctrl: ctrl}
	mock.recorder = &MockProductRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductRepository) EXPECT() *MockProductRepositoryMockRecorder {
	return m.recorder
}

// AddProduct mocks base method.
func (m *MockProductRepository) AddProduct(ctx context.Context, product models.Product) (models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProduct", ctx, product)
	ret0, _ := ret[0].(models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddProduct indicates an expected call of AddProduct.
func (mr *MockProductRepositoryMockRecorder) AddProduct(ctx, product any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProduct", reflect.TypeOf((*MockProductRepository)(nil).AddProduct), ctx, product)
}

// AddProducts mocks base method.
func (m *MockProductRepository) AddProducts(ctx context.Context, products []models.Product) ([]models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProducts", ctx, products)
	ret0, _ := ret[0].([]models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddProducts indicates an expected call of AddProducts.
func (mr *MockProductRepositoryMockRecorder) AddProducts(ctx, products any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProducts", reflect.TypeOf((*MockProductRepository)(nil).AddProducts), ctx, products)
}

// FindProductById mocks base method.
func (m *MockProductRepository) FindProductById(ctx context.Context, id string) (models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindProductById", ctx, id)
	ret0, _ := ret[0].(models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindProductById indicates an expected call of FindProductById.
func (mr *MockProductRepositoryMockRecorder) FindProductById(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindProductById", reflect.TypeOf((*MockProductRepository)(nil).FindProductById), ctx, id)
}

// FindProductsByIds mocks base method.
func (m *MockProductRepository) FindProductsByIds(ctx context.Context, ids []string) ([]models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindProductsByIds", ctx, ids)
	ret0, _ := ret[0].([]models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindProductsByIds indicates an expected call of FindProductsByIds.
func (mr *MockProductRepositoryMockRecorder) FindProductsByIds(ctx, ids any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindProductsByIds", reflect.TypeOf((*MockProductRepository)(nil).FindProductsByIds), ctx, ids)
}
