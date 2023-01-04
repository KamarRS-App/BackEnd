package service

import (
	"errors"
	"testing"

	checkupreservation "github.com/KamarRS-App/KamarRS-App/features/checkupReservation"
	"github.com/KamarRS-App/KamarRS-App/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	repo := new(mocks.CheckRepo)
	userId := 1
	t.Run("Succes Create Reservation", func(t *testing.T) {
		inputRepo := checkupreservation.CheckupReservationCore{PatientID: 1, PracticeID: 1, NamaDokter: "umum", ID: 1}
		inputData := checkupreservation.CheckupReservationCore{PatientID: 1, PracticeID: 1, NamaDokter: "umum", ID: 1}

		repo.On("Create", inputRepo, userId).Return(nil).Once()
		srv := New(repo)
		err := srv.Create(inputData, userId)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Create, Empty Name", func(t *testing.T) {
		inputData := checkupreservation.CheckupReservationCore{PracticeID: 1, NamaDokter: "umum", ID: 1}
		srv := New(repo)
		err := srv.Create(inputData, userId)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Create, Duplicate", func(t *testing.T) {
		inputRepo := checkupreservation.CheckupReservationCore{PatientID: 1, PracticeID: 1, NamaDokter: "umum", ID: 1}
		inputData := checkupreservation.CheckupReservationCore{PatientID: 1, PracticeID: 1, NamaDokter: "umum", ID: 1}
		repo.On("Create", inputRepo, userId).Return(errors.New("reservasi checkup gagal")).Once()
		srv := New(repo)
		err := srv.Create(inputData, userId)
		assert.NotNil(t, err)
		assert.Equal(t, " reservasi checkup gagal", err.Error()) // samakan dengan di logic
		repo.AssertExpectations(t)
	})

}

func TestGetByPracticesId(t *testing.T) {
	repo := new(mocks.CheckRepo)
	returnData := []checkupreservation.CheckupReservationCore{{PatientID: 1, PracticeID: 1, NamaDokter: "Teguh", ID: 1}}
	pagination := 1
	PracticeID := 1
	limit := 10
	offset := (pagination - 1) * limit
	t.Run("Succes Get user by id", func(t *testing.T) {
		repo.On("GetByPracticesId", limit, offset, PracticeID).Return(returnData, 1, nil).Once() //menentukan fungsi yang akan dijalankan//

		srv := New(repo) //membuat dependency injection
		data, totalPage, err := srv.GetByPracticesId(pagination, limit, PracticeID)
		assert.Nil(t, err)
		assert.Equal(t, totalPage, 1)
		assert.Equal(t, data[0].NamaDokter, returnData[0].NamaDokter)
		repo.AssertExpectations(t)
	})

	t.Run("Failed to Get user by id", func(t *testing.T) {
		repo.On("GetByPracticesId", limit, offset, PracticeID).Return(nil, 0, errors.New("Gagal menampilkan data")).Once() //menentukan fungsi yang akan dijalankan//

		srv := New(repo)
		data, totalPage, err := srv.GetByPracticesId(pagination, limit, PracticeID)
		assert.NotNil(t, err)
		assert.NotEqual(t, totalPage, 1)
		assert.Nil(t, data)
		assert.Equal(t, "gagal menampilkan reservasi", err.Error())
		repo.AssertExpectations(t)
	})

}

func TestGetBedsbyid(t *testing.T) {
	repo := new(mocks.CheckRepo)
	returnData := checkupreservation.CheckupReservationCore{PatientID: 1, PracticeID: 1, NamaDokter: "Teguh", ID: 1}
	id := 1

	t.Run("Succes Get user by id", func(t *testing.T) {
		repo.On("GetByreservationId", id).Return(returnData, nil).Once() //menentukan fungsi yang akan dijalankan//

		srv := New(repo) //membuat dependency injection
		data, err := srv.GetByreservationId(id)
		assert.Nil(t, err)
		assert.Equal(t, returnData.NamaDokter, data.NamaDokter)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Succes Get user by id", func(t *testing.T) {
		repo.On("GetByreservationId", id).Return(checkupreservation.CheckupReservationCore{}, errors.New("failed get reservation by id")).Once() //menentukan fungsi yang akan dijalankan//

		srv := New(repo) //membuat dependency injection
		data, err := srv.GetByreservationId(id)
		assert.NotNil(t, err)
		assert.NotEqual(t, returnData.NamaDokter, data.NamaDokter)
		assert.Equal(t, "gagal menampilkan reservasi", err.Error())
		repo.AssertExpectations(t)
	})

}
