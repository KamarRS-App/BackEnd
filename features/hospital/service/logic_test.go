package service

import (
	"errors"
	"testing"

	"github.com/KamarRS-App/KamarRS-App/features/hospital"
	"github.com/KamarRS-App/KamarRS-App/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	repo := new(mocks.HospitalRepo)
	t.Run("Success Create Hospital", func(t *testing.T) {
		inputRepo := hospital.HospitalCore{KodeRs: "0123456", Nama: "RS. Soebandi", Foto: "abc123.jpg", Alamat: "jl. Kenangan no 123", Provinsi: "Jatim", KabupatenKota: "Surabaya", Kecamatan: "Rungkut", KodePos: "61123", NoTelpon: "08123", Email: "soebandi@gmail.com", KelasRs: "Daerah", PemilikPengelola: "Daerah", JumlahTempatTidur: 100, StatusPenggunaan: "Daerah", BiayaRegistrasi: 25000}
		inputData := hospital.HospitalCore{KodeRs: "0123456", Nama: "RS. Soebandi", Foto: "abc123.jpg", Alamat: "jl. Kenangan no 123", Provinsi: "Jatim", KabupatenKota: "Surabaya", Kecamatan: "Rungkut", KodePos: "61123", NoTelpon: "08123", Email: "soebandi@gmail.com", KelasRs: "Daerah", PemilikPengelola: "Daerah", JumlahTempatTidur: 100, StatusPenggunaan: "Daerah", BiayaRegistrasi: 25000}

		repo.On("Create", inputRepo).Return(nil).Once()
		srv := New(repo)
		err := srv.Create(inputData)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Create Hospital, errValidate", func(t *testing.T) {
		input := hospital.HospitalCore{Nama: "RS. Soebandi", Foto: "abc123.jpg", Alamat: "jl. Kenangan no 123", Provinsi: "Jatim", KabupatenKota: "Surabaya", Kecamatan: "Rungkut", KodePos: "61123", NoTelpon: "08123", Email: "soebandi@gmail.com", KelasRs: "Daerah", PemilikPengelola: "Daerah", JumlahTempatTidur: 100, StatusPenggunaan: "Daerah", BiayaRegistrasi: 25000}

		// repo.On("Create", mock.Anything).Return(nil).Once()
		srv := New(repo)
		err := srv.Create(input)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)

	})

	t.Run("Failed Create Hospital, errCreate", func(t *testing.T) {
		input := hospital.HospitalCore{KodeRs: "0123456", Nama: "RS. Soebandi", Foto: "abc123.jpg", Alamat: "jl. Kenangan no 123", Provinsi: "Jatim", KabupatenKota: "Surabaya", Kecamatan: "Rungkut", KodePos: "61123", NoTelpon: "08123", Email: "soebandi@gmail.com", KelasRs: "Daerah", PemilikPengelola: "Daerah", JumlahTempatTidur: 100, StatusPenggunaan: "Daerah", BiayaRegistrasi: 25000}

		repo.On("Create", mock.Anything).Return(errors.New("failed to insert data, error query")).Once()
		srv := New(repo)
		err := srv.Create(input)
		assert.NotNil(t, err)
		assert.Equal(t, "failed to insert data, error query", err.Error())
		repo.AssertExpectations(t)
	})

}

func TestGetById(t *testing.T) {
	repo := new(mocks.HospitalRepo)
	returnData := hospital.HospitalCore{ID: 1, KodeRs: "0123456", Nama: "RS. Soebandi", Foto: "abc123.jpg", Alamat: "jl. Kenangan no 123", Provinsi: "Jatim", KabupatenKota: "Surabaya", Kecamatan: "Rungkut", KodePos: "61123", NoTelpon: "08123", Email: "soebandi@gmail.com", KelasRs: "Daerah", PemilikPengelola: "Daerah", JumlahTempatTidur: 100, StatusPenggunaan: "Daerah", BiayaRegistrasi: 25000}

	t.Run("Success Get by Id", func(t *testing.T) {
		repo.On("GetById", 1).Return(returnData, nil).Once()

		srv := New(repo)
		response, err := srv.GetById(1)
		assert.Nil(t, err)
		assert.Equal(t, returnData, response)
		repo.AssertExpectations(t)
	})

	t.Run("Failed to Get by Id", func(t *testing.T) {
		repo.On("GetById", 1).Return(hospital.HospitalCore{}, errors.New("failed get hospital by id data, error query")).Once()

		srv := New(repo)
		response, err := srv.GetById(1)
		assert.NotNil(t, err)
		assert.NotEqual(t, returnData, response)
		assert.Equal(t, "failed get hospital by id data, error query", err.Error())
		repo.AssertExpectations(t)
	})
}
