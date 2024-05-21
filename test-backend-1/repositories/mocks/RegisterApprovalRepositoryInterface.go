// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks_repositories

import (
	entities "test-backend-1/entities"
	db "test-backend-1/utils/db"

	mock "github.com/stretchr/testify/mock"

	repositories "test-backend-1/repositories"
)

// RegisterApprovalRepositoryInterface is an autogenerated mock type for the RegisterApprovalRepositoryInterface type
type RegisterApprovalRepositoryInterface struct {
	mock.Mock
}

type RegisterApprovalRepositoryInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *RegisterApprovalRepositoryInterface) EXPECT() *RegisterApprovalRepositoryInterface_Expecter {
	return &RegisterApprovalRepositoryInterface_Expecter{mock: &_m.Mock}
}

// Begin provides a mock function with given fields: transactor
func (_m *RegisterApprovalRepositoryInterface) Begin(transactor db.Transactor) repositories.RegisterApprovalRepositoryInterface {
	ret := _m.Called(transactor)

	var r0 repositories.RegisterApprovalRepositoryInterface
	if rf, ok := ret.Get(0).(func(db.Transactor) repositories.RegisterApprovalRepositoryInterface); ok {
		r0 = rf(transactor)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(repositories.RegisterApprovalRepositoryInterface)
		}
	}

	return r0
}

// RegisterApprovalRepositoryInterface_Begin_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Begin'
type RegisterApprovalRepositoryInterface_Begin_Call struct {
	*mock.Call
}

// Begin is a helper method to define mock.On call
//   - transactor db.Transactor
func (_e *RegisterApprovalRepositoryInterface_Expecter) Begin(transactor interface{}) *RegisterApprovalRepositoryInterface_Begin_Call {
	return &RegisterApprovalRepositoryInterface_Begin_Call{Call: _e.mock.On("Begin", transactor)}
}

func (_c *RegisterApprovalRepositoryInterface_Begin_Call) Run(run func(transactor db.Transactor)) *RegisterApprovalRepositoryInterface_Begin_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(db.Transactor))
	})
	return _c
}

func (_c *RegisterApprovalRepositoryInterface_Begin_Call) Return(_a0 repositories.RegisterApprovalRepositoryInterface) *RegisterApprovalRepositoryInterface_Begin_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RegisterApprovalRepositoryInterface_Begin_Call) RunAndReturn(run func(db.Transactor) repositories.RegisterApprovalRepositoryInterface) *RegisterApprovalRepositoryInterface_Begin_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: approval
func (_m *RegisterApprovalRepositoryInterface) Create(approval entities.RegisterApproval) (entities.RegisterApproval, error) {
	ret := _m.Called(approval)

	var r0 entities.RegisterApproval
	var r1 error
	if rf, ok := ret.Get(0).(func(entities.RegisterApproval) (entities.RegisterApproval, error)); ok {
		return rf(approval)
	}
	if rf, ok := ret.Get(0).(func(entities.RegisterApproval) entities.RegisterApproval); ok {
		r0 = rf(approval)
	} else {
		r0 = ret.Get(0).(entities.RegisterApproval)
	}

	if rf, ok := ret.Get(1).(func(entities.RegisterApproval) error); ok {
		r1 = rf(approval)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterApprovalRepositoryInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type RegisterApprovalRepositoryInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - approval entities.RegisterApproval
func (_e *RegisterApprovalRepositoryInterface_Expecter) Create(approval interface{}) *RegisterApprovalRepositoryInterface_Create_Call {
	return &RegisterApprovalRepositoryInterface_Create_Call{Call: _e.mock.On("Create", approval)}
}

func (_c *RegisterApprovalRepositoryInterface_Create_Call) Run(run func(approval entities.RegisterApproval)) *RegisterApprovalRepositoryInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(entities.RegisterApproval))
	})
	return _c
}

func (_c *RegisterApprovalRepositoryInterface_Create_Call) Return(result entities.RegisterApproval, err error) *RegisterApprovalRepositoryInterface_Create_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *RegisterApprovalRepositoryInterface_Create_Call) RunAndReturn(run func(entities.RegisterApproval) (entities.RegisterApproval, error)) *RegisterApprovalRepositoryInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteByAdminId provides a mock function with given fields: id
func (_m *RegisterApprovalRepositoryInterface) DeleteByAdminId(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RegisterApprovalRepositoryInterface_DeleteByAdminId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteByAdminId'
type RegisterApprovalRepositoryInterface_DeleteByAdminId_Call struct {
	*mock.Call
}

// DeleteByAdminId is a helper method to define mock.On call
//   - id uint
func (_e *RegisterApprovalRepositoryInterface_Expecter) DeleteByAdminId(id interface{}) *RegisterApprovalRepositoryInterface_DeleteByAdminId_Call {
	return &RegisterApprovalRepositoryInterface_DeleteByAdminId_Call{Call: _e.mock.On("DeleteByAdminId", id)}
}

func (_c *RegisterApprovalRepositoryInterface_DeleteByAdminId_Call) Run(run func(id uint)) *RegisterApprovalRepositoryInterface_DeleteByAdminId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *RegisterApprovalRepositoryInterface_DeleteByAdminId_Call) Return(err error) *RegisterApprovalRepositoryInterface_DeleteByAdminId_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *RegisterApprovalRepositoryInterface_DeleteByAdminId_Call) RunAndReturn(run func(uint) error) *RegisterApprovalRepositoryInterface_DeleteByAdminId_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllPendingApproval provides a mock function with given fields:
func (_m *RegisterApprovalRepositoryInterface) GetAllPendingApproval() ([]entities.RegisterApproval, error) {
	ret := _m.Called()

	var r0 []entities.RegisterApproval
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]entities.RegisterApproval, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []entities.RegisterApproval); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.RegisterApproval)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterApprovalRepositoryInterface_GetAllPendingApproval_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllPendingApproval'
type RegisterApprovalRepositoryInterface_GetAllPendingApproval_Call struct {
	*mock.Call
}

// GetAllPendingApproval is a helper method to define mock.On call
func (_e *RegisterApprovalRepositoryInterface_Expecter) GetAllPendingApproval() *RegisterApprovalRepositoryInterface_GetAllPendingApproval_Call {
	return &RegisterApprovalRepositoryInterface_GetAllPendingApproval_Call{Call: _e.mock.On("GetAllPendingApproval")}
}

func (_c *RegisterApprovalRepositoryInterface_GetAllPendingApproval_Call) Run(run func()) *RegisterApprovalRepositoryInterface_GetAllPendingApproval_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RegisterApprovalRepositoryInterface_GetAllPendingApproval_Call) Return(approvals []entities.RegisterApproval, err error) *RegisterApprovalRepositoryInterface_GetAllPendingApproval_Call {
	_c.Call.Return(approvals, err)
	return _c
}

func (_c *RegisterApprovalRepositoryInterface_GetAllPendingApproval_Call) RunAndReturn(run func() ([]entities.RegisterApproval, error)) *RegisterApprovalRepositoryInterface_GetAllPendingApproval_Call {
	_c.Call.Return(run)
	return _c
}

// GetByAdminID provides a mock function with given fields: id
func (_m *RegisterApprovalRepositoryInterface) GetByAdminID(id uint) (entities.RegisterApproval, error) {
	ret := _m.Called(id)

	var r0 entities.RegisterApproval
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (entities.RegisterApproval, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) entities.RegisterApproval); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(entities.RegisterApproval)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterApprovalRepositoryInterface_GetByAdminID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByAdminID'
type RegisterApprovalRepositoryInterface_GetByAdminID_Call struct {
	*mock.Call
}

// GetByAdminID is a helper method to define mock.On call
//   - id uint
func (_e *RegisterApprovalRepositoryInterface_Expecter) GetByAdminID(id interface{}) *RegisterApprovalRepositoryInterface_GetByAdminID_Call {
	return &RegisterApprovalRepositoryInterface_GetByAdminID_Call{Call: _e.mock.On("GetByAdminID", id)}
}

func (_c *RegisterApprovalRepositoryInterface_GetByAdminID_Call) Run(run func(id uint)) *RegisterApprovalRepositoryInterface_GetByAdminID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *RegisterApprovalRepositoryInterface_GetByAdminID_Call) Return(approval entities.RegisterApproval, err error) *RegisterApprovalRepositoryInterface_GetByAdminID_Call {
	_c.Call.Return(approval, err)
	return _c
}

func (_c *RegisterApprovalRepositoryInterface_GetByAdminID_Call) RunAndReturn(run func(uint) (entities.RegisterApproval, error)) *RegisterApprovalRepositoryInterface_GetByAdminID_Call {
	_c.Call.Return(run)
	return _c
}

// GetByAdminIdBatch provides a mock function with given fields: id
func (_m *RegisterApprovalRepositoryInterface) GetByAdminIdBatch(id []uint) ([]entities.RegisterApproval, error) {
	ret := _m.Called(id)

	var r0 []entities.RegisterApproval
	var r1 error
	if rf, ok := ret.Get(0).(func([]uint) ([]entities.RegisterApproval, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func([]uint) []entities.RegisterApproval); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.RegisterApproval)
		}
	}

	if rf, ok := ret.Get(1).(func([]uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterApprovalRepositoryInterface_GetByAdminIdBatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByAdminIdBatch'
type RegisterApprovalRepositoryInterface_GetByAdminIdBatch_Call struct {
	*mock.Call
}

// GetByAdminIdBatch is a helper method to define mock.On call
//   - id []uint
func (_e *RegisterApprovalRepositoryInterface_Expecter) GetByAdminIdBatch(id interface{}) *RegisterApprovalRepositoryInterface_GetByAdminIdBatch_Call {
	return &RegisterApprovalRepositoryInterface_GetByAdminIdBatch_Call{Call: _e.mock.On("GetByAdminIdBatch", id)}
}

func (_c *RegisterApprovalRepositoryInterface_GetByAdminIdBatch_Call) Run(run func(id []uint)) *RegisterApprovalRepositoryInterface_GetByAdminIdBatch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]uint))
	})
	return _c
}

func (_c *RegisterApprovalRepositoryInterface_GetByAdminIdBatch_Call) Return(approvals []entities.RegisterApproval, err error) *RegisterApprovalRepositoryInterface_GetByAdminIdBatch_Call {
	_c.Call.Return(approvals, err)
	return _c
}

func (_c *RegisterApprovalRepositoryInterface_GetByAdminIdBatch_Call) RunAndReturn(run func([]uint) ([]entities.RegisterApproval, error)) *RegisterApprovalRepositoryInterface_GetByAdminIdBatch_Call {
	_c.Call.Return(run)
	return _c
}

// InitTransaction provides a mock function with given fields:
func (_m *RegisterApprovalRepositoryInterface) InitTransaction() (db.Transactor, error) {
	ret := _m.Called()

	var r0 db.Transactor
	var r1 error
	if rf, ok := ret.Get(0).(func() (db.Transactor, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() db.Transactor); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(db.Transactor)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterApprovalRepositoryInterface_InitTransaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InitTransaction'
type RegisterApprovalRepositoryInterface_InitTransaction_Call struct {
	*mock.Call
}

// InitTransaction is a helper method to define mock.On call
func (_e *RegisterApprovalRepositoryInterface_Expecter) InitTransaction() *RegisterApprovalRepositoryInterface_InitTransaction_Call {
	return &RegisterApprovalRepositoryInterface_InitTransaction_Call{Call: _e.mock.On("InitTransaction")}
}

func (_c *RegisterApprovalRepositoryInterface_InitTransaction_Call) Run(run func()) *RegisterApprovalRepositoryInterface_InitTransaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RegisterApprovalRepositoryInterface_InitTransaction_Call) Return(_a0 db.Transactor, _a1 error) *RegisterApprovalRepositoryInterface_InitTransaction_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RegisterApprovalRepositoryInterface_InitTransaction_Call) RunAndReturn(run func() (db.Transactor, error)) *RegisterApprovalRepositoryInterface_InitTransaction_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: approval
func (_m *RegisterApprovalRepositoryInterface) Update(approval entities.RegisterApproval) error {
	ret := _m.Called(approval)

	var r0 error
	if rf, ok := ret.Get(0).(func(entities.RegisterApproval) error); ok {
		r0 = rf(approval)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RegisterApprovalRepositoryInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type RegisterApprovalRepositoryInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - approval entities.RegisterApproval
func (_e *RegisterApprovalRepositoryInterface_Expecter) Update(approval interface{}) *RegisterApprovalRepositoryInterface_Update_Call {
	return &RegisterApprovalRepositoryInterface_Update_Call{Call: _e.mock.On("Update", approval)}
}

func (_c *RegisterApprovalRepositoryInterface_Update_Call) Run(run func(approval entities.RegisterApproval)) *RegisterApprovalRepositoryInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(entities.RegisterApproval))
	})
	return _c
}

func (_c *RegisterApprovalRepositoryInterface_Update_Call) Return(err error) *RegisterApprovalRepositoryInterface_Update_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *RegisterApprovalRepositoryInterface_Update_Call) RunAndReturn(run func(entities.RegisterApproval) error) *RegisterApprovalRepositoryInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewRegisterApprovalRepositoryInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewRegisterApprovalRepositoryInterface creates a new instance of RegisterApprovalRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRegisterApprovalRepositoryInterface(t mockConstructorTestingTNewRegisterApprovalRepositoryInterface) *RegisterApprovalRepositoryInterface {
	mock := &RegisterApprovalRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}