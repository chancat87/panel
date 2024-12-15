// Code generated by mockery. DO NOT EDIT.

package biz

import (
	context "context"

	biz "github.com/TheTNB/panel/internal/biz"

	mock "github.com/stretchr/testify/mock"

	request "github.com/TheTNB/panel/internal/http/request"

	types "github.com/TheTNB/panel/pkg/types"
)

// WebsiteRepo is an autogenerated mock type for the WebsiteRepo type
type WebsiteRepo struct {
	mock.Mock
}

type WebsiteRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *WebsiteRepo) EXPECT() *WebsiteRepo_Expecter {
	return &WebsiteRepo_Expecter{mock: &_m.Mock}
}

// ClearLog provides a mock function with given fields: id
func (_m *WebsiteRepo) ClearLog(id uint) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for ClearLog")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WebsiteRepo_ClearLog_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ClearLog'
type WebsiteRepo_ClearLog_Call struct {
	*mock.Call
}

// ClearLog is a helper method to define mock.On call
//   - id uint
func (_e *WebsiteRepo_Expecter) ClearLog(id interface{}) *WebsiteRepo_ClearLog_Call {
	return &WebsiteRepo_ClearLog_Call{Call: _e.mock.On("ClearLog", id)}
}

func (_c *WebsiteRepo_ClearLog_Call) Run(run func(id uint)) *WebsiteRepo_ClearLog_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *WebsiteRepo_ClearLog_Call) Return(_a0 error) *WebsiteRepo_ClearLog_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *WebsiteRepo_ClearLog_Call) RunAndReturn(run func(uint) error) *WebsiteRepo_ClearLog_Call {
	_c.Call.Return(run)
	return _c
}

// Count provides a mock function with no fields
func (_m *WebsiteRepo) Count() (int64, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Count")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func() (int64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WebsiteRepo_Count_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Count'
type WebsiteRepo_Count_Call struct {
	*mock.Call
}

// Count is a helper method to define mock.On call
func (_e *WebsiteRepo_Expecter) Count() *WebsiteRepo_Count_Call {
	return &WebsiteRepo_Count_Call{Call: _e.mock.On("Count")}
}

func (_c *WebsiteRepo_Count_Call) Run(run func()) *WebsiteRepo_Count_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *WebsiteRepo_Count_Call) Return(_a0 int64, _a1 error) *WebsiteRepo_Count_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *WebsiteRepo_Count_Call) RunAndReturn(run func() (int64, error)) *WebsiteRepo_Count_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: req
func (_m *WebsiteRepo) Create(req *request.WebsiteCreate) (*biz.Website, error) {
	ret := _m.Called(req)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *biz.Website
	var r1 error
	if rf, ok := ret.Get(0).(func(*request.WebsiteCreate) (*biz.Website, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(*request.WebsiteCreate) *biz.Website); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*biz.Website)
		}
	}

	if rf, ok := ret.Get(1).(func(*request.WebsiteCreate) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WebsiteRepo_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type WebsiteRepo_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - req *request.WebsiteCreate
func (_e *WebsiteRepo_Expecter) Create(req interface{}) *WebsiteRepo_Create_Call {
	return &WebsiteRepo_Create_Call{Call: _e.mock.On("Create", req)}
}

func (_c *WebsiteRepo_Create_Call) Run(run func(req *request.WebsiteCreate)) *WebsiteRepo_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*request.WebsiteCreate))
	})
	return _c
}

func (_c *WebsiteRepo_Create_Call) Return(_a0 *biz.Website, _a1 error) *WebsiteRepo_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *WebsiteRepo_Create_Call) RunAndReturn(run func(*request.WebsiteCreate) (*biz.Website, error)) *WebsiteRepo_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: req
func (_m *WebsiteRepo) Delete(req *request.WebsiteDelete) error {
	ret := _m.Called(req)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*request.WebsiteDelete) error); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WebsiteRepo_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type WebsiteRepo_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - req *request.WebsiteDelete
func (_e *WebsiteRepo_Expecter) Delete(req interface{}) *WebsiteRepo_Delete_Call {
	return &WebsiteRepo_Delete_Call{Call: _e.mock.On("Delete", req)}
}

func (_c *WebsiteRepo_Delete_Call) Run(run func(req *request.WebsiteDelete)) *WebsiteRepo_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*request.WebsiteDelete))
	})
	return _c
}

func (_c *WebsiteRepo_Delete_Call) Return(_a0 error) *WebsiteRepo_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *WebsiteRepo_Delete_Call) RunAndReturn(run func(*request.WebsiteDelete) error) *WebsiteRepo_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: id
func (_m *WebsiteRepo) Get(id uint) (*types.WebsiteSetting, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *types.WebsiteSetting
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*types.WebsiteSetting, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) *types.WebsiteSetting); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.WebsiteSetting)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WebsiteRepo_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type WebsiteRepo_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - id uint
func (_e *WebsiteRepo_Expecter) Get(id interface{}) *WebsiteRepo_Get_Call {
	return &WebsiteRepo_Get_Call{Call: _e.mock.On("Get", id)}
}

func (_c *WebsiteRepo_Get_Call) Run(run func(id uint)) *WebsiteRepo_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *WebsiteRepo_Get_Call) Return(_a0 *types.WebsiteSetting, _a1 error) *WebsiteRepo_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *WebsiteRepo_Get_Call) RunAndReturn(run func(uint) (*types.WebsiteSetting, error)) *WebsiteRepo_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetByName provides a mock function with given fields: name
func (_m *WebsiteRepo) GetByName(name string) (*types.WebsiteSetting, error) {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for GetByName")
	}

	var r0 *types.WebsiteSetting
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*types.WebsiteSetting, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) *types.WebsiteSetting); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.WebsiteSetting)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WebsiteRepo_GetByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByName'
type WebsiteRepo_GetByName_Call struct {
	*mock.Call
}

// GetByName is a helper method to define mock.On call
//   - name string
func (_e *WebsiteRepo_Expecter) GetByName(name interface{}) *WebsiteRepo_GetByName_Call {
	return &WebsiteRepo_GetByName_Call{Call: _e.mock.On("GetByName", name)}
}

func (_c *WebsiteRepo_GetByName_Call) Run(run func(name string)) *WebsiteRepo_GetByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *WebsiteRepo_GetByName_Call) Return(_a0 *types.WebsiteSetting, _a1 error) *WebsiteRepo_GetByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *WebsiteRepo_GetByName_Call) RunAndReturn(run func(string) (*types.WebsiteSetting, error)) *WebsiteRepo_GetByName_Call {
	_c.Call.Return(run)
	return _c
}

// GetRewrites provides a mock function with no fields
func (_m *WebsiteRepo) GetRewrites() (map[string]string, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetRewrites")
	}

	var r0 map[string]string
	var r1 error
	if rf, ok := ret.Get(0).(func() (map[string]string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() map[string]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WebsiteRepo_GetRewrites_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRewrites'
type WebsiteRepo_GetRewrites_Call struct {
	*mock.Call
}

// GetRewrites is a helper method to define mock.On call
func (_e *WebsiteRepo_Expecter) GetRewrites() *WebsiteRepo_GetRewrites_Call {
	return &WebsiteRepo_GetRewrites_Call{Call: _e.mock.On("GetRewrites")}
}

func (_c *WebsiteRepo_GetRewrites_Call) Run(run func()) *WebsiteRepo_GetRewrites_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *WebsiteRepo_GetRewrites_Call) Return(_a0 map[string]string, _a1 error) *WebsiteRepo_GetRewrites_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *WebsiteRepo_GetRewrites_Call) RunAndReturn(run func() (map[string]string, error)) *WebsiteRepo_GetRewrites_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: page, limit
func (_m *WebsiteRepo) List(page uint, limit uint) ([]*biz.Website, int64, error) {
	ret := _m.Called(page, limit)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []*biz.Website
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(uint, uint) ([]*biz.Website, int64, error)); ok {
		return rf(page, limit)
	}
	if rf, ok := ret.Get(0).(func(uint, uint) []*biz.Website); ok {
		r0 = rf(page, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*biz.Website)
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

// WebsiteRepo_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type WebsiteRepo_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - page uint
//   - limit uint
func (_e *WebsiteRepo_Expecter) List(page interface{}, limit interface{}) *WebsiteRepo_List_Call {
	return &WebsiteRepo_List_Call{Call: _e.mock.On("List", page, limit)}
}

func (_c *WebsiteRepo_List_Call) Run(run func(page uint, limit uint)) *WebsiteRepo_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint), args[1].(uint))
	})
	return _c
}

func (_c *WebsiteRepo_List_Call) Return(_a0 []*biz.Website, _a1 int64, _a2 error) *WebsiteRepo_List_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *WebsiteRepo_List_Call) RunAndReturn(run func(uint, uint) ([]*biz.Website, int64, error)) *WebsiteRepo_List_Call {
	_c.Call.Return(run)
	return _c
}

// ObtainCert provides a mock function with given fields: ctx, id
func (_m *WebsiteRepo) ObtainCert(ctx context.Context, id uint) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for ObtainCert")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WebsiteRepo_ObtainCert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ObtainCert'
type WebsiteRepo_ObtainCert_Call struct {
	*mock.Call
}

// ObtainCert is a helper method to define mock.On call
//   - ctx context.Context
//   - id uint
func (_e *WebsiteRepo_Expecter) ObtainCert(ctx interface{}, id interface{}) *WebsiteRepo_ObtainCert_Call {
	return &WebsiteRepo_ObtainCert_Call{Call: _e.mock.On("ObtainCert", ctx, id)}
}

func (_c *WebsiteRepo_ObtainCert_Call) Run(run func(ctx context.Context, id uint)) *WebsiteRepo_ObtainCert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint))
	})
	return _c
}

func (_c *WebsiteRepo_ObtainCert_Call) Return(_a0 error) *WebsiteRepo_ObtainCert_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *WebsiteRepo_ObtainCert_Call) RunAndReturn(run func(context.Context, uint) error) *WebsiteRepo_ObtainCert_Call {
	_c.Call.Return(run)
	return _c
}

// ResetConfig provides a mock function with given fields: id
func (_m *WebsiteRepo) ResetConfig(id uint) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for ResetConfig")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WebsiteRepo_ResetConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ResetConfig'
type WebsiteRepo_ResetConfig_Call struct {
	*mock.Call
}

// ResetConfig is a helper method to define mock.On call
//   - id uint
func (_e *WebsiteRepo_Expecter) ResetConfig(id interface{}) *WebsiteRepo_ResetConfig_Call {
	return &WebsiteRepo_ResetConfig_Call{Call: _e.mock.On("ResetConfig", id)}
}

func (_c *WebsiteRepo_ResetConfig_Call) Run(run func(id uint)) *WebsiteRepo_ResetConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *WebsiteRepo_ResetConfig_Call) Return(_a0 error) *WebsiteRepo_ResetConfig_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *WebsiteRepo_ResetConfig_Call) RunAndReturn(run func(uint) error) *WebsiteRepo_ResetConfig_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: req
func (_m *WebsiteRepo) Update(req *request.WebsiteUpdate) error {
	ret := _m.Called(req)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*request.WebsiteUpdate) error); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WebsiteRepo_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type WebsiteRepo_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - req *request.WebsiteUpdate
func (_e *WebsiteRepo_Expecter) Update(req interface{}) *WebsiteRepo_Update_Call {
	return &WebsiteRepo_Update_Call{Call: _e.mock.On("Update", req)}
}

func (_c *WebsiteRepo_Update_Call) Run(run func(req *request.WebsiteUpdate)) *WebsiteRepo_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*request.WebsiteUpdate))
	})
	return _c
}

func (_c *WebsiteRepo_Update_Call) Return(_a0 error) *WebsiteRepo_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *WebsiteRepo_Update_Call) RunAndReturn(run func(*request.WebsiteUpdate) error) *WebsiteRepo_Update_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateDefaultConfig provides a mock function with given fields: req
func (_m *WebsiteRepo) UpdateDefaultConfig(req *request.WebsiteDefaultConfig) error {
	ret := _m.Called(req)

	if len(ret) == 0 {
		panic("no return value specified for UpdateDefaultConfig")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*request.WebsiteDefaultConfig) error); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WebsiteRepo_UpdateDefaultConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateDefaultConfig'
type WebsiteRepo_UpdateDefaultConfig_Call struct {
	*mock.Call
}

// UpdateDefaultConfig is a helper method to define mock.On call
//   - req *request.WebsiteDefaultConfig
func (_e *WebsiteRepo_Expecter) UpdateDefaultConfig(req interface{}) *WebsiteRepo_UpdateDefaultConfig_Call {
	return &WebsiteRepo_UpdateDefaultConfig_Call{Call: _e.mock.On("UpdateDefaultConfig", req)}
}

func (_c *WebsiteRepo_UpdateDefaultConfig_Call) Run(run func(req *request.WebsiteDefaultConfig)) *WebsiteRepo_UpdateDefaultConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*request.WebsiteDefaultConfig))
	})
	return _c
}

func (_c *WebsiteRepo_UpdateDefaultConfig_Call) Return(_a0 error) *WebsiteRepo_UpdateDefaultConfig_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *WebsiteRepo_UpdateDefaultConfig_Call) RunAndReturn(run func(*request.WebsiteDefaultConfig) error) *WebsiteRepo_UpdateDefaultConfig_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateRemark provides a mock function with given fields: id, remark
func (_m *WebsiteRepo) UpdateRemark(id uint, remark string) error {
	ret := _m.Called(id, remark)

	if len(ret) == 0 {
		panic("no return value specified for UpdateRemark")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, string) error); ok {
		r0 = rf(id, remark)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WebsiteRepo_UpdateRemark_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateRemark'
type WebsiteRepo_UpdateRemark_Call struct {
	*mock.Call
}

// UpdateRemark is a helper method to define mock.On call
//   - id uint
//   - remark string
func (_e *WebsiteRepo_Expecter) UpdateRemark(id interface{}, remark interface{}) *WebsiteRepo_UpdateRemark_Call {
	return &WebsiteRepo_UpdateRemark_Call{Call: _e.mock.On("UpdateRemark", id, remark)}
}

func (_c *WebsiteRepo_UpdateRemark_Call) Run(run func(id uint, remark string)) *WebsiteRepo_UpdateRemark_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint), args[1].(string))
	})
	return _c
}

func (_c *WebsiteRepo_UpdateRemark_Call) Return(_a0 error) *WebsiteRepo_UpdateRemark_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *WebsiteRepo_UpdateRemark_Call) RunAndReturn(run func(uint, string) error) *WebsiteRepo_UpdateRemark_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateStatus provides a mock function with given fields: id, status
func (_m *WebsiteRepo) UpdateStatus(id uint, status bool) error {
	ret := _m.Called(id, status)

	if len(ret) == 0 {
		panic("no return value specified for UpdateStatus")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, bool) error); ok {
		r0 = rf(id, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WebsiteRepo_UpdateStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateStatus'
type WebsiteRepo_UpdateStatus_Call struct {
	*mock.Call
}

// UpdateStatus is a helper method to define mock.On call
//   - id uint
//   - status bool
func (_e *WebsiteRepo_Expecter) UpdateStatus(id interface{}, status interface{}) *WebsiteRepo_UpdateStatus_Call {
	return &WebsiteRepo_UpdateStatus_Call{Call: _e.mock.On("UpdateStatus", id, status)}
}

func (_c *WebsiteRepo_UpdateStatus_Call) Run(run func(id uint, status bool)) *WebsiteRepo_UpdateStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint), args[1].(bool))
	})
	return _c
}

func (_c *WebsiteRepo_UpdateStatus_Call) Return(_a0 error) *WebsiteRepo_UpdateStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *WebsiteRepo_UpdateStatus_Call) RunAndReturn(run func(uint, bool) error) *WebsiteRepo_UpdateStatus_Call {
	_c.Call.Return(run)
	return _c
}

// NewWebsiteRepo creates a new instance of WebsiteRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWebsiteRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *WebsiteRepo {
	mock := &WebsiteRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
