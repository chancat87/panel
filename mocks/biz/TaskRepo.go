// Code generated by mockery. DO NOT EDIT.

package biz

import (
	biz "github.com/TheTNB/panel/internal/biz"
	mock "github.com/stretchr/testify/mock"
)

// TaskRepo is an autogenerated mock type for the TaskRepo type
type TaskRepo struct {
	mock.Mock
}

type TaskRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *TaskRepo) EXPECT() *TaskRepo_Expecter {
	return &TaskRepo_Expecter{mock: &_m.Mock}
}

// Delete provides a mock function with given fields: id
func (_m *TaskRepo) Delete(id uint) error {
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

// TaskRepo_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type TaskRepo_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - id uint
func (_e *TaskRepo_Expecter) Delete(id interface{}) *TaskRepo_Delete_Call {
	return &TaskRepo_Delete_Call{Call: _e.mock.On("Delete", id)}
}

func (_c *TaskRepo_Delete_Call) Run(run func(id uint)) *TaskRepo_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *TaskRepo_Delete_Call) Return(_a0 error) *TaskRepo_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TaskRepo_Delete_Call) RunAndReturn(run func(uint) error) *TaskRepo_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: id
func (_m *TaskRepo) Get(id uint) (*biz.Task, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *biz.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*biz.Task, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) *biz.Task); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*biz.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TaskRepo_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type TaskRepo_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - id uint
func (_e *TaskRepo_Expecter) Get(id interface{}) *TaskRepo_Get_Call {
	return &TaskRepo_Get_Call{Call: _e.mock.On("Get", id)}
}

func (_c *TaskRepo_Get_Call) Run(run func(id uint)) *TaskRepo_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *TaskRepo_Get_Call) Return(_a0 *biz.Task, _a1 error) *TaskRepo_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TaskRepo_Get_Call) RunAndReturn(run func(uint) (*biz.Task, error)) *TaskRepo_Get_Call {
	_c.Call.Return(run)
	return _c
}

// HasRunningTask provides a mock function with no fields
func (_m *TaskRepo) HasRunningTask() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for HasRunningTask")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// TaskRepo_HasRunningTask_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HasRunningTask'
type TaskRepo_HasRunningTask_Call struct {
	*mock.Call
}

// HasRunningTask is a helper method to define mock.On call
func (_e *TaskRepo_Expecter) HasRunningTask() *TaskRepo_HasRunningTask_Call {
	return &TaskRepo_HasRunningTask_Call{Call: _e.mock.On("HasRunningTask")}
}

func (_c *TaskRepo_HasRunningTask_Call) Run(run func()) *TaskRepo_HasRunningTask_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *TaskRepo_HasRunningTask_Call) Return(_a0 bool) *TaskRepo_HasRunningTask_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TaskRepo_HasRunningTask_Call) RunAndReturn(run func() bool) *TaskRepo_HasRunningTask_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: page, limit
func (_m *TaskRepo) List(page uint, limit uint) ([]*biz.Task, int64, error) {
	ret := _m.Called(page, limit)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []*biz.Task
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(uint, uint) ([]*biz.Task, int64, error)); ok {
		return rf(page, limit)
	}
	if rf, ok := ret.Get(0).(func(uint, uint) []*biz.Task); ok {
		r0 = rf(page, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*biz.Task)
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

// TaskRepo_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type TaskRepo_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - page uint
//   - limit uint
func (_e *TaskRepo_Expecter) List(page interface{}, limit interface{}) *TaskRepo_List_Call {
	return &TaskRepo_List_Call{Call: _e.mock.On("List", page, limit)}
}

func (_c *TaskRepo_List_Call) Run(run func(page uint, limit uint)) *TaskRepo_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint), args[1].(uint))
	})
	return _c
}

func (_c *TaskRepo_List_Call) Return(_a0 []*biz.Task, _a1 int64, _a2 error) *TaskRepo_List_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *TaskRepo_List_Call) RunAndReturn(run func(uint, uint) ([]*biz.Task, int64, error)) *TaskRepo_List_Call {
	_c.Call.Return(run)
	return _c
}

// Push provides a mock function with given fields: task
func (_m *TaskRepo) Push(task *biz.Task) error {
	ret := _m.Called(task)

	if len(ret) == 0 {
		panic("no return value specified for Push")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*biz.Task) error); ok {
		r0 = rf(task)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TaskRepo_Push_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Push'
type TaskRepo_Push_Call struct {
	*mock.Call
}

// Push is a helper method to define mock.On call
//   - task *biz.Task
func (_e *TaskRepo_Expecter) Push(task interface{}) *TaskRepo_Push_Call {
	return &TaskRepo_Push_Call{Call: _e.mock.On("Push", task)}
}

func (_c *TaskRepo_Push_Call) Run(run func(task *biz.Task)) *TaskRepo_Push_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*biz.Task))
	})
	return _c
}

func (_c *TaskRepo_Push_Call) Return(_a0 error) *TaskRepo_Push_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TaskRepo_Push_Call) RunAndReturn(run func(*biz.Task) error) *TaskRepo_Push_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateStatus provides a mock function with given fields: id, status
func (_m *TaskRepo) UpdateStatus(id uint, status biz.TaskStatus) error {
	ret := _m.Called(id, status)

	if len(ret) == 0 {
		panic("no return value specified for UpdateStatus")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, biz.TaskStatus) error); ok {
		r0 = rf(id, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TaskRepo_UpdateStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateStatus'
type TaskRepo_UpdateStatus_Call struct {
	*mock.Call
}

// UpdateStatus is a helper method to define mock.On call
//   - id uint
//   - status biz.TaskStatus
func (_e *TaskRepo_Expecter) UpdateStatus(id interface{}, status interface{}) *TaskRepo_UpdateStatus_Call {
	return &TaskRepo_UpdateStatus_Call{Call: _e.mock.On("UpdateStatus", id, status)}
}

func (_c *TaskRepo_UpdateStatus_Call) Run(run func(id uint, status biz.TaskStatus)) *TaskRepo_UpdateStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint), args[1].(biz.TaskStatus))
	})
	return _c
}

func (_c *TaskRepo_UpdateStatus_Call) Return(_a0 error) *TaskRepo_UpdateStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TaskRepo_UpdateStatus_Call) RunAndReturn(run func(uint, biz.TaskStatus) error) *TaskRepo_UpdateStatus_Call {
	_c.Call.Return(run)
	return _c
}

// NewTaskRepo creates a new instance of TaskRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTaskRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *TaskRepo {
	mock := &TaskRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
