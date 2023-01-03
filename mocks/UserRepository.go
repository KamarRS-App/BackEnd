// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	user "github.com/KamarRS-App/KamarRS-App/features/user"
	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the RepositoryInterface type
type UserRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: input
func (_m *UserRepository) Create(input user.CoreUser) error {
	ret := _m.Called(input)

	var r0 error
	if rf, ok := ret.Get(0).(func(user.CoreUser) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteById provides a mock function with given fields: id
func (_m *UserRepository) DeleteById(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetById provides a mock function with given fields: id
func (_m *UserRepository) GetById(id int) (user.CoreUser, error) {
	ret := _m.Called(id)

	var r0 user.CoreUser
	if rf, ok := ret.Get(0).(func(int) user.CoreUser); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(user.CoreUser)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, input
func (_m *UserRepository) Update(id int, input user.CoreUser) error {
	ret := _m.Called(id, input)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, user.CoreUser) error); ok {
		r0 = rf(id, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepository(t mockConstructorTestingTNewUserRepository) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
