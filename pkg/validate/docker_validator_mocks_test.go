// Code generated by MockGen. DO NOT EDIT.
// Source: docker_validator.go

// Package validate is a generated GoMock package.
package validate

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockdockerRepositoryValidator is a mock of dockerRepositoryValidator interface
type MockdockerRepositoryValidator struct {
	ctrl     *gomock.Controller
	recorder *MockdockerRepositoryValidatorMockRecorder
}

// MockdockerRepositoryValidatorMockRecorder is the mock recorder for MockdockerRepositoryValidator
type MockdockerRepositoryValidatorMockRecorder struct {
	mock *MockdockerRepositoryValidator
}

// NewMockdockerRepositoryValidator creates a new mock instance
func NewMockdockerRepositoryValidator(ctrl *gomock.Controller) *MockdockerRepositoryValidator {
	mock := &MockdockerRepositoryValidator{ctrl: ctrl}
	mock.recorder = &MockdockerRepositoryValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockdockerRepositoryValidator) EXPECT() *MockdockerRepositoryValidatorMockRecorder {
	return m.recorder
}

// repository mocks base method
func (m *MockdockerRepositoryValidator) repository(registry dockerRegistryAccount, service *dockerRegistryService) []error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "repository", registry, service)
	ret0, _ := ret[0].([]error)
	return ret0
}

// repository indicates an expected call of repository
func (mr *MockdockerRepositoryValidatorMockRecorder) repository(registry, service interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "repository", reflect.TypeOf((*MockdockerRepositoryValidator)(nil).repository), registry, service)
}

// imageTags mocks base method
func (m *MockdockerRepositoryValidator) imageTags(repository string, service *dockerRegistryService) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "imageTags", repository, service)
	ret0, _ := ret[0].(error)
	return ret0
}

// imageTags indicates an expected call of imageTags
func (mr *MockdockerRepositoryValidatorMockRecorder) imageTags(repository, service interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "imageTags", reflect.TypeOf((*MockdockerRepositoryValidator)(nil).imageTags), repository, service)
}