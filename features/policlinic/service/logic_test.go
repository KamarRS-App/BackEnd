package service

import (
	"errors"
	"testing"

	"github.com/KamarRS-App/KamarRS-App/features/policlinic"
	"github.com/KamarRS-App/KamarRS-App/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	repo := new(mocks.PoliclinicRepo)

	t.Run("Succes Create Patient", func(t *testing.T) {
		inputRepo := policlinic.CorePoliclinic{ID: 1, NamaPoli: "Poli Gigi", HospitalID: 1, DoctorID: 1, JamPraktik: "18:40"}
		inputData := policlinic.CorePoliclinic{ID: 1, NamaPoli: "Poli Gigi", HospitalID: 1, DoctorID: 1, JamPraktik: "18:40"}

		repo.On("Create", inputRepo).Return(1, nil).Once()
		srv := New(repo)
		err := srv.Create(inputData)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Create, Empty Name", func(t *testing.T) {
		inputData := policlinic.CorePoliclinic{ID: 1, HospitalID: 1, DoctorID: 1, JamPraktik: "18:40"}

		srv := New(repo)
		err := srv.Create(inputData)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Create new patient", func(t *testing.T) {
		inputRepo := policlinic.CorePoliclinic{ID: 1, NamaPoli: "Poli Gigi", HospitalID: 1, DoctorID: 1, JamPraktik: "18:40"}
		inputData := policlinic.CorePoliclinic{ID: 1, NamaPoli: "Poli Gigi", HospitalID: 1, DoctorID: 1, JamPraktik: "18:40"}
		repo.On("Create", inputRepo).Return(0, errors.New(" Gagal mendaftarkan Poli")).Once()
		srv := New(repo)
		err := srv.Create(inputData)
		assert.NotNil(t, err)
		assert.Equal(t, "failed to insert data, error logic", err.Error()) // samakan dengan di logic
		repo.AssertExpectations(t)
	})

}

func TestUpdate(t *testing.T) {
	repo := new(mocks.PoliclinicRepo)
	id := 1
	t.Run("Success Update", func(t *testing.T) {
		inputRepo := policlinic.CorePoliclinic{ID: 1, NamaPoli: "Poli Gigi", HospitalID: 1, DoctorID: 1, JamPraktik: "18:40"}
		inputData := policlinic.CorePoliclinic{ID: 1, NamaPoli: "Poli Gigi", HospitalID: 1, DoctorID: 1, JamPraktik: "18:40"}
		repo.On("Update", inputRepo, id).Return(nil).Once()
		srv := New(repo)
		err := srv.Update(inputData, id)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Update data", func(t *testing.T) {
		inputRepo := policlinic.CorePoliclinic{ID: 1, NamaPoli: "Poli Gigi", HospitalID: 1, DoctorID: 1, JamPraktik: "18:40"}
		inputData := policlinic.CorePoliclinic{ID: 1, NamaPoli: "Poli Gigi", HospitalID: 1, DoctorID: 1, JamPraktik: "18:40"}
		repo.On("Update", inputRepo, id).Return(errors.New("gagal mengupdate data , querry error")).Once()
		srv := New(repo)
		err := srv.Update(inputData, id)
		assert.NotNil(t, err)
		assert.Equal(t, "failed update policlinic data, error logic", err.Error())
		repo.AssertExpectations(t)
	})

}

func TestDelete(t *testing.T) {
	repo := new(mocks.PoliclinicRepo)
	Poliid := 1
	t.Run("Success Delete", func(t *testing.T) {

		repo.On("Delete", Poliid).Return(1, nil).Once()
		srv := New(repo)
		err := srv.Delete(Poliid)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Delete data", func(t *testing.T) {
		repo.On("Delete", Poliid).Return(0, errors.New("Gagal Delete")).Once()
		srv := New(repo)
		err := srv.Delete(Poliid)
		assert.NotNil(t, err)
		assert.Equal(t, "failed delete policlinic, error logic", err.Error())

		repo.AssertExpectations(t)
	})

}

func TestGetPatientByUserId(t *testing.T) {
	repo := new(mocks.PoliclinicRepo)
	returnData := []policlinic.CorePoliclinic{{ID: 1, NamaPoli: "Poli Gigi", HospitalID: 1, DoctorID: 1, JamPraktik: "18:40"}}
	pagination := 1
	limit := 10
	id := 1
	offset := (pagination - 1) * limit
	t.Run("Succes Get user by id", func(t *testing.T) {
		repo.On("GetAllbyHospitalID", limit, offset, id).Return(returnData, 1, nil).Once() //menentukan fungsi yang akan dijalankan//

		srv := New(repo) //membuat dependency injection
		data, totalPage, err := srv.GetAllbyHospitalID(pagination, limit, id)
		assert.Nil(t, err)
		assert.Equal(t, totalPage, 1)
		assert.Equal(t, data[0].NamaPoli, returnData[0].NamaPoli)
		repo.AssertExpectations(t)
	})

	t.Run("Failed to Get user by id", func(t *testing.T) {
		repo.On("GetAllbyHospitalID", limit, offset, id).Return(nil, 0, errors.New("Gagal menampilkan data")).Once() //menentukan fungsi yang akan dijalankan//

		srv := New(repo)
		data, totalPage, err := srv.GetAllbyHospitalID(pagination, limit, id)
		assert.NotNil(t, err)
		assert.NotEqual(t, totalPage, 1)
		assert.Nil(t, data)
		assert.Equal(t, "failed get policlinic by hospital id, error logic", err.Error())
		repo.AssertExpectations(t)
	})

}

func TestAllGetPatient(t *testing.T) {
	repo := new(mocks.PoliclinicRepo)
	returnData := []policlinic.CorePoliclinic{{ID: 1, NamaPoli: "Poli Gigi", HospitalID: 1, DoctorID: 1, JamPraktik: "18:40"}}
	t.Run("Succes Get user by id", func(t *testing.T) {
		repo.On("GetAll").Return(returnData, nil).Once() //menentukan fungsi yang akan dijalankan//

		srv := New(repo) //membuat dependency injection
		data, err := srv.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, data[0].NamaPoli, returnData[0].NamaPoli)
		repo.AssertExpectations(t)
	})

	t.Run("Failed to Get user by id", func(t *testing.T) {
		repo.On("GetAll").Return(nil, errors.New("Gagal menampilkan data")).Once() //menentukan fungsi yang akan dijalankan//

		srv := New(repo)
		data, err := srv.GetAll()
		assert.NotNil(t, err)
		assert.Nil(t, data)
		assert.Equal(t, "failed get policlinic , error logic", err.Error())
		repo.AssertExpectations(t)
	})

}

func TestGetpatientbyid(t *testing.T) {
	repo := new(mocks.PoliclinicRepo)
	returnData := policlinic.CorePoliclinic{ID: 1, NamaPoli: "Poli Gigi", HospitalID: 1, DoctorID: 1, JamPraktik: "18:40"}

	poliId := 1

	t.Run("Succes Get Poli by id", func(t *testing.T) {
		repo.On("GetById", poliId).Return(returnData, nil).Once() //menentukan fungsi yang akan dijalankan//

		srv := New(repo) //membuat dependency injection
		data, err := srv.GetById(poliId)
		assert.Nil(t, err)
		assert.Equal(t, data.NamaPoli, returnData.NamaPoli)
		repo.AssertExpectations(t)
	})

	t.Run("Failed to Get Poli by id", func(t *testing.T) {
		repo.On("GetById", poliId).Return(policlinic.CorePoliclinic{}, errors.New("Gagal menampilkan data")).Once() //menentukan fungsi yang akan dijalankan//
		srv := New(repo)
		data, err := srv.GetById(poliId)
		assert.NotNil(t, err)
		assert.NotEqual(t, data.NamaPoli, returnData.NamaPoli)
		assert.Equal(t, "failed get policlinic by id data, error logic", err.Error())
		repo.AssertExpectations(t)
	})

}
