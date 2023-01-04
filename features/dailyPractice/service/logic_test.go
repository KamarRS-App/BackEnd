package service

import (
	"errors"
	"testing"

	dailypractice "github.com/KamarRS-App/KamarRS-App/features/dailyPractice"
	"github.com/KamarRS-App/KamarRS-App/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	repo := new(mocks.DailyPracticeRepository)
	t.Run("success create", func(t *testing.T) {
		input := dailypractice.PracticeCore{PoliclinicID: 1, TanggalPraktik: "2022-12-27", KuotaHarian: 10, Status: "available"}
		repo.On("Create", mock.Anything).Return(1, nil).Once()
		srv := New(repo)
		err := srv.Create(input)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed create, err", func(t *testing.T) {
		input := dailypractice.PracticeCore{PoliclinicID: 1, TanggalPraktik: "2022-12-27", KuotaHarian: 10, Status: "available"}
		repo.On("Create", mock.Anything).Return(0, errors.New("failed to insert data, error logic")).Once()
		srv := New(repo)
		err := srv.Create(input)
		assert.EqualError(t, err, "failed to insert data, error logic")
		repo.AssertExpectations(t)
	})

}

func TestGetById(t *testing.T) {
	repo := new(mocks.DailyPracticeRepository)
	t.Run("success get by id", func(t *testing.T) {
		inputId := 1
		returnData := dailypractice.PracticeCore{PoliclinicID: 1, TanggalPraktik: "2022-12-27", KuotaHarian: 10, Status: "available"}
		repo.On("GetById", mock.Anything).Return(returnData, nil).Once()
		srv := New(repo)
		res, err := srv.GetById(inputId)
		assert.Nil(t, err)
		assert.Equal(t, res.Status, returnData.Status)
		repo.AssertExpectations(t)
	})

	t.Run("failed get by id", func(t *testing.T) {
		inputId := 2
		repo.On("GetById", mock.Anything).Return(dailypractice.PracticeCore{}, errors.New("failed get practice by id data, error logic")).Once()
		srv := New(repo)
		_, err := srv.GetById(inputId)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	repo := new(mocks.DailyPracticeRepository)
	t.Run("success get all", func(t *testing.T) {
		inputPagination := 1
		inputLimit := 10
		inputPoliId := 1
		inputOffset := (inputPagination - 1) * inputLimit
		returnData := []dailypractice.PracticeCore{{PoliclinicID: 1, TanggalPraktik: "2022-12-27", KuotaHarian: 10, Status: "available"}, {PoliclinicID: 1, TanggalPraktik: "2022-12-27", KuotaHarian: 10, Status: "available"}}
		repo.On("GetAll", inputLimit, inputOffset, inputPoliId).Return(returnData, 1, nil).Once()
		srv := New(repo)
		res, pages, err := srv.GetAll(inputPagination, inputLimit, inputPoliId)
		assert.Nil(t, err)
		assert.Equal(t, pages, 1)
		assert.Equal(t, res[0].Status, returnData[0].Status)
		repo.AssertExpectations(t)
	})

	t.Run("failed get all", func(t *testing.T) {
		inputPagination := 1
		inputLimit := 10
		inputPoliId := 1
		inputOffset := (inputPagination - 1) * inputLimit
		repo.On("GetAll", inputLimit, inputOffset, inputPoliId).Return(nil, 0, errors.New("failed get practice by policlinic id, error logic")).Once()
		srv := New(repo)
		_, pages, err := srv.GetAll(inputPagination, inputLimit, inputPoliId)
		assert.NotNil(t, err)
		assert.Equal(t, pages, 0)
		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	repo := new(mocks.DailyPracticeRepository)
	t.Run("success update", func(t *testing.T) {
		input := dailypractice.PracticeCore{ID: 1, PoliclinicID: 1, TanggalPraktik: "2022-12-27", KuotaHarian: 10, Status: "available"}
		repo.On("Update", input, 1).Return(nil).Once()
		srv := New(repo)
		err := srv.Update(input, 1)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed update", func(t *testing.T) {
		input := dailypractice.PracticeCore{ID: 1, PoliclinicID: 1, TanggalPraktik: "2022-12-27", KuotaHarian: 10, Status: "available"}
		repo.On("Update", input, 1).Return(errors.New("failed update practice data, error logic")).Once()
		srv := New(repo)
		err := srv.Update(input, 1)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}
