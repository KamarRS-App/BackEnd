package service

import (
	"errors"
	"testing"

	"github.com/KamarRS-App/KamarRS-App/features/bed"
	"github.com/KamarRS-App/KamarRS-App/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	repo := new(mocks.BedRepository)
	t.Run("Succes Create User", func(t *testing.T) {
		inputRepo := bed.BedCore{NamaTempatTidur: "bunga", Ruangan: "A.17", Kelas: "umum", Status: "available", HospitalID: 1}
		inputData := bed.BedCore{NamaTempatTidur: "bunga", Ruangan: "A.17", Kelas: "umum", Status: "available", HospitalID: 1}

		repo.On("Create", inputRepo).Return(1, nil).Once()
		srv := New(repo)
		err := srv.Create(inputData)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Create, Empty Name", func(t *testing.T) {
		inputData := bed.BedCore{ID: 1, Ruangan: "A.17", Kelas: "umum", Status: "available", HospitalID: 1}
		srv := New(repo)
		err := srv.Create(inputData)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Create, Duplicate", func(t *testing.T) {
		inputRepo := bed.BedCore{NamaTempatTidur: "bunga", Ruangan: "A.17", Kelas: "umum", Status: "available", HospitalID: 1}
		inputData := bed.BedCore{NamaTempatTidur: "bunga", Ruangan: "A.17", Kelas: "umum", Status: "available", HospitalID: 1}
		repo.On("Create", inputRepo).Return(0, errors.New(" Gagal membuat akun, intput data salah atau Email sudah terdaftar")).Once()
		srv := New(repo)
		err := srv.Create(inputData)
		assert.NotNil(t, err)
		assert.Equal(t, "failed to insert data, error logic", err.Error()) // samakan dengan di logic
		repo.AssertExpectations(t)
	})

}

func TestGetAllBeds(t *testing.T) {
	repo := new(mocks.BedRepository)
	returnData := []bed.BedCore{{ID: 1, Ruangan: "A.17", Kelas: "umum", Status: "available", HospitalID: 1}}
	TotalPage := 1
	kelasreq := "umum"
	statusreq := "available"
	t.Run("Succes Get user by id", func(t *testing.T) {
		repo.On("GetAll", "umum", "available", 10, 0, 1).Return(returnData, TotalPage, nil).Once() //menentukan fungsi yang akan dijalankan//

		srv := New(repo) //membuat dependency injection
		data, totalPage, err := srv.GetAll(kelasreq, statusreq, 1, 10, 1)
		assert.Nil(t, err)
		assert.Equal(t, totalPage, 1)
		assert.Equal(t, data[0].Kelas, returnData[0].Kelas)
		repo.AssertExpectations(t)
	})

	t.Run("Failed to Get user by id", func(t *testing.T) {
		repo.On("GetAll", "umum", "available", 10, 0, 1).Return(nil, 0, errors.New("Gagal menampilkan data")).Once() //menentukan fungsi yang akan dijalankan//

		srv := New(repo)
		data, totalPage, err := srv.GetAll(kelasreq, statusreq, 1, 10, 1)
		assert.NotNil(t, err)
		assert.NotEqual(t, totalPage, 1)
		assert.Nil(t, data)
		assert.Equal(t, "failed get bed by hospital id, error logic", err.Error())
		repo.AssertExpectations(t)
	})

}

func TestGetBedsbyid(t *testing.T) {
	repo := new(mocks.BedRepository)
	returnData := bed.BedCore{ID: 1, Ruangan: "A.17", Kelas: "umum", Status: "available", HospitalID: 1}
	id := 1

	t.Run("Succes Get user by id", func(t *testing.T) {
		repo.On("GetById", id).Return(returnData, nil).Once() //menentukan fungsi yang akan dijalankan//

		srv := New(repo) //membuat dependency injection
		data, err := srv.GetById(id)
		assert.Nil(t, err)
		assert.Equal(t, returnData.Kelas, data.Kelas)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Succes Get user by id", func(t *testing.T) {
		repo.On("GetById", id).Return(bed.BedCore{}, errors.New("failed get bed by id data, error logic")).Once() //menentukan fungsi yang akan dijalankan//

		srv := New(repo) //membuat dependency injection
		data, err := srv.GetById(id)
		assert.NotNil(t, err)
		assert.NotEqual(t, returnData.Kelas, data.Kelas)
		assert.Equal(t, "failed get bed by id data, error logic", err.Error())
		repo.AssertExpectations(t)
	})

}

func TestDelete(t *testing.T) {
	repo := new(mocks.BedRepository)
	t.Run("Success Delete", func(t *testing.T) {
		inputRepo := bed.BedCore{ID: 1}
		inputData := bed.BedCore{ID: 1}
		repo.On("Delete", int(inputRepo.ID)).Return(1, nil).Once()
		srv := New(repo)
		err := srv.Delete(int(inputData.ID))
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Delete data", func(t *testing.T) {
		inputRepo := bed.BedCore{ID: 1}
		inputData := bed.BedCore{ID: 1}
		repo.On("Delete", int(inputRepo.ID)).Return(0, errors.New("Gagal Delete")).Once()
		srv := New(repo)
		err := srv.Delete(int(inputData.ID))
		assert.NotNil(t, err)

		repo.AssertExpectations(t)
	})

}
func TestUpdate(t *testing.T) {
	repo := new(mocks.BedRepository)
	t.Run("Success Update", func(t *testing.T) {
		inputRepo := bed.BedCore{NamaTempatTidur: "bunga", Ruangan: "A.17", Kelas: "umum", Status: "available", HospitalID: 1}
		inputData := bed.BedCore{NamaTempatTidur: "bunga", Ruangan: "A.17", Kelas: "umum", Status: "available", HospitalID: 1}
		repo.On("Update", inputRepo, int(inputRepo.ID)).Return(nil).Once()
		srv := New(repo)
		err := srv.Update(inputData, int(inputData.ID))
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Update data", func(t *testing.T) {
		inputRepo := bed.BedCore{NamaTempatTidur: "bunga", Ruangan: "A.17", Kelas: "umum", Status: "available", HospitalID: 1}
		inputData := bed.BedCore{NamaTempatTidur: "bunga", Ruangan: "A.17", Kelas: "umum", Status: "available", HospitalID: 1}
		repo.On("Update", inputRepo, int(inputRepo.ID)).Return(errors.New("gagal mengupdate data , querry error")).Once()
		srv := New(repo)
		err := srv.Update(inputData, int(inputData.ID))
		assert.NotNil(t, err)
		assert.Equal(t, "failed update bed data, error logic", err.Error())

		repo.AssertExpectations(t)
	})

}
