package service

import (
	"errors"
	"testing"

	"github.com/KamarRS-App/KamarRS-App/features/kamarrsteam"
	"github.com/KamarRS-App/KamarRS-App/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	repo := new(mocks.KamarRsTeamRepository)
	t.Run("success create", func(t *testing.T) {
		input := kamarrsteam.KamarRsTeamCore{Email: "sheena@gmail.com", KataSandi: "password", Peran: "super admin"}
		repo.On("FindTeam", input.Email).Return(kamarrsteam.KamarRsTeamCore{}, nil).Once()
		repo.On("Create", mock.Anything).Return(nil).Once()
		srv := New(repo)
		err := srv.Create(input)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed create, errCreate", func(t *testing.T) {
		input := kamarrsteam.KamarRsTeamCore{Email: "sheena@gmail.com", KataSandi: "password", Peran: "super admin"}
		repo.On("FindTeam", input.Email).Return(kamarrsteam.KamarRsTeamCore{}, nil).Once()
		repo.On("Create", mock.Anything).Return(errors.New("error query")).Once()
		srv := New(repo)
		err := srv.Create(input)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed create, errValidate", func(t *testing.T) {
		input := kamarrsteam.KamarRsTeamCore{Email: "sheena@gmail.com", Peran: "super admin"}
		srv := New(repo)
		err := srv.Create(input)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed create, used email", func(t *testing.T) {
		input := kamarrsteam.KamarRsTeamCore{Email: "sheena@gmail.com", KataSandi: "password", Peran: "super admin"}
		repo.On("FindTeam", input.Email).Return(kamarrsteam.KamarRsTeamCore{Email: "sheena@gmail.com", KataSandi: "password", Peran: "super admin"}, errors.New("use another email")).Once()
		srv := New(repo)
		err := srv.Create(input)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)

	})
}
