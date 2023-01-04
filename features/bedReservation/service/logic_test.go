package service

import (
	"errors"
	"testing"

	bedreservation "github.com/KamarRS-App/KamarRS-App/features/bedReservation"
	"github.com/KamarRS-App/KamarRS-App/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateBedReservation(t *testing.T) {
	repo := new(mocks.BedReservationRepository)
	t.Run("succes update", func(t *testing.T) {
		input := bedreservation.BedReservationCore{ID: 1, StatusPasien: "pemeriksaan awal", BedID: 1}
		repo.On("UpdateBedReservation", mock.Anything).Return(nil).Once()
		srv := New(repo)
		err := srv.UpdateBedReservation(input)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed update, err", func(t *testing.T) {
		input := bedreservation.BedReservationCore{ID: 1, StatusPasien: "pemeriksaan awal", BedID: 1}
		repo.On("UpdateBedReservation", mock.Anything).Return(errors.New("failed query")).Once()
		srv := New(repo)
		err := srv.UpdateBedReservation(input)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

}

func TestDelete(t *testing.T) {
	repo := new(mocks.BedReservationRepository)
	t.Run("succes delete", func(t *testing.T) {
		inputId := 1
		repo.On("Delete", mock.Anything).Return(nil).Once()
		srv := New(repo)
		err := srv.Delete(uint(inputId))
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed delete, err", func(t *testing.T) {
		inputId := 1
		repo.On("Delete", mock.Anything).Return(errors.New("failed query")).Once()
		srv := New(repo)
		err := srv.Delete(uint(inputId))
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	repo := new(mocks.BedReservationRepository)
	t.Run("succes get by id", func(t *testing.T) {
		inputId := 1
		returnData := bedreservation.BedReservationCore{ID: 1, StatusPasien: "pemeriksaan awal", BiayaRegistrasi: 37000, KodeDaftar: "order-UiK85", BedID: 1}
		repo.On("GetById", mock.Anything).Return(returnData, nil).Once()
		srv := New(repo)
		res, err := srv.GetById(uint(inputId))
		assert.Nil(t, err)
		assert.Equal(t, res.KodeDaftar, returnData.KodeDaftar)
		repo.AssertExpectations(t)
	})

	t.Run("failed get by id", func(t *testing.T) {
		inputId := 1
		repo.On("GetById", mock.Anything).Return(bedreservation.BedReservationCore{}, errors.New("error query")).Once()
		srv := New(repo)
		res, err := srv.GetById(uint(inputId))
		assert.EqualError(t, err, "error query")
		assert.Equal(t, res, bedreservation.BedReservationCore{})
		repo.AssertExpectations(t)
	})
}

func TestGetRegistrations(t *testing.T) {
	repo := new(mocks.BedReservationRepository)
	t.Run("succes get registrations", func(t *testing.T) {
		inputPage := 1
		inputLimit := 10
		inputHospitalId := 1
		inputOffset := (inputPage - 1) * inputLimit
		returnData := []bedreservation.BedReservationCore{{ID: 1, StatusPasien: "pemeriksaan awal", BiayaRegistrasi: 37000, KodeDaftar: "order-UiK85", BedID: 1, HospitalID: 1}, {ID: 2, StatusPasien: "pemeriksaan awal", BiayaRegistrasi: 37000, KodeDaftar: "order-UiK86", BedID: 2, HospitalID: 1}}
		repo.On("GetRegistrations", inputLimit, inputOffset, inputHospitalId).Return(returnData, 1, nil).Once()
		srv := New(repo)
		res, pages, err := srv.GetRegistrations(inputPage, inputLimit, inputHospitalId)
		assert.Nil(t, err)
		assert.Equal(t, res[0].HospitalID, returnData[0].HospitalID)
		assert.Equal(t, pages, 1)
		repo.AssertExpectations(t)
	})

	t.Run("failed get registrations, err", func(t *testing.T) {
		inputPage := 1
		inputLimit := 10
		inputHospitalId := 1
		inputOffset := (inputPage - 1) * inputLimit
		repo.On("GetRegistrations", inputLimit, inputOffset, inputHospitalId).Return(nil, 0, errors.New("error query")).Once()
		srv := New(repo)
		res, pages, err := srv.GetRegistrations(inputPage, inputLimit, inputHospitalId)
		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, pages, 0)
		repo.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	repo := new(mocks.BedReservationRepository)
	t.Run("succes create registrations", func(t *testing.T) {
		inputUserId := uint(1)
		inputData := bedreservation.BedReservationCore{HospitalID: 1, PatientID: 1}
		returnData := bedreservation.BedReservationCore{ID: 1, BiayaRegistrasi: 37000, KodeDaftar: "order-UiK85", PatientID: 1, HospitalID: 1, StatusPembayaran: "belum dibayar"}
		repo.On("Create", inputData, inputUserId).Return(returnData, nil).Once()
		srv := New(repo)
		res, err := srv.Create(inputData, uint(inputUserId))
		assert.Nil(t, err)
		assert.Equal(t, res.StatusPembayaran, "belum dibayar")
		repo.AssertExpectations(t)
	})

	t.Run("failed create registrations", func(t *testing.T) {
		inputUserId := uint(1)
		inputData := bedreservation.BedReservationCore{HospitalID: 1, PatientID: 1}
		repo.On("Create", inputData, inputUserId).Return(bedreservation.BedReservationCore{}, errors.New("error query")).Once()
		srv := New(repo)
		res, err := srv.Create(inputData, uint(inputUserId))
		assert.NotNil(t, err)
		assert.Equal(t, res, bedreservation.BedReservationCore{})
		repo.AssertExpectations(t)
	})
}

func TestGetPayment(t *testing.T) {
	repo := new(mocks.BedReservationRepository)
	t.Run("succes get payment", func(t *testing.T) {
		inputKodeDaftar := "order-UiK85"
		returnData := bedreservation.BedReservationCore{ID: 1, BiayaRegistrasi: 37000, KodeDaftar: "order-UiK85", PatientID: 1, HospitalID: 1, StatusPembayaran: "belum dibayar"}
		repo.On("GetPayment", mock.Anything).Return(returnData, nil).Once()
		srv := New(repo)
		res, err := srv.GetPayment(inputKodeDaftar)
		assert.Nil(t, err)
		assert.Equal(t, res.StatusPembayaran, "belum dibayar")
		repo.AssertExpectations(t)
	})

	t.Run("failed get payment", func(t *testing.T) {
		inputKodeDaftar := "order-UiK85"
		repo.On("GetPayment", mock.Anything).Return(bedreservation.BedReservationCore{}, errors.New("cannot find data")).Once()
		srv := New(repo)
		_, err := srv.GetPayment(inputKodeDaftar)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestCreatePayment(t *testing.T) {
	repo := new(mocks.BedReservationRepository)
	t.Run("succes get payment", func(t *testing.T) {
		inputData := bedreservation.BedReservationCore{KodeDaftar: "order-UiK85", PaymentMethod: "transfer_va_bca"}
		returnData := bedreservation.BedReservationCore{KodeDaftar: "order-UiK85", PatientID: 1, HospitalID: 1, StatusPembayaran: "pending", StatusPasien: "waiting list", BiayaRegistrasi: 37000, PaymentMethod: "transfer_va_bca", LinkPembayaran: "https://simulator.sandbox.midtrans.com/bca/va/index", BankPenerima: "bank bca"}
		repo.On("CreatePayment", mock.Anything).Return(returnData, nil).Once()
		srv := New(repo)
		res, err := srv.CreatePayment(inputData)
		assert.Nil(t, err)
		assert.Equal(t, res.StatusPembayaran, "pending")
		repo.AssertExpectations(t)
	})

	t.Run("failed get payment", func(t *testing.T) {
		inputData := bedreservation.BedReservationCore{KodeDaftar: "order-UiK85", PaymentMethod: "qris"}
		repo.On("CreatePayment", mock.Anything).Return(bedreservation.BedReservationCore{}, errors.New("pilih metode pembayaran lain")).Once()
		srv := New(repo)
		_, err := srv.CreatePayment(inputData)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestPaymentNotif(t *testing.T) {
	repo := new(mocks.BedReservationRepository)
	t.Run("succes callback", func(t *testing.T) {
		inputData := bedreservation.BedReservationCore{KodeDaftar: "order-UiK85"}
		repo.On("PaymentNotif", mock.Anything).Return(nil).Once()
		srv := New(repo)
		err := srv.PaymentNotif(inputData)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
}
