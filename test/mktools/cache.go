// FILE AUTOMATICALLY GENERATED. ADAPTER TO USE GENERICS
package mktools

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCache is a mock of Cache interface.
type MockCache[T any] struct {
	ctrl     *gomock.Controller
	recorder *MockCacheMockRecorder[T]
}

// MockCacheMockRecorder is the mock recorder for MockCache.
type MockCacheMockRecorder[T any] struct {
	mock *MockCache[T]
}

// NewMockCache creates a new mock instance.
func NewMockCache[T any](ctrl *gomock.Controller) *MockCache[T] {
	mock := &MockCache[T]{ctrl: ctrl}
	mock.recorder = &MockCacheMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCache[T]) EXPECT() *MockCacheMockRecorder[T] {
	return m.recorder
}

// Add mocks base method.
func (m *MockCache[T]) Add(key string, value *T) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockCacheMockRecorder[T]) Add(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockCache[T])(nil).Add), key, value)
}

// Get mocks base method.
func (m *MockCache[T]) Get(key string) (*T, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(*T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCacheMockRecorder[T]) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCache[T])(nil).Get), key)
}

// Remove mocks base method.
func (m *MockCache[T]) Remove(key string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Remove", key)
}

// Remove indicates an expected call of Remove.
func (mr *MockCacheMockRecorder[T]) Remove(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockCache[T])(nil).Remove), key)
}
