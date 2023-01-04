package service

import (
	"errors"

	checkupreservation "github.com/KamarRS-App/KamarRS-App/features/checkupReservation"
	"github.com/go-playground/validator/v10"
)

type CheckUpService struct {
	checkupRepository checkupreservation.RepositoryInterface //data repository dri entities
	validate          *validator.Validate
}

func New(repo checkupreservation.RepositoryInterface) checkupreservation.ServiceInterface { //dengan kembalian user.service
	return &CheckUpService{
		checkupRepository: repo,
		validate:          validator.New(),
	}
}

// Create implements checkupreservation.ServiceInterface
func (service *CheckUpService) Create(input checkupreservation.CheckupReservationCore, userId int) (err error) {
	if validateERR := service.validate.Struct(input); validateERR != nil {
		return validateERR
	}

	errCreate := service.checkupRepository.Create(input, userId)

	if errCreate != nil {
		return errors.New(" reservasi checkup gagal")
	}

	return nil
}

// GetByPracticesId implements checkupreservation.ServiceInterface
func (service *CheckUpService) GetByPracticesId(pagination int, limit int, id int) (data []checkupreservation.CheckupReservationCore, totalpage int, err error) {
	offset := (pagination - 1) * limit
	data, totalpage, err = service.checkupRepository.GetByPracticesId(limit, offset, id) // memanggil struct entities repository yang ada di entities yang berisi coding logic
	if err != nil {
		return nil, 0, errors.New("gagal menampilkan reservasi")
	}
	return
}

// GetByreservationId implements checkupreservation.ServiceInterface
func (service *CheckUpService) GetByreservationId(id int) (data checkupreservation.CheckupReservationCore, err error) {
	data, err = service.checkupRepository.GetByreservationId(id) // memanggil struct entities repository yang ada di entities yang berisi coding logic
	if err != nil {
		return checkupreservation.CheckupReservationCore{}, errors.New("gagal menampilkan reservasi")
	}
	return
}
