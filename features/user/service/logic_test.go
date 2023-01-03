package service

import (
	"errors"
	"testing"

	"github.com/KamarRS-App/KamarRS-App/features/user"
	"github.com/KamarRS-App/KamarRS-App/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	repo := new(mocks.UserRepository)
	t.Run("Succes Create User", func(t *testing.T) {
		inputRepo := user.CoreUser{Nama: "teguh", Email: "teguh@mail.id", KataSandi: "qwerty", Nokk: "123", Nik: "321", NoTelpon: "08123"}
		inputData := user.CoreUser{Nama: "teguh", Email: "teguh@mail.id", KataSandi: "qwerty", Nokk: "123", Nik: "321", NoTelpon: "08123"}

		repo.On("Create", inputRepo).Return(nil).Once()
		srv := New(repo)
		err := srv.Create(inputData)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Create, Empty Name", func(t *testing.T) {
		inputData := user.CoreUser{Nama: "teguh", Nokk: "123", Nik: "321", NoTelpon: "08123"}
		srv := New(repo)
		err := srv.Create(inputData)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Create, Duplicate", func(t *testing.T) {
		inputRepo := user.CoreUser{Nama: "teguh", Email: "teguh@mail.id", Nokk: "123", Nik: "321", NoTelpon: "08123"}
		inputData := user.CoreUser{Nama: "teguh", Email: "teguh@mail.id", Nokk: "123", Nik: "321", NoTelpon: "08123"}
		repo.On("Create", inputRepo).Return(errors.New(" Gagal membuat akun, intput data salah atau Email sudah terdaftar")).Once()
		srv := New(repo)
		err := srv.Create(inputData)
		assert.NotNil(t, err)
		assert.Equal(t, " Gagal membuat akun, intput data salah atau Email sudah terdaftar", err.Error()) // samakan dengan di logic
		repo.AssertExpectations(t)
	})

}

func TestUpdate(t *testing.T) {
	repo := new(mocks.UserRepository)
	t.Run("Success Update", func(t *testing.T) {
		inputRepo := user.CoreUser{ID: 1, Nama: "teguh", Email: "teguh@mail.id", KataSandi: "qwerty", Nokk: "123", Nik: "321", NoTelpon: "08123"}
		inputData := user.CoreUser{ID: 1, Nama: "teguh", Email: "teguh@mail.id", KataSandi: "qwerty", Nokk: "123", Nik: "321", NoTelpon: "08123"}
		repo.On("Update", int(inputRepo.ID), inputRepo).Return(nil).Once()
		srv := New(repo)
		err := srv.Update(int(inputData.ID), inputData)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Update data", func(t *testing.T) {
		inputRepo := user.CoreUser{ID: 1, Nama: "teguh", Email: "teguh@mail.id", KataSandi: "qwerty", Nokk: "123", Nik: "321", NoTelpon: "08123"}
		inputData := user.CoreUser{ID: 1, Nama: "teguh", Email: "teguh@mail.id", KataSandi: "qwerty", Nokk: "123", Nik: "321", NoTelpon: "08123"}
		repo.On("Update", int(inputRepo.ID), inputRepo).Return(errors.New("gagal mengupdate data , querry error")).Once()
		srv := New(repo)
		err := srv.Update(int(inputData.ID), inputData)
		assert.NotNil(t, err)
		assert.Equal(t, "gagal mengupdate data , querry error", err.Error())

		repo.AssertExpectations(t)
	})

}

func TestDelete(t *testing.T) {
	repo := new(mocks.UserRepository)
	t.Run("Success Delete", func(t *testing.T) {
		inputRepo := user.CoreUser{ID: 1}
		inputData := user.CoreUser{ID: 1}
		repo.On("DeleteById", int(inputRepo.ID)).Return(nil).Once()
		srv := New(repo)
		err := srv.DeleteById(int(inputData.ID))
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Delete data", func(t *testing.T) {
		inputRepo := user.CoreUser{ID: 1}
		inputData := user.CoreUser{ID: 1}
		repo.On("DeleteById", int(inputRepo.ID)).Return(errors.New("Id not Found")).Once()
		srv := New(repo)
		err := srv.DeleteById(int(inputData.ID))
		assert.NotNil(t, err)

		repo.AssertExpectations(t)
	})

}

func TestGetbyId(t *testing.T) {
	repo := new(mocks.UserRepository)
	returnData := user.CoreUser{ID: 1, Nama: "teguh", Email: "teguh@mail.id", KataSandi: "qwerty", Nokk: "123", Nik: "321", NoTelpon: "08123"}
	t.Run("Succes Get user by id", func(t *testing.T) {
		repo.On("GetById", 1).Return(returnData, nil).Once() //menentukan fungsi yang akan dijalankan//

		srv := New(repo) //membuat dependency injection
		response, err := srv.GetById(1)
		assert.Nil(t, err)
		assert.Equal(t, returnData.Nama, response.Nama)
		repo.AssertExpectations(t)
	})

	t.Run("Failed to Get user by id", func(t *testing.T) {
		repo.On("GetById", 1).Return(user.CoreUser{}, errors.New("Id not Found")).Once() //menentukan fungsi yang akan dijalankan//

		srv := New(repo) //membuat dependency injection
		response, err := srv.GetById(1)
		assert.NotNil(t, err)
		assert.NotEqual(t, returnData.Nama, response.Nama)
		assert.Equal(t, "gagal mengupdate data , querry error", err.Error())
		repo.AssertExpectations(t)
	})

}
