// Code generated by mockery. DO NOT EDIT.

package biz

import (
	mock "github.com/stretchr/testify/mock"
	request "github.com/tnborg/panel/internal/http/request"
	types "github.com/tnborg/panel/pkg/types"
)

// ContainerNetworkRepo is an autogenerated mock type for the ContainerNetworkRepo type
type ContainerNetworkRepo struct {
	mock.Mock
}

type ContainerNetworkRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *ContainerNetworkRepo) EXPECT() *ContainerNetworkRepo_Expecter {
	return &ContainerNetworkRepo_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: req
func (_m *ContainerNetworkRepo) Create(req *request.ContainerNetworkCreate) (string, error) {
	ret := _m.Called(req)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*request.ContainerNetworkCreate) (string, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(*request.ContainerNetworkCreate) string); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*request.ContainerNetworkCreate) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContainerNetworkRepo_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type ContainerNetworkRepo_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - req *request.ContainerNetworkCreate
func (_e *ContainerNetworkRepo_Expecter) Create(req interface{}) *ContainerNetworkRepo_Create_Call {
	return &ContainerNetworkRepo_Create_Call{Call: _e.mock.On("Create", req)}
}

func (_c *ContainerNetworkRepo_Create_Call) Run(run func(req *request.ContainerNetworkCreate)) *ContainerNetworkRepo_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*request.ContainerNetworkCreate))
	})
	return _c
}

func (_c *ContainerNetworkRepo_Create_Call) Return(_a0 string, _a1 error) *ContainerNetworkRepo_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ContainerNetworkRepo_Create_Call) RunAndReturn(run func(*request.ContainerNetworkCreate) (string, error)) *ContainerNetworkRepo_Create_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with no fields
func (_m *ContainerNetworkRepo) List() ([]types.ContainerNetwork, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []types.ContainerNetwork
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]types.ContainerNetwork, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []types.ContainerNetwork); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.ContainerNetwork)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContainerNetworkRepo_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type ContainerNetworkRepo_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
func (_e *ContainerNetworkRepo_Expecter) List() *ContainerNetworkRepo_List_Call {
	return &ContainerNetworkRepo_List_Call{Call: _e.mock.On("List")}
}

func (_c *ContainerNetworkRepo_List_Call) Run(run func()) *ContainerNetworkRepo_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ContainerNetworkRepo_List_Call) Return(_a0 []types.ContainerNetwork, _a1 error) *ContainerNetworkRepo_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ContainerNetworkRepo_List_Call) RunAndReturn(run func() ([]types.ContainerNetwork, error)) *ContainerNetworkRepo_List_Call {
	_c.Call.Return(run)
	return _c
}

// Prune provides a mock function with no fields
func (_m *ContainerNetworkRepo) Prune() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Prune")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ContainerNetworkRepo_Prune_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Prune'
type ContainerNetworkRepo_Prune_Call struct {
	*mock.Call
}

// Prune is a helper method to define mock.On call
func (_e *ContainerNetworkRepo_Expecter) Prune() *ContainerNetworkRepo_Prune_Call {
	return &ContainerNetworkRepo_Prune_Call{Call: _e.mock.On("Prune")}
}

func (_c *ContainerNetworkRepo_Prune_Call) Run(run func()) *ContainerNetworkRepo_Prune_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ContainerNetworkRepo_Prune_Call) Return(_a0 error) *ContainerNetworkRepo_Prune_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ContainerNetworkRepo_Prune_Call) RunAndReturn(run func() error) *ContainerNetworkRepo_Prune_Call {
	_c.Call.Return(run)
	return _c
}

// Remove provides a mock function with given fields: id
func (_m *ContainerNetworkRepo) Remove(id string) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Remove")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ContainerNetworkRepo_Remove_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Remove'
type ContainerNetworkRepo_Remove_Call struct {
	*mock.Call
}

// Remove is a helper method to define mock.On call
//   - id string
func (_e *ContainerNetworkRepo_Expecter) Remove(id interface{}) *ContainerNetworkRepo_Remove_Call {
	return &ContainerNetworkRepo_Remove_Call{Call: _e.mock.On("Remove", id)}
}

func (_c *ContainerNetworkRepo_Remove_Call) Run(run func(id string)) *ContainerNetworkRepo_Remove_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *ContainerNetworkRepo_Remove_Call) Return(_a0 error) *ContainerNetworkRepo_Remove_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ContainerNetworkRepo_Remove_Call) RunAndReturn(run func(string) error) *ContainerNetworkRepo_Remove_Call {
	_c.Call.Return(run)
	return _c
}

// NewContainerNetworkRepo creates a new instance of ContainerNetworkRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewContainerNetworkRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *ContainerNetworkRepo {
	mock := &ContainerNetworkRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
