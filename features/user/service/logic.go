package service

import (
	"errors"

	"github.com/KamarRS-App/features/user"

	"github.com/go-playground/validator/v10"
)

type UserService struct {
	userRepository user.RepositoryInterface //data repository dri entities
	validate       *validator.Validate
}

func New(repo user.RepositoryInterface) user.ServiceInterface { //dengan kembalian user.service
	return &UserService{
		userRepository: repo,
		validate:       validator.New(),
	}
}

// Create implements user.ServiceEntities
func (service *UserService) Create(input user.CoreUser) (err error) {
	// input.Role = "User"
	if validateERR := service.validate.Struct(input); validateERR != nil {
		return validateERR
	}

	errCreate := service.userRepository.Create(input)
	if errCreate != nil {
		return errors.New("GAGAL MENAMBAH DATA , QUERY ERROR")
	}

	return nil
}
