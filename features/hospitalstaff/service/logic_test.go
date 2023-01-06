package service

import (
	"errors"
	"testing"

	"github.com/KamarRS-App/KamarRS-App/features/hospitalstaff"
	"github.com/KamarRS-App/KamarRS-App/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	repo := new(mocks.StaffRepository)
	t.Run("Succes Create User", func(t *testing.T) {
		inputRepo := hospitalstaff.HospitalStaffCore{Nama: "teguh", Email: "teguh@mail.id", KataSandi: "qwerty", Peran: "admin", HospitalID: 1, HospitalName: "RSJIWA"}
		inputData := hospitalstaff.HospitalStaffCore{Nama: "teguh", Email: "teguh@mail.id", KataSandi: "qwerty", Peran: "admin", HospitalID: 1, HospitalName: "RSJIWA"}

		repo.On("Create", inputRepo).Return(nil).Once()
		srv := New(repo)
		err := srv.Create(inputData)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Create staff", func(t *testing.T) {
		inputData := hospitalstaff.HospitalStaffCore{Nama: "teguh", KataSandi: "qwerty", Peran: "admin", HospitalID: 1, HospitalName: "RSJIWA"}
		srv := New(repo)
		err := srv.Create(inputData)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Create, Duplicate", func(t *testing.T) {
		inputRepo := hospitalstaff.HospitalStaffCore{Nama: "teguh", Email: "teguh@mail.id", KataSandi: "qwerty", Peran: "admin", HospitalID: 1, HospitalName: "RSJIWA"}
		inputData := hospitalstaff.HospitalStaffCore{Nama: "teguh", Email: "teguh@mail.id", KataSandi: "qwerty", Peran: "admin", HospitalID: 1, HospitalName: "RSJIWA"}
		repo.On("Create", inputRepo).Return(errors.New(" Gagal membuat akun, intput data salah atau Email sudah terdaftar")).Once()
		srv := New(repo)
		err := srv.Create(inputData)
		assert.NotNil(t, err)
		assert.Equal(t, " Gagal membuat akun staff", err.Error()) // samakan dengan di logic
		repo.AssertExpectations(t)
	})

}

func TestUpdate(t *testing.T) {
	repo := new(mocks.StaffRepository)
	t.Run("Success Update", func(t *testing.T) {
		inputRepo := hospitalstaff.HospitalStaffCore{ID: 1, Nama: "teguh", Email: "teguh@mail.id", KataSandi: "qwerty", Peran: "admin", HospitalID: 1, HospitalName: "RSJIWA"}
		inputData := hospitalstaff.HospitalStaffCore{ID: 1, Nama: "teguh", Email: "teguh@mail.id", KataSandi: "qwerty", Peran: "admin", HospitalID: 1, HospitalName: "RSJIWA"}
		repo.On("Update", int(inputRepo.ID), inputRepo).Return(nil).Once()
		srv := New(repo)
		err := srv.Update(int(inputData.ID), inputData)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Update data", func(t *testing.T) {
		inputRepo := hospitalstaff.HospitalStaffCore{ID: 1, Nama: "teguh", Email: "teguh@mail.id", KataSandi: "qwerty", Peran: "admin", HospitalID: 1, HospitalName: "RSJIWA"}
		inputData := hospitalstaff.HospitalStaffCore{ID: 1, Nama: "teguh", Email: "teguh@mail.id", KataSandi: "qwerty", Peran: "admin", HospitalID: 1, HospitalName: "RSJIWA"}
		repo.On("Update", int(inputRepo.ID), inputRepo).Return(errors.New("gagal mengupdate data , querry error")).Once()
		srv := New(repo)
		err := srv.Update(int(inputData.ID), inputData)
		assert.NotNil(t, err)
		assert.Equal(t, "gagal mengupdate data , querry error", err.Error())

		repo.AssertExpectations(t)
	})

}

func TestDelete(t *testing.T) {
	repo := new(mocks.StaffRepository)
	t.Run("Success Delete", func(t *testing.T) {
		inputRepo := hospitalstaff.HospitalStaffCore{ID: 1}
		inputData := hospitalstaff.HospitalStaffCore{ID: 1}
		repo.On("DeleteById", int(inputRepo.ID)).Return(nil).Once()
		srv := New(repo)
		err := srv.DeleteById(int(inputData.ID))
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Delete data", func(t *testing.T) {
		inputRepo := hospitalstaff.HospitalStaffCore{ID: 1}
		inputData := hospitalstaff.HospitalStaffCore{ID: 1}
		repo.On("DeleteById", int(inputRepo.ID)).Return(errors.New("Gagal Delete")).Once()
		srv := New(repo)
		err := srv.DeleteById(int(inputData.ID))
		assert.NotNil(t, err)
		assert.Equal(t, "gagal menghapus data , querry error", err.Error())
		repo.AssertExpectations(t)
	})

}

func TestGetAllStaff(t *testing.T) {
	repo := new(mocks.StaffRepository)
	returnData := []hospitalstaff.HospitalStaffCore{{ID: 1, Nama: "teguh", Email: "teguh@mail.id", KataSandi: "qwerty", Peran: "admin", HospitalID: 1, HospitalName: "RSJIWA"}}
	TotalPage := 1

	t.Run("Succes Get user by id", func(t *testing.T) {
		repo.On("GetAllStaff", "RSJIWA", 10, 0).Return(returnData, TotalPage, nil).Once() //menentukan fungsi yang akan dijalankan//

		srv := New(repo) //membuat dependency injection
		data, totalPage, err := srv.GetAllStaff("RSJIWA", 10, 1)
		assert.Nil(t, err)
		assert.Equal(t, totalPage, 1)
		assert.Equal(t, data[0].Nama, returnData[0].Nama)
		repo.AssertExpectations(t)
	})

	t.Run("Failed to Get user by id", func(t *testing.T) {
		repo.On("GetAllStaff", "RSJIWA", 10, 0).Return(nil, 0, errors.New("Gagal menampilkan data")).Once() //menentukan fungsi yang akan dijalankan//

		srv := New(repo)
		data, totalPage, err := srv.GetAllStaff("RSJIWA", 10, 1)
		assert.NotNil(t, err)
		assert.NotEqual(t, totalPage, 1)
		assert.Nil(t, data)
		assert.Equal(t, "failed get staff by hospital id, error logic", err.Error())
		repo.AssertExpectations(t)
	})

}

func TestGetstaffbyid(t *testing.T) {
	repo := new(mocks.StaffRepository)
	returnData := hospitalstaff.HospitalStaffCore{ID: 1, Nama: "teguh", Email: "teguh@mail.id", KataSandi: "qwerty", Peran: "admin", HospitalID: 1, HospitalName: "RSJIWA"}

	staffID := 1

	t.Run("Succes Get staff by id", func(t *testing.T) {
		repo.On("GetStaff", staffID).Return(returnData, nil).Once() //menentukan fungsi yang akan dijalankan//

		srv := New(repo) //membuat dependency injection
		data, err := srv.GetStaff(staffID)
		assert.Nil(t, err)
		assert.Equal(t, data.Nama, returnData.Nama)
		repo.AssertExpectations(t)
	})

	t.Run("Failed to Get staff by id", func(t *testing.T) {
		repo.On("GetStaff", staffID).Return(hospitalstaff.HospitalStaffCore{}, errors.New("Gagal menampilkan data")).Once() //menentukan fungsi yang akan dijalankan//
		srv := New(repo)
		data, err := srv.GetStaff(staffID)
		assert.NotNil(t, err)
		assert.NotEqual(t, data.Nama, returnData.Nama)
		assert.Equal(t, "gagal menampilkan data , querry error", err.Error())
		repo.AssertExpectations(t)
	})

}
