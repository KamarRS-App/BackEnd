package service

import (
	"errors"
	"testing"

	"github.com/KamarRS-App/KamarRS-App/features/doctor"
	"github.com/KamarRS-App/KamarRS-App/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	repo := new(mocks.DoctorRepository)
	t.Run("success create", func(t *testing.T) {
		inputData := doctor.DoctorCore{Nama: "sheena", Email: "sheena@kronokedow.jp", Spesialis: "feeling doctor", NoTelpon: "123456789"}
		repo.On("Create", mock.Anything).Return(1, nil).Once()
		srv := New(repo)
		err := srv.Create(inputData)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed create", func(t *testing.T) {
		inputData := doctor.DoctorCore{Nama: "sheena", Email: "sheena@kronokedow.jp", Spesialis: "feeling doctor", NoTelpon: "123456789"}
		repo.On("Create", mock.Anything).Return(0, errors.New("failed to insert data, error logic")).Once()
		srv := New(repo)
		err := srv.Create(inputData)
		assert.EqualError(t, err, "failed to insert data, error logic")
		repo.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	repo := new(mocks.DoctorRepository)
	t.Run("success get by id", func(t *testing.T) {
		inputId := 1
		returnData := doctor.DoctorCore{ID: 1, Nama: "sheena", Email: "sheena@kronokedow.jp", Spesialis: "feeling doctor", NoTelpon: "123456789"}
		repo.On("GetById", mock.Anything).Return(returnData, nil).Once()
		srv := New(repo)
		res, err := srv.GetById(inputId)
		assert.Nil(t, err)
		assert.Equal(t, res.ID, returnData.ID)
		repo.AssertExpectations(t)
	})

	t.Run("failed get by id", func(t *testing.T) {
		inputId := 1
		repo.On("GetById", mock.Anything).Return(doctor.DoctorCore{}, errors.New("failed get doctor by id data, error logic")).Once()
		srv := New(repo)
		res, err := srv.GetById(inputId)
		assert.NotNil(t, err)
		assert.Equal(t, res, doctor.DoctorCore{})
		repo.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	repo := new(mocks.DoctorRepository)
	t.Run("success get by id", func(t *testing.T) {
		returnData := []doctor.DoctorCore{{ID: 1, Nama: "sheena", Email: "sheena@kronokedow.jp", Spesialis: "feeling doctor", NoTelpon: "123456789"}}
		repo.On("GetAll", mock.Anything).Return(returnData, nil).Once()
		srv := New(repo)
		res, err := srv.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, res[0].ID, returnData[0].ID)
		repo.AssertExpectations(t)
	})

	t.Run("failed get all", func(t *testing.T) {
		repo.On("GetAll", mock.Anything).Return(nil, errors.New("failed get all doctors")).Once()
		srv := New(repo)
		res, err := srv.GetAll()
		assert.NotNil(t, err)
		assert.Nil(t, res)
		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	repo := new(mocks.DoctorRepository)
	t.Run("success update", func(t *testing.T) {
		inputData := doctor.DoctorCore{ID: 1, Nama: "sheena", Email: "sheena@kronokedow.jp", Spesialis: "feeling doctor", NoTelpon: "123456789"}
		inputId := 1
		repo.On("Update", inputData, inputId).Return(nil).Once()
		srv := New(repo)
		err := srv.Update(inputData, inputId)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed update", func(t *testing.T) {
		inputData := doctor.DoctorCore{Nama: "sheena", Email: "sheena@kronokedow.jp", Spesialis: "feeling doctor", NoTelpon: "123456789"}
		inputId := 1
		repo.On("Update", inputData, inputId).Return(errors.New("failed update doctor data, error logic")).Once()
		srv := New(repo)
		err := srv.Update(inputData, inputId)
		assert.EqualError(t, err, "failed update doctor data, error logic")
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := new(mocks.DoctorRepository)
	t.Run("success delete", func(t *testing.T) {
		inputId := 1
		repo.On("Delete", mock.Anything).Return(1, nil).Once()
		srv := New(repo)
		err := srv.Delete(inputId)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed delete", func(t *testing.T) {
		inputId := 1
		repo.On("Delete", mock.Anything).Return(0, errors.New("failed delete doctor, error logic")).Once()
		srv := New(repo)
		err := srv.Delete(inputId)
		assert.EqualError(t, err, "failed delete doctor, error logic")
		repo.AssertExpectations(t)
	})
}
