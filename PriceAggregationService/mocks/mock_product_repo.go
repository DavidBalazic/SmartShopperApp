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

// FindAllProductPrices mocks base method.
func (m *MockProductRepository) FindAllProductPrices(ctx context.Context, name string) ([]models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllProductPrices", ctx, name)
	ret0, _ := ret[0].([]models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllProductPrices indicates an expected call of FindAllProductPrices.
func (mr *MockProductRepositoryMockRecorder) FindAllProductPrices(ctx, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllProductPrices", reflect.TypeOf((*MockProductRepository)(nil).FindAllProductPrices), ctx, name)
}

// FindCheapestProduct mocks base method.
func (m *MockProductRepository) FindCheapestProduct(ctx context.Context, name string) (models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindCheapestProduct", ctx, name)
	ret0, _ := ret[0].(models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindCheapestProduct indicates an expected call of FindCheapestProduct.
func (mr *MockProductRepositoryMockRecorder) FindCheapestProduct(ctx, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCheapestProduct", reflect.TypeOf((*MockProductRepository)(nil).FindCheapestProduct), ctx, name)
}

// FindCheapestProductByStore mocks base method.
func (m *MockProductRepository) FindCheapestProductByStore(ctx context.Context, name, store string) (models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindCheapestProductByStore", ctx, name, store)
	ret0, _ := ret[0].(models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindCheapestProductByStore indicates an expected call of FindCheapestProductByStore.
func (mr *MockProductRepositoryMockRecorder) FindCheapestProductByStore(ctx, name, store any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCheapestProductByStore", reflect.TypeOf((*MockProductRepository)(nil).FindCheapestProductByStore), ctx, name, store)
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
