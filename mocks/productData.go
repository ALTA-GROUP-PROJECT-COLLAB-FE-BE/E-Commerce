// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	product "project/e-commerce/features/product"

	mock "github.com/stretchr/testify/mock"
)

// ProductData is an autogenerated mock type for the Data type
type ProductData struct {
	mock.Mock
}

// DeleteDataByIdDB provides a mock function with given fields: idUser, idFromToken
func (_m *ProductData) DeleteDataByIdDB(idUser int, idFromToken int) (int, error) {
	ret := _m.Called(idUser, idFromToken)

	var r0 int
	if rf, ok := ret.Get(0).(func(int, int) int); ok {
		r0 = rf(idUser, idFromToken)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(idUser, idFromToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDataByMeDB provides a mock function with given fields: idUser
func (_m *ProductData) GetDataByMeDB(idUser int) ([]product.Core, error) {
	ret := _m.Called(idUser)

	var r0 []product.Core
	if rf, ok := ret.Get(0).(func(int) []product.Core); ok {
		r0 = rf(idUser)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]product.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(idUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertData provides a mock function with given fields: data
func (_m *ProductData) InsertData(data product.Core) (int, error) {
	ret := _m.Called(data)

	var r0 int
	if rf, ok := ret.Get(0).(func(product.Core) int); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(product.Core) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectData provides a mock function with given fields: limit, offset
func (_m *ProductData) SelectData(limit int, offset int) ([]product.Core, error) {
	ret := _m.Called(limit, offset)

	var r0 []product.Core
	if rf, ok := ret.Get(0).(func(int, int) []product.Core); ok {
		r0 = rf(limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]product.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectDataById provides a mock function with given fields: idProd
func (_m *ProductData) SelectDataById(idProd int) (product.Core, error) {
	ret := _m.Called(idProd)

	var r0 product.Core
	if rf, ok := ret.Get(0).(func(int) product.Core); ok {
		r0 = rf(idProd)
	} else {
		r0 = ret.Get(0).(product.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(idProd)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDataDB provides a mock function with given fields: data, idProd, idFromToken
func (_m *ProductData) UpdateDataDB(data map[string]interface{}, idProd int, idFromToken int) (int, error) {
	ret := _m.Called(data, idProd, idFromToken)

	var r0 int
	if rf, ok := ret.Get(0).(func(map[string]interface{}, int, int) int); ok {
		r0 = rf(data, idProd, idFromToken)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(map[string]interface{}, int, int) error); ok {
		r1 = rf(data, idProd, idFromToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewProductData interface {
	mock.TestingT
	Cleanup(func())
}

// NewProductData creates a new instance of ProductData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProductData(t mockConstructorTestingTNewProductData) *ProductData {
	mock := &ProductData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}