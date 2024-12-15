// Code generated by mockery. DO NOT EDIT.

package biz

import (
	biz "github.com/TheTNB/panel/internal/biz"
	mock "github.com/stretchr/testify/mock"

	request "github.com/TheTNB/panel/internal/http/request"
)

// CertDNSRepo is an autogenerated mock type for the CertDNSRepo type
type CertDNSRepo struct {
	mock.Mock
}

type CertDNSRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *CertDNSRepo) EXPECT() *CertDNSRepo_Expecter {
	return &CertDNSRepo_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: req
func (_m *CertDNSRepo) Create(req *request.CertDNSCreate) (*biz.CertDNS, error) {
	ret := _m.Called(req)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *biz.CertDNS
	var r1 error
	if rf, ok := ret.Get(0).(func(*request.CertDNSCreate) (*biz.CertDNS, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(*request.CertDNSCreate) *biz.CertDNS); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*biz.CertDNS)
		}
	}

	if rf, ok := ret.Get(1).(func(*request.CertDNSCreate) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CertDNSRepo_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type CertDNSRepo_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - req *request.CertDNSCreate
func (_e *CertDNSRepo_Expecter) Create(req interface{}) *CertDNSRepo_Create_Call {
	return &CertDNSRepo_Create_Call{Call: _e.mock.On("Create", req)}
}

func (_c *CertDNSRepo_Create_Call) Run(run func(req *request.CertDNSCreate)) *CertDNSRepo_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*request.CertDNSCreate))
	})
	return _c
}

func (_c *CertDNSRepo_Create_Call) Return(_a0 *biz.CertDNS, _a1 error) *CertDNSRepo_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CertDNSRepo_Create_Call) RunAndReturn(run func(*request.CertDNSCreate) (*biz.CertDNS, error)) *CertDNSRepo_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: id
func (_m *CertDNSRepo) Delete(id uint) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CertDNSRepo_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type CertDNSRepo_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - id uint
func (_e *CertDNSRepo_Expecter) Delete(id interface{}) *CertDNSRepo_Delete_Call {
	return &CertDNSRepo_Delete_Call{Call: _e.mock.On("Delete", id)}
}

func (_c *CertDNSRepo_Delete_Call) Run(run func(id uint)) *CertDNSRepo_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *CertDNSRepo_Delete_Call) Return(_a0 error) *CertDNSRepo_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CertDNSRepo_Delete_Call) RunAndReturn(run func(uint) error) *CertDNSRepo_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: id
func (_m *CertDNSRepo) Get(id uint) (*biz.CertDNS, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *biz.CertDNS
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*biz.CertDNS, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) *biz.CertDNS); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*biz.CertDNS)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CertDNSRepo_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type CertDNSRepo_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - id uint
func (_e *CertDNSRepo_Expecter) Get(id interface{}) *CertDNSRepo_Get_Call {
	return &CertDNSRepo_Get_Call{Call: _e.mock.On("Get", id)}
}

func (_c *CertDNSRepo_Get_Call) Run(run func(id uint)) *CertDNSRepo_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *CertDNSRepo_Get_Call) Return(_a0 *biz.CertDNS, _a1 error) *CertDNSRepo_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CertDNSRepo_Get_Call) RunAndReturn(run func(uint) (*biz.CertDNS, error)) *CertDNSRepo_Get_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: page, limit
func (_m *CertDNSRepo) List(page uint, limit uint) ([]*biz.CertDNS, int64, error) {
	ret := _m.Called(page, limit)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []*biz.CertDNS
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(uint, uint) ([]*biz.CertDNS, int64, error)); ok {
		return rf(page, limit)
	}
	if rf, ok := ret.Get(0).(func(uint, uint) []*biz.CertDNS); ok {
		r0 = rf(page, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*biz.CertDNS)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, uint) int64); ok {
		r1 = rf(page, limit)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(uint, uint) error); ok {
		r2 = rf(page, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CertDNSRepo_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type CertDNSRepo_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - page uint
//   - limit uint
func (_e *CertDNSRepo_Expecter) List(page interface{}, limit interface{}) *CertDNSRepo_List_Call {
	return &CertDNSRepo_List_Call{Call: _e.mock.On("List", page, limit)}
}

func (_c *CertDNSRepo_List_Call) Run(run func(page uint, limit uint)) *CertDNSRepo_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint), args[1].(uint))
	})
	return _c
}

func (_c *CertDNSRepo_List_Call) Return(_a0 []*biz.CertDNS, _a1 int64, _a2 error) *CertDNSRepo_List_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *CertDNSRepo_List_Call) RunAndReturn(run func(uint, uint) ([]*biz.CertDNS, int64, error)) *CertDNSRepo_List_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: req
func (_m *CertDNSRepo) Update(req *request.CertDNSUpdate) error {
	ret := _m.Called(req)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*request.CertDNSUpdate) error); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CertDNSRepo_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type CertDNSRepo_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - req *request.CertDNSUpdate
func (_e *CertDNSRepo_Expecter) Update(req interface{}) *CertDNSRepo_Update_Call {
	return &CertDNSRepo_Update_Call{Call: _e.mock.On("Update", req)}
}

func (_c *CertDNSRepo_Update_Call) Run(run func(req *request.CertDNSUpdate)) *CertDNSRepo_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*request.CertDNSUpdate))
	})
	return _c
}

func (_c *CertDNSRepo_Update_Call) Return(_a0 error) *CertDNSRepo_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CertDNSRepo_Update_Call) RunAndReturn(run func(*request.CertDNSUpdate) error) *CertDNSRepo_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewCertDNSRepo creates a new instance of CertDNSRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCertDNSRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *CertDNSRepo {
	mock := &CertDNSRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
