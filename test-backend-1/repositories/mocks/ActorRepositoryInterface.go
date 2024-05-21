// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks_repositories

import (
	entities "test-backend-1/entities"
	db "test-backend-1/utils/db"

	mock "github.com/stretchr/testify/mock"

	repositories "test-backend-1/repositories"
)

// ActorRepositoryInterface is an autogenerated mock type for the ActorRepositoryInterface type
type ActorRepositoryInterface struct {
	mock.Mock
}

type ActorRepositoryInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *ActorRepositoryInterface) EXPECT() *ActorRepositoryInterface_Expecter {
	return &ActorRepositoryInterface_Expecter{mock: &_m.Mock}
}

// Begin provides a mock function with given fields: transactor
func (_m *ActorRepositoryInterface) Begin(transactor db.Transactor) repositories.ActorRepositoryInterface {
	ret := _m.Called(transactor)

	var r0 repositories.ActorRepositoryInterface
	if rf, ok := ret.Get(0).(func(db.Transactor) repositories.ActorRepositoryInterface); ok {
		r0 = rf(transactor)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(repositories.ActorRepositoryInterface)
		}
	}

	return r0
}

// ActorRepositoryInterface_Begin_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Begin'
type ActorRepositoryInterface_Begin_Call struct {
	*mock.Call
}

// Begin is a helper method to define mock.On call
//   - transactor db.Transactor
func (_e *ActorRepositoryInterface_Expecter) Begin(transactor interface{}) *ActorRepositoryInterface_Begin_Call {
	return &ActorRepositoryInterface_Begin_Call{Call: _e.mock.On("Begin", transactor)}
}

func (_c *ActorRepositoryInterface_Begin_Call) Run(run func(transactor db.Transactor)) *ActorRepositoryInterface_Begin_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(db.Transactor))
	})
	return _c
}

func (_c *ActorRepositoryInterface_Begin_Call) Return(_a0 repositories.ActorRepositoryInterface) *ActorRepositoryInterface_Begin_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ActorRepositoryInterface_Begin_Call) RunAndReturn(run func(db.Transactor) repositories.ActorRepositoryInterface) *ActorRepositoryInterface_Begin_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: actor
func (_m *ActorRepositoryInterface) Create(actor entities.Actor) (entities.Actor, error) {
	ret := _m.Called(actor)

	var r0 entities.Actor
	var r1 error
	if rf, ok := ret.Get(0).(func(entities.Actor) (entities.Actor, error)); ok {
		return rf(actor)
	}
	if rf, ok := ret.Get(0).(func(entities.Actor) entities.Actor); ok {
		r0 = rf(actor)
	} else {
		r0 = ret.Get(0).(entities.Actor)
	}

	if rf, ok := ret.Get(1).(func(entities.Actor) error); ok {
		r1 = rf(actor)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ActorRepositoryInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type ActorRepositoryInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - actor entities.Actor
func (_e *ActorRepositoryInterface_Expecter) Create(actor interface{}) *ActorRepositoryInterface_Create_Call {
	return &ActorRepositoryInterface_Create_Call{Call: _e.mock.On("Create", actor)}
}

func (_c *ActorRepositoryInterface_Create_Call) Run(run func(actor entities.Actor)) *ActorRepositoryInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(entities.Actor))
	})
	return _c
}

func (_c *ActorRepositoryInterface_Create_Call) Return(result entities.Actor, err error) *ActorRepositoryInterface_Create_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *ActorRepositoryInterface_Create_Call) RunAndReturn(run func(entities.Actor) (entities.Actor, error)) *ActorRepositoryInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: id
func (_m *ActorRepositoryInterface) Delete(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ActorRepositoryInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type ActorRepositoryInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - id uint
func (_e *ActorRepositoryInterface_Expecter) Delete(id interface{}) *ActorRepositoryInterface_Delete_Call {
	return &ActorRepositoryInterface_Delete_Call{Call: _e.mock.On("Delete", id)}
}

func (_c *ActorRepositoryInterface_Delete_Call) Run(run func(id uint)) *ActorRepositoryInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *ActorRepositoryInterface_Delete_Call) Return(err error) *ActorRepositoryInterface_Delete_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *ActorRepositoryInterface_Delete_Call) RunAndReturn(run func(uint) error) *ActorRepositoryInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllByUsername provides a mock function with given fields: username, limit, offset
func (_m *ActorRepositoryInterface) GetAllByUsername(username string, limit uint, offset uint) ([]entities.Actor, error) {
	ret := _m.Called(username, limit, offset)

	var r0 []entities.Actor
	var r1 error
	if rf, ok := ret.Get(0).(func(string, uint, uint) ([]entities.Actor, error)); ok {
		return rf(username, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(string, uint, uint) []entities.Actor); ok {
		r0 = rf(username, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.Actor)
		}
	}

	if rf, ok := ret.Get(1).(func(string, uint, uint) error); ok {
		r1 = rf(username, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ActorRepositoryInterface_GetAllByUsername_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllByUsername'
type ActorRepositoryInterface_GetAllByUsername_Call struct {
	*mock.Call
}

// GetAllByUsername is a helper method to define mock.On call
//   - username string
//   - limit uint
//   - offset uint
func (_e *ActorRepositoryInterface_Expecter) GetAllByUsername(username interface{}, limit interface{}, offset interface{}) *ActorRepositoryInterface_GetAllByUsername_Call {
	return &ActorRepositoryInterface_GetAllByUsername_Call{Call: _e.mock.On("GetAllByUsername", username, limit, offset)}
}

func (_c *ActorRepositoryInterface_GetAllByUsername_Call) Run(run func(username string, limit uint, offset uint)) *ActorRepositoryInterface_GetAllByUsername_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(uint), args[2].(uint))
	})
	return _c
}

func (_c *ActorRepositoryInterface_GetAllByUsername_Call) Return(actor []entities.Actor, err error) *ActorRepositoryInterface_GetAllByUsername_Call {
	_c.Call.Return(actor, err)
	return _c
}

func (_c *ActorRepositoryInterface_GetAllByUsername_Call) RunAndReturn(run func(string, uint, uint) ([]entities.Actor, error)) *ActorRepositoryInterface_GetAllByUsername_Call {
	_c.Call.Return(run)
	return _c
}

// GetByID provides a mock function with given fields: id
func (_m *ActorRepositoryInterface) GetByID(id uint) (entities.Actor, error) {
	ret := _m.Called(id)

	var r0 entities.Actor
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (entities.Actor, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) entities.Actor); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(entities.Actor)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ActorRepositoryInterface_GetByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByID'
type ActorRepositoryInterface_GetByID_Call struct {
	*mock.Call
}

// GetByID is a helper method to define mock.On call
//   - id uint
func (_e *ActorRepositoryInterface_Expecter) GetByID(id interface{}) *ActorRepositoryInterface_GetByID_Call {
	return &ActorRepositoryInterface_GetByID_Call{Call: _e.mock.On("GetByID", id)}
}

func (_c *ActorRepositoryInterface_GetByID_Call) Run(run func(id uint)) *ActorRepositoryInterface_GetByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *ActorRepositoryInterface_GetByID_Call) Return(actor entities.Actor, err error) *ActorRepositoryInterface_GetByID_Call {
	_c.Call.Return(actor, err)
	return _c
}

func (_c *ActorRepositoryInterface_GetByID_Call) RunAndReturn(run func(uint) (entities.Actor, error)) *ActorRepositoryInterface_GetByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetByUsername provides a mock function with given fields: username
func (_m *ActorRepositoryInterface) GetByUsername(username string) (entities.Actor, error) {
	ret := _m.Called(username)

	var r0 entities.Actor
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (entities.Actor, error)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(string) entities.Actor); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(entities.Actor)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ActorRepositoryInterface_GetByUsername_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByUsername'
type ActorRepositoryInterface_GetByUsername_Call struct {
	*mock.Call
}

// GetByUsername is a helper method to define mock.On call
//   - username string
func (_e *ActorRepositoryInterface_Expecter) GetByUsername(username interface{}) *ActorRepositoryInterface_GetByUsername_Call {
	return &ActorRepositoryInterface_GetByUsername_Call{Call: _e.mock.On("GetByUsername", username)}
}

func (_c *ActorRepositoryInterface_GetByUsername_Call) Run(run func(username string)) *ActorRepositoryInterface_GetByUsername_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *ActorRepositoryInterface_GetByUsername_Call) Return(actor entities.Actor, err error) *ActorRepositoryInterface_GetByUsername_Call {
	_c.Call.Return(actor, err)
	return _c
}

func (_c *ActorRepositoryInterface_GetByUsername_Call) RunAndReturn(run func(string) (entities.Actor, error)) *ActorRepositoryInterface_GetByUsername_Call {
	_c.Call.Return(run)
	return _c
}

// GetByUsernameBatch provides a mock function with given fields: username
func (_m *ActorRepositoryInterface) GetByUsernameBatch(username []string) ([]entities.Actor, error) {
	ret := _m.Called(username)

	var r0 []entities.Actor
	var r1 error
	if rf, ok := ret.Get(0).(func([]string) ([]entities.Actor, error)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func([]string) []entities.Actor); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.Actor)
		}
	}

	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ActorRepositoryInterface_GetByUsernameBatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByUsernameBatch'
type ActorRepositoryInterface_GetByUsernameBatch_Call struct {
	*mock.Call
}

// GetByUsernameBatch is a helper method to define mock.On call
//   - username []string
func (_e *ActorRepositoryInterface_Expecter) GetByUsernameBatch(username interface{}) *ActorRepositoryInterface_GetByUsernameBatch_Call {
	return &ActorRepositoryInterface_GetByUsernameBatch_Call{Call: _e.mock.On("GetByUsernameBatch", username)}
}

func (_c *ActorRepositoryInterface_GetByUsernameBatch_Call) Run(run func(username []string)) *ActorRepositoryInterface_GetByUsernameBatch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string))
	})
	return _c
}

func (_c *ActorRepositoryInterface_GetByUsernameBatch_Call) Return(actors []entities.Actor, err error) *ActorRepositoryInterface_GetByUsernameBatch_Call {
	_c.Call.Return(actors, err)
	return _c
}

func (_c *ActorRepositoryInterface_GetByUsernameBatch_Call) RunAndReturn(run func([]string) ([]entities.Actor, error)) *ActorRepositoryInterface_GetByUsernameBatch_Call {
	_c.Call.Return(run)
	return _c
}

// InitTransaction provides a mock function with given fields:
func (_m *ActorRepositoryInterface) InitTransaction() (db.Transactor, error) {
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

// ActorRepositoryInterface_InitTransaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InitTransaction'
type ActorRepositoryInterface_InitTransaction_Call struct {
	*mock.Call
}

// InitTransaction is a helper method to define mock.On call
func (_e *ActorRepositoryInterface_Expecter) InitTransaction() *ActorRepositoryInterface_InitTransaction_Call {
	return &ActorRepositoryInterface_InitTransaction_Call{Call: _e.mock.On("InitTransaction")}
}

func (_c *ActorRepositoryInterface_InitTransaction_Call) Run(run func()) *ActorRepositoryInterface_InitTransaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ActorRepositoryInterface_InitTransaction_Call) Return(_a0 db.Transactor, _a1 error) *ActorRepositoryInterface_InitTransaction_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ActorRepositoryInterface_InitTransaction_Call) RunAndReturn(run func() (db.Transactor, error)) *ActorRepositoryInterface_InitTransaction_Call {
	_c.Call.Return(run)
	return _c
}

// IsExist provides a mock function with given fields: id
func (_m *ActorRepositoryInterface) IsExist(id uint) (bool, error) {
	ret := _m.Called(id)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (bool, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ActorRepositoryInterface_IsExist_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsExist'
type ActorRepositoryInterface_IsExist_Call struct {
	*mock.Call
}

// IsExist is a helper method to define mock.On call
//   - id uint
func (_e *ActorRepositoryInterface_Expecter) IsExist(id interface{}) *ActorRepositoryInterface_IsExist_Call {
	return &ActorRepositoryInterface_IsExist_Call{Call: _e.mock.On("IsExist", id)}
}

func (_c *ActorRepositoryInterface_IsExist_Call) Run(run func(id uint)) *ActorRepositoryInterface_IsExist_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *ActorRepositoryInterface_IsExist_Call) Return(exist bool, err error) *ActorRepositoryInterface_IsExist_Call {
	_c.Call.Return(exist, err)
	return _c
}

func (_c *ActorRepositoryInterface_IsExist_Call) RunAndReturn(run func(uint) (bool, error)) *ActorRepositoryInterface_IsExist_Call {
	_c.Call.Return(run)
	return _c
}

// IsUsernameExist provides a mock function with given fields: actor
func (_m *ActorRepositoryInterface) IsUsernameExist(actor entities.Actor) (bool, error) {
	ret := _m.Called(actor)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(entities.Actor) (bool, error)); ok {
		return rf(actor)
	}
	if rf, ok := ret.Get(0).(func(entities.Actor) bool); ok {
		r0 = rf(actor)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(entities.Actor) error); ok {
		r1 = rf(actor)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ActorRepositoryInterface_IsUsernameExist_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsUsernameExist'
type ActorRepositoryInterface_IsUsernameExist_Call struct {
	*mock.Call
}

// IsUsernameExist is a helper method to define mock.On call
//   - actor entities.Actor
func (_e *ActorRepositoryInterface_Expecter) IsUsernameExist(actor interface{}) *ActorRepositoryInterface_IsUsernameExist_Call {
	return &ActorRepositoryInterface_IsUsernameExist_Call{Call: _e.mock.On("IsUsernameExist", actor)}
}

func (_c *ActorRepositoryInterface_IsUsernameExist_Call) Run(run func(actor entities.Actor)) *ActorRepositoryInterface_IsUsernameExist_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(entities.Actor))
	})
	return _c
}

func (_c *ActorRepositoryInterface_IsUsernameExist_Call) Return(exist bool, err error) *ActorRepositoryInterface_IsUsernameExist_Call {
	_c.Call.Return(exist, err)
	return _c
}

func (_c *ActorRepositoryInterface_IsUsernameExist_Call) RunAndReturn(run func(entities.Actor) (bool, error)) *ActorRepositoryInterface_IsUsernameExist_Call {
	_c.Call.Return(run)
	return _c
}

// Save provides a mock function with given fields: actor
func (_m *ActorRepositoryInterface) Save(actor entities.Actor) error {
	ret := _m.Called(actor)

	var r0 error
	if rf, ok := ret.Get(0).(func(entities.Actor) error); ok {
		r0 = rf(actor)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ActorRepositoryInterface_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type ActorRepositoryInterface_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - actor entities.Actor
func (_e *ActorRepositoryInterface_Expecter) Save(actor interface{}) *ActorRepositoryInterface_Save_Call {
	return &ActorRepositoryInterface_Save_Call{Call: _e.mock.On("Save", actor)}
}

func (_c *ActorRepositoryInterface_Save_Call) Run(run func(actor entities.Actor)) *ActorRepositoryInterface_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(entities.Actor))
	})
	return _c
}

func (_c *ActorRepositoryInterface_Save_Call) Return(err error) *ActorRepositoryInterface_Save_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *ActorRepositoryInterface_Save_Call) RunAndReturn(run func(entities.Actor) error) *ActorRepositoryInterface_Save_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: actor
func (_m *ActorRepositoryInterface) Update(actor entities.Actor) error {
	ret := _m.Called(actor)

	var r0 error
	if rf, ok := ret.Get(0).(func(entities.Actor) error); ok {
		r0 = rf(actor)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ActorRepositoryInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type ActorRepositoryInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - actor entities.Actor
func (_e *ActorRepositoryInterface_Expecter) Update(actor interface{}) *ActorRepositoryInterface_Update_Call {
	return &ActorRepositoryInterface_Update_Call{Call: _e.mock.On("Update", actor)}
}

func (_c *ActorRepositoryInterface_Update_Call) Run(run func(actor entities.Actor)) *ActorRepositoryInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(entities.Actor))
	})
	return _c
}

func (_c *ActorRepositoryInterface_Update_Call) Return(err error) *ActorRepositoryInterface_Update_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *ActorRepositoryInterface_Update_Call) RunAndReturn(run func(entities.Actor) error) *ActorRepositoryInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewActorRepositoryInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewActorRepositoryInterface creates a new instance of ActorRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewActorRepositoryInterface(t mockConstructorTestingTNewActorRepositoryInterface) *ActorRepositoryInterface {
	mock := &ActorRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
