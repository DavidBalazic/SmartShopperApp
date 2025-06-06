// Code generated by MockGen. DO NOT EDIT.
// Source: internal/rabbitmq/publisher.go
//
// Generated by this command:
//
//	mockgen -source=internal/rabbitmq/publisher.go -destination=mocks/mock_publisher.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	dtos "github.com/DavidBalazic/SmartShopperApp/internal/dtos"
	gomock "go.uber.org/mock/gomock"
)

// MockPublisher is a mock of Publisher interface.
type MockPublisher struct {
	ctrl     *gomock.Controller
	recorder *MockPublisherMockRecorder
}

// MockPublisherMockRecorder is the mock recorder for MockPublisher.
type MockPublisherMockRecorder struct {
	mock *MockPublisher
}

// NewMockPublisher creates a new mock instance.
func NewMockPublisher(ctrl *gomock.Controller) *MockPublisher {
	mock := &MockPublisher{ctrl: ctrl}
	mock.recorder = &MockPublisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPublisher) EXPECT() *MockPublisherMockRecorder {
	return m.recorder
}

// PublishMultipleProducts mocks base method.
func (m *MockPublisher) PublishMultipleProducts(messages []dtos.ProductMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishMultipleProducts", messages)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishMultipleProducts indicates an expected call of PublishMultipleProducts.
func (mr *MockPublisherMockRecorder) PublishMultipleProducts(messages any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishMultipleProducts", reflect.TypeOf((*MockPublisher)(nil).PublishMultipleProducts), messages)
}

// PublishSingleProduct mocks base method.
func (m *MockPublisher) PublishSingleProduct(message dtos.ProductMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishSingleProduct", message)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishSingleProduct indicates an expected call of PublishSingleProduct.
func (mr *MockPublisherMockRecorder) PublishSingleProduct(message any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishSingleProduct", reflect.TypeOf((*MockPublisher)(nil).PublishSingleProduct), message)
}
