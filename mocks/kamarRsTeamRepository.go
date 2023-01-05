// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	kamarrsteam "github.com/KamarRS-App/KamarRS-App/features/kamarrsteam"
	mock "github.com/stretchr/testify/mock"
)

// KamarRsTeamRepository is an autogenerated mock type for the RepositoryInterface type
type KamarRsTeamRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: input
func (_m *KamarRsTeamRepository) Create(input kamarrsteam.KamarRsTeamCore) error {
	ret := _m.Called(input)

	var r0 error
	if rf, ok := ret.Get(0).(func(kamarrsteam.KamarRsTeamCore) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindTeam provides a mock function with given fields: email
func (_m *KamarRsTeamRepository) FindTeam(email string) (kamarrsteam.KamarRsTeamCore, error) {
	ret := _m.Called(email)

	var r0 kamarrsteam.KamarRsTeamCore
	if rf, ok := ret.Get(0).(func(string) kamarrsteam.KamarRsTeamCore); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(kamarrsteam.KamarRsTeamCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewKamarRsTeamRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewKamarRsTeamRepository creates a new instance of KamarRsTeamRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewKamarRsTeamRepository(t mockConstructorTestingTNewKamarRsTeamRepository) *KamarRsTeamRepository {
	mock := &KamarRsTeamRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}