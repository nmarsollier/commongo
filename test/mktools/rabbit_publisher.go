// Code generated by MockGen. DO NOT EDIT.
// Source: ./rbt/publisher.go

// Package mktools is a generated GoMock package.
package mktools

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	log "github.com/nmarsollier/commongo/log"
)

// MockRabbitPublisher is a mock of RabbitPublisher interface.
type MockRabbitPublisher[T any] struct {
	ctrl     *gomock.Controller
	recorder *MockRabbitPublisherMockRecorder[T]
}

// MockRabbitPublisherMockRecorder is the mock recorder for MockRabbitPublisher.
type MockRabbitPublisherMockRecorder[T any] struct {
	mock *MockRabbitPublisher[T]
}

// NewMockRabbitPublisher creates a new mock instance.
func NewMockRabbitPublisher[T any](ctrl *gomock.Controller) *MockRabbitPublisher[T] {
	mock := &MockRabbitPublisher[T]{ctrl: ctrl}
	mock.recorder = &MockRabbitPublisherMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRabbitPublisher[T]) EXPECT() *MockRabbitPublisherMockRecorder[T] {
	return m.recorder
}

// Logger mocks base method.
func (m *MockRabbitPublisher[T]) Logger() log.LogRusEntry {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logger")
	ret0, _ := ret[0].(log.LogRusEntry)
	return ret0
}

// Logger indicates an expected call of Logger.
func (mr *MockRabbitPublisherMockRecorder[T]) Logger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logger", reflect.TypeOf((*MockRabbitPublisher[T])(nil).Logger))
}

// Publish mocks base method.
func (m *MockRabbitPublisher[T]) Publish(data T) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish.
func (mr *MockRabbitPublisherMockRecorder[T]) Publish(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockRabbitPublisher[T])(nil).Publish), data)
}

// PublishForResult mocks base method.
func (m *MockRabbitPublisher[T]) PublishForResult(data T, exchange, routingKey string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishForResult", data, exchange, routingKey)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishForResult indicates an expected call of PublishForResult.
func (mr *MockRabbitPublisherMockRecorder[T]) PublishForResult(data, exchange, routingKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishForResult", reflect.TypeOf((*MockRabbitPublisher[T])(nil).PublishForResult), data, exchange, routingKey)
}

// PublishTo mocks base method.
func (m *MockRabbitPublisher[T]) PublishTo(exchange, routingKey string, data T) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishTo", exchange, routingKey, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishTo indicates an expected call of PublishTo.
func (mr *MockRabbitPublisherMockRecorder[T]) PublishTo(exchange, routingKey, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishTo", reflect.TypeOf((*MockRabbitPublisher[T])(nil).PublishTo), exchange, routingKey, data)
}
