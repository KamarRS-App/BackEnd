package service

import (
	"errors"
	"testing"

	"github.com/KamarRS-App/KamarRS-App/features/auth"
	staff "github.com/KamarRS-App/KamarRS-App/features/hospitalstaff/repository"
	teamrepo "github.com/KamarRS-App/KamarRS-App/features/kamarrsteam/repository"
	"github.com/KamarRS-App/KamarRS-App/features/user/repository"
	"github.com/KamarRS-App/KamarRS-App/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin(t *testing.T) {
	repo := new(mocks.AuthRepository)
	t.Run("success login", func(t *testing.T) {
		inputEmail := "sheena@duck.com"
		inputPass := "ringo"
		returnData := repository.User{Nama: "sheena", Email: "sheena@duck.com", Nokk: "1234567890", Nik: "0987654321", KataSandi: "ringo", NoTelpon: "12309876534"}
		returnToken := "klafanasndfiuweof"
		repo.On("Login", inputEmail, inputPass).Return(returnToken, returnData, nil).Once()
		srv := New(repo)
		token, response, err := srv.Login(inputEmail, inputPass)
		assert.Nil(t, err)
		assert.Equal(t, returnData, response)
		assert.Equal(t, token, returnToken)
		repo.AssertExpectations(t)
	})
}

func TestStaff(t *testing.T) {
	repo := new(mocks.AuthRepository)
	t.Run("success login", func(t *testing.T) {
		inputEmail := "sheena@duck.com"
		inputPass := "ringo"
		returnData := staff.HospitalStaff{Nama: "sheena", Email: "sheena@duck.com", KataSandi: "ringo", Peran: "admin", HospitalID: 1, HospitalName: "rs"}
		returnToken := "klafanasndfiuweof"
		repo.On("LoginStaff", inputEmail, inputPass).Return(returnToken, returnData, nil).Once()
		srv := New(repo)
		token, response, err := srv.LoginStaff(inputEmail, inputPass)
		assert.Nil(t, err)
		assert.Equal(t, returnData, response)
		assert.Equal(t, token, returnToken)
		repo.AssertExpectations(t)
	})
}

func TestLoginOauth(t *testing.T) {
	repo := new(mocks.AuthRepository)
	t.Run("success login", func(t *testing.T) {
		inputData := auth.Oauth{Email: "sheena@gmail.com", Name: "sheena"}
		returnData := repository.User{Nama: "sheena", Email: "sheena@duck.com"}
		returnToken := "klafanasndfiuweof"
		repo.On("LoginOauth", mock.Anything).Return(returnToken, returnData, nil).Once()
		srv := New(repo)
		token, response, err := srv.LoginOauth(inputData)
		assert.Nil(t, err)
		assert.Equal(t, returnData, response)
		assert.Equal(t, token, returnToken)
		repo.AssertExpectations(t)
	})
}

func TestLoginTeam(t *testing.T) {
	repo := new(mocks.AuthRepository)
	t.Run("success login", func(t *testing.T) {
		inputEmail := "sheena@duck.com"
		inputPass := "ringo"
		returnData := teamrepo.KamarRsTeam{Email: "sheena@duck.com", KataSandi: "ringo", Peran: "super admin"}
		returnToken := "klafanasndfiuweof"
		repo.On("LoginTeam", inputEmail, inputPass).Return(returnToken, returnData, nil).Once()
		srv := New(repo)
		token, response, err := srv.LoginTeam(inputEmail, inputPass)
		assert.Nil(t, err)
		assert.Equal(t, returnData, response)
		assert.Equal(t, token, returnToken)
		repo.AssertExpectations(t)
	})

	t.Run("failed login, empty pass", func(t *testing.T) {
		inputEmail := "sheena@duck.com"
		inputPass := ""
		// repo.On("LoginTeam", inputEmail, inputPass).Return("", teamrepo.KamarRsTeam{}, errors.New("field must be filled")).Once()
		srv := New(repo)
		token, data, err := srv.LoginTeam(inputEmail, inputPass)
		assert.Error(t, errors.New("field must be filled"))
		assert.EqualError(t, err, "field must be filled")
		assert.Equal(t, token, "")
		assert.Equal(t, data, teamrepo.KamarRsTeam{})
		repo.AssertExpectations(t)
	})

	t.Run("failed login, err", func(t *testing.T) {
		inputEmail := "sheena@duck.com"
		inputPass := "ringo"
		// returnData := teamrepo.KamarRsTeam{Email: "sheena@duck.com", KataSandi: "ringo", Peran: "super admin"}
		returnToken := ""
		repo.On("LoginTeam", inputEmail, inputPass).Return("", teamrepo.KamarRsTeam{}, errors.New("cannot find data")).Once()
		srv := New(repo)
		token, _, err := srv.LoginTeam(inputEmail, inputPass)
		assert.NotNil(t, err)
		assert.Equal(t, token, returnToken)
		repo.AssertExpectations(t)

	})
}
