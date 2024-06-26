/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/digitalocean/godo (interfaces: LoadBalancersService)

// Package mock_networking is a generated GoMock package.
package mock_networking

import (
	context "context"
	reflect "reflect"

	godo "github.com/digitalocean/godo"
	gomock "github.com/golang/mock/gomock"
)

// MockLoadBalancersService is a mock of LoadBalancersService interface.
type MockLoadBalancersService struct {
	ctrl     *gomock.Controller
	recorder *MockLoadBalancersServiceMockRecorder
}

// MockLoadBalancersServiceMockRecorder is the mock recorder for MockLoadBalancersService.
type MockLoadBalancersServiceMockRecorder struct {
	mock *MockLoadBalancersService
}

// NewMockLoadBalancersService creates a new mock instance.
func NewMockLoadBalancersService(ctrl *gomock.Controller) *MockLoadBalancersService {
	mock := &MockLoadBalancersService{ctrl: ctrl}
	mock.recorder = &MockLoadBalancersServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoadBalancersService) EXPECT() *MockLoadBalancersServiceMockRecorder {
	return m.recorder
}

// AddDroplets mocks base method.
func (m *MockLoadBalancersService) AddDroplets(arg0 context.Context, arg1 string, arg2 ...int) (*godo.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddDroplets", varargs...)
	ret0, _ := ret[0].(*godo.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddDroplets indicates an expected call of AddDroplets.
func (mr *MockLoadBalancersServiceMockRecorder) AddDroplets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDroplets", reflect.TypeOf((*MockLoadBalancersService)(nil).AddDroplets), varargs...)
}

// AddForwardingRules mocks base method.
func (m *MockLoadBalancersService) AddForwardingRules(arg0 context.Context, arg1 string, arg2 ...godo.ForwardingRule) (*godo.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddForwardingRules", varargs...)
	ret0, _ := ret[0].(*godo.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddForwardingRules indicates an expected call of AddForwardingRules.
func (mr *MockLoadBalancersServiceMockRecorder) AddForwardingRules(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddForwardingRules", reflect.TypeOf((*MockLoadBalancersService)(nil).AddForwardingRules), varargs...)
}

// Create mocks base method.
func (m *MockLoadBalancersService) Create(arg0 context.Context, arg1 *godo.LoadBalancerRequest) (*godo.LoadBalancer, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*godo.LoadBalancer)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Create indicates an expected call of Create.
func (mr *MockLoadBalancersServiceMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockLoadBalancersService)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockLoadBalancersService) Delete(arg0 context.Context, arg1 string) (*godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(*godo.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockLoadBalancersServiceMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockLoadBalancersService)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *MockLoadBalancersService) Get(arg0 context.Context, arg1 string) (*godo.LoadBalancer, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*godo.LoadBalancer)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockLoadBalancersServiceMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockLoadBalancersService)(nil).Get), arg0, arg1)
}

// List mocks base method.
func (m *MockLoadBalancersService) List(arg0 context.Context, arg1 *godo.ListOptions) ([]godo.LoadBalancer, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]godo.LoadBalancer)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockLoadBalancersServiceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockLoadBalancersService)(nil).List), arg0, arg1)
}

// PurgeCache mocks base method.
func (m *MockLoadBalancersService) PurgeCache(arg0 context.Context, arg1 string) (*godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PurgeCache", arg0, arg1)
	ret0, _ := ret[0].(*godo.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PurgeCache indicates an expected call of PurgeCache.
func (mr *MockLoadBalancersServiceMockRecorder) PurgeCache(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PurgeCache", reflect.TypeOf((*MockLoadBalancersService)(nil).PurgeCache), arg0, arg1)
}

// RemoveDroplets mocks base method.
func (m *MockLoadBalancersService) RemoveDroplets(arg0 context.Context, arg1 string, arg2 ...int) (*godo.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveDroplets", varargs...)
	ret0, _ := ret[0].(*godo.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveDroplets indicates an expected call of RemoveDroplets.
func (mr *MockLoadBalancersServiceMockRecorder) RemoveDroplets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveDroplets", reflect.TypeOf((*MockLoadBalancersService)(nil).RemoveDroplets), varargs...)
}

// RemoveForwardingRules mocks base method.
func (m *MockLoadBalancersService) RemoveForwardingRules(arg0 context.Context, arg1 string, arg2 ...godo.ForwardingRule) (*godo.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveForwardingRules", varargs...)
	ret0, _ := ret[0].(*godo.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveForwardingRules indicates an expected call of RemoveForwardingRules.
func (mr *MockLoadBalancersServiceMockRecorder) RemoveForwardingRules(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveForwardingRules", reflect.TypeOf((*MockLoadBalancersService)(nil).RemoveForwardingRules), varargs...)
}

// Update mocks base method.
func (m *MockLoadBalancersService) Update(arg0 context.Context, arg1 string, arg2 *godo.LoadBalancerRequest) (*godo.LoadBalancer, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*godo.LoadBalancer)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Update indicates an expected call of Update.
func (mr *MockLoadBalancersServiceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockLoadBalancersService)(nil).Update), arg0, arg1, arg2)
}
