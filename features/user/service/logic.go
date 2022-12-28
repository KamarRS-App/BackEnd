package service

import (
	"errors"
	"strings"

	"github.com/KamarRS-App/KamarRS-App/features/user"
	"github.com/KamarRS-App/KamarRS-App/utils/helper"

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
	lower := strings.ToLower(input.Email)

	input.Email = lower

	generatePass := helper.Bcript(input.KataSandi)

	input.KataSandi = generatePass

	if validateERR := service.validate.Struct(input); validateERR != nil {
		return validateERR
	}

	errCreate := service.userRepository.Create(input)
	if errCreate != nil {
		return errors.New(" Gagal membuat akun, Email sudah terdaftar")
	}

	return nil
}

// Update implements user.ServiceInterface
func (service *UserService) Update(id int, input user.CoreUser) error {
	errUpdate := service.userRepository.Update(id, input)

	if errUpdate != nil {
		return errors.New("gagal mengupdate data , querry error")
	}

	return nil
}

// GetById implements user.ServiceInterface
func (service *UserService) GetById(id int) (data user.CoreUser, err error) {
	data, err = service.userRepository.GetById(id) // memanggil struct entities repository yang ada di entities yang berisi coding logic
	return
}

// DeleteById implements user.ServiceInterface
func (service *UserService) DeleteById(id int) error {
	err := service.userRepository.DeleteById(id) // memanggil struct entities repository yang ada di entities yang berisi coding logic
	return err
}
