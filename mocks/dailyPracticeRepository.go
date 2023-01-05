// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	dailypractice "github.com/KamarRS-App/KamarRS-App/features/dailyPractice"
	mock "github.com/stretchr/testify/mock"
)

// DailyPracticeRepository is an autogenerated mock type for the RepositoryInterface type
type DailyPracticeRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: input
func (_m *DailyPracticeRepository) Create(input dailypractice.PracticeCore) (int, error) {
	ret := _m.Called(input)

	var r0 int
	if rf, ok := ret.Get(0).(func(dailypractice.PracticeCore) int); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(dailypractice.PracticeCore) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: limit, offset, id
func (_m *DailyPracticeRepository) GetAll(limit int, offset int, id int) ([]dailypractice.PracticeCore, int, error) {
	ret := _m.Called(limit, offset, id)

	var r0 []dailypractice.PracticeCore
	if rf, ok := ret.Get(0).(func(int, int, int) []dailypractice.PracticeCore); ok {
		r0 = rf(limit, offset, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dailypractice.PracticeCore)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(int, int, int) int); ok {
		r1 = rf(limit, offset, id)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int, int, int) error); ok {
		r2 = rf(limit, offset, id)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetById provides a mock function with given fields: id
func (_m *DailyPracticeRepository) GetById(id int) (dailypractice.PracticeCore, error) {
	ret := _m.Called(id)

	var r0 dailypractice.PracticeCore
	if rf, ok := ret.Get(0).(func(int) dailypractice.PracticeCore); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(dailypractice.PracticeCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: input, id
func (_m *DailyPracticeRepository) Update(input dailypractice.PracticeCore, id int) error {
	ret := _m.Called(input, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(dailypractice.PracticeCore, int) error); ok {
		r0 = rf(input, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewDailyPracticeRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewDailyPracticeRepository creates a new instance of DailyPracticeRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDailyPracticeRepository(t mockConstructorTestingTNewDailyPracticeRepository) *DailyPracticeRepository {
	mock := &DailyPracticeRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}