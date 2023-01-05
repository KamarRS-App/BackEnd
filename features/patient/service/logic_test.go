package service

import (
	"errors"
	"testing"

	"github.com/KamarRS-App/KamarRS-App/features/patient"
	"github.com/KamarRS-App/KamarRS-App/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	repo := new(mocks.PatientRepo)

	t.Run("Succes Create Patient", func(t *testing.T) {
		inputRepo := patient.CorePatient{ID: 1, NoKk: "123", Nik: "321", NamaPasien: "teguh", JenisKelamin: "laki", TanggalLahir: "16-07-1998", Usia: 20, NamaWali: "Ahmad", EmailWali: "ahmad@mail.com", NoTelponWali: "08123", AlamatKtp: "Bekasi", AlamatDomisili: "Manado"}
		inputData := patient.CorePatient{ID: 1, NoKk: "123", Nik: "321", NamaPasien: "teguh", JenisKelamin: "laki", TanggalLahir: "16-07-1998", Usia: 20, NamaWali: "Ahmad", EmailWali: "ahmad@mail.com", NoTelponWali: "08123", AlamatKtp: "Bekasi", AlamatDomisili: "Manado"}

		repo.On("Create", inputRepo).Return(nil).Once()
		srv := New(repo)
		err := srv.Create(inputData)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Create, Empty Name", func(t *testing.T) {
		inputData := patient.CorePatient{ID: 1, NoKk: "123", Nik: "321", JenisKelamin: "laki", TanggalLahir: "16-07-1998", Usia: 20, NamaWali: "Ahmad", EmailWali: "ahmad@mail.com", NoTelponWali: "08123", AlamatKtp: "Bekasi", AlamatDomisili: "Manado"}
		srv := New(repo)
		err := srv.Create(inputData)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Create new patient", func(t *testing.T) {
		inputRepo := patient.CorePatient{ID: 1, NoKk: "123", Nik: "321", NamaPasien: "teguh", JenisKelamin: "laki", TanggalLahir: "16-07-1998", Usia: 20, NamaWali: "Ahmad", EmailWali: "ahmad@mail.com", NoTelponWali: "08123", AlamatKtp: "Bekasi", AlamatDomisili: "Manado"}
		inputData := patient.CorePatient{ID: 1, NoKk: "123", Nik: "321", NamaPasien: "teguh", JenisKelamin: "laki", TanggalLahir: "16-07-1998", Usia: 20, NamaWali: "Ahmad", EmailWali: "ahmad@mail.com", NoTelponWali: "08123", AlamatKtp: "Bekasi", AlamatDomisili: "Manado"}
		repo.On("Create", inputRepo).Return(errors.New(" Gagal mendaftarkan pasien")).Once()
		srv := New(repo)
		err := srv.Create(inputData)
		assert.NotNil(t, err)
		assert.Equal(t, " Kesalahan pada input data pasien", err.Error()) // samakan dengan di logic
		repo.AssertExpectations(t)
	})

}

func TestUpdate(t *testing.T) {
	repo := new(mocks.PatientRepo)
	userid := 1
	t.Run("Success Update", func(t *testing.T) {
		inputRepo := patient.CorePatient{ID: 1, NoKk: "123", Nik: "321", NamaPasien: "teguh", JenisKelamin: "laki", TanggalLahir: "16-07-1998", Usia: 20, NamaWali: "Ahmad", EmailWali: "ahmad@mail.com", NoTelponWali: "08123", AlamatKtp: "Bekasi", AlamatDomisili: "Manado", UserID: 1}
		inputData := patient.CorePatient{ID: 1, NoKk: "123", Nik: "321", NamaPasien: "teguh", JenisKelamin: "laki", TanggalLahir: "16-07-1998", Usia: 20, NamaWali: "Ahmad", EmailWali: "ahmad@mail.com", NoTelponWali: "08123", AlamatKtp: "Bekasi", AlamatDomisili: "Manado", UserID: 1}
		repo.On("Update", int(inputRepo.ID), userid, inputRepo).Return(nil).Once()
		srv := New(repo)
		err := srv.Update(int(inputData.ID), userid, inputData)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Update data", func(t *testing.T) {
		inputRepo := patient.CorePatient{ID: 1, NoKk: "123", Nik: "321", NamaPasien: "teguh", JenisKelamin: "laki", TanggalLahir: "16-07-1998", Usia: 20, NamaWali: "Ahmad", EmailWali: "ahmad@mail.com", NoTelponWali: "08123", AlamatKtp: "Bekasi", AlamatDomisili: "Manado", UserID: 1}
		inputData := patient.CorePatient{ID: 1, NoKk: "123", Nik: "321", NamaPasien: "teguh", JenisKelamin: "laki", TanggalLahir: "16-07-1998", Usia: 20, NamaWali: "Ahmad", EmailWali: "ahmad@mail.com", NoTelponWali: "08123", AlamatKtp: "Bekasi", AlamatDomisili: "Manado", UserID: 1}
		repo.On("Update", int(inputRepo.ID), userid, inputRepo).Return(errors.New("gagal mengupdate data , querry error")).Once()
		srv := New(repo)
		err := srv.Update(int(inputData.ID), userid, inputData)
		assert.NotNil(t, err)
		assert.Equal(t, "gagal mengupdate data , querry error", err.Error())

		repo.AssertExpectations(t)
	})

}

func TestDelete(t *testing.T) {
	repo := new(mocks.PatientRepo)
	patientId := 1
	t.Run("Success Delete", func(t *testing.T) {

		repo.On("DeleteById", patientId).Return(nil).Once()
		srv := New(repo)
		err := srv.DeleteById(patientId)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Delete data", func(t *testing.T) {
		repo.On("DeleteById", patientId).Return(errors.New("Gagal Delete")).Once()
		srv := New(repo)
		err := srv.DeleteById(patientId)
		assert.NotNil(t, err)
		assert.Equal(t, "gagal menghapus data pasien", err.Error())

		repo.AssertExpectations(t)
	})

}

func TestGetPatientByUserId(t *testing.T) {
	repo := new(mocks.PatientRepo)
	returnData := []patient.CorePatient{{ID: 1, NoKk: "123", Nik: "321", NamaPasien: "teguh", JenisKelamin: "laki", TanggalLahir: "16-07-1998", Usia: 20, NamaWali: "Ahmad", EmailWali: "ahmad@mail.com", NoTelponWali: "08123", AlamatKtp: "Bekasi", AlamatDomisili: "Manado", UserID: 1}}
	pagination := 1
	limit := 10
	userid := 1
	offset := (pagination - 1) * limit
	t.Run("Succes Get user by id", func(t *testing.T) {
		repo.On("GetByUserId", limit, offset, userid).Return(returnData, 1, nil).Once() //menentukan fungsi yang akan dijalankan//

		srv := New(repo) //membuat dependency injection
		data, totalPage, err := srv.GetByUserId(pagination, limit, userid)
		assert.Nil(t, err)
		assert.Equal(t, totalPage, 1)
		assert.Equal(t, data[0].NamaPasien, returnData[0].NamaPasien)
		repo.AssertExpectations(t)
	})

	t.Run("Failed to Get user by id", func(t *testing.T) {
		repo.On("GetByUserId", limit, offset, userid).Return(nil, 0, errors.New("Gagal menampilkan data")).Once() //menentukan fungsi yang akan dijalankan//

		srv := New(repo)
		data, totalPage, err := srv.GetByUserId(pagination, limit, userid)
		assert.NotNil(t, err)
		assert.NotEqual(t, totalPage, 1)
		assert.Nil(t, data)
		assert.Equal(t, "gagal menampilkan pasien", err.Error())
		repo.AssertExpectations(t)
	})

}

func TestGetpatientbyid(t *testing.T) {
	repo := new(mocks.PatientRepo)
	returnData := patient.CorePatient{ID: 1, NoKk: "123", Nik: "321", NamaPasien: "teguh", JenisKelamin: "laki", TanggalLahir: "16-07-1998", Usia: 20, NamaWali: "Ahmad", EmailWali: "ahmad@mail.com", NoTelponWali: "08123", AlamatKtp: "Bekasi", AlamatDomisili: "Manado", UserID: 1}

	patientID := 1

	t.Run("Succes Get user by id", func(t *testing.T) {
		repo.On("GetByPatientId", patientID).Return(returnData, nil).Once() //menentukan fungsi yang akan dijalankan//

		srv := New(repo) //membuat dependency injection
		data, err := srv.GetByPatientId(patientID)
		assert.Nil(t, err)
		assert.Equal(t, data.NamaPasien, returnData.NamaPasien)
		repo.AssertExpectations(t)
	})

	t.Run("Failed to Get user by id", func(t *testing.T) {
		repo.On("GetByPatientId", patientID).Return(patient.CorePatient{}, errors.New("Gagal menampilkan data")).Once() //menentukan fungsi yang akan dijalankan//
		srv := New(repo)
		data, err := srv.GetByPatientId(patientID)
		assert.NotNil(t, err)
		assert.NotEqual(t, data.NamaPasien, returnData.NamaPasien)
		assert.Equal(t, "gagal menampilkan pasien", err.Error())
		repo.AssertExpectations(t)
	})

}

func TestAllGetPatient(t *testing.T) {
	repo := new(mocks.PatientRepo)
	returnData := []patient.CorePatient{{ID: 1, NoKk: "123", Nik: "321", NamaPasien: "teguh", JenisKelamin: "laki", TanggalLahir: "16-07-1998", Usia: 20, NamaWali: "Ahmad", EmailWali: "ahmad@mail.com", NoTelponWali: "08123", AlamatKtp: "Bekasi", AlamatDomisili: "Manado", UserID: 1}}

	t.Run("Succes Get user by id", func(t *testing.T) {
		repo.On("GetAllPatient").Return(returnData, nil).Once() //menentukan fungsi yang akan dijalankan//

		srv := New(repo) //membuat dependency injection
		data, err := srv.GetAllPatient()
		assert.Nil(t, err)
		assert.Equal(t, data[0].NamaPasien, returnData[0].NamaPasien)
		repo.AssertExpectations(t)
	})

	t.Run("Failed to Get user by id", func(t *testing.T) {
		repo.On("GetAllPatient").Return(nil, errors.New("Gagal menampilkan data")).Once() //menentukan fungsi yang akan dijalankan//

		srv := New(repo)
		data, err := srv.GetAllPatient()
		assert.NotNil(t, err)
		assert.Nil(t, data)
		assert.Equal(t, "gagal menampilkan pasien", err.Error())
		repo.AssertExpectations(t)
	})

}
