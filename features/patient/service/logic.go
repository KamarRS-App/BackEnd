package service

import (
	"errors"

	"github.com/KamarRS-App/features/patient"
	"github.com/go-playground/validator/v10"
)

type patientService struct {
	patientRepository patient.RepositoryInterface //data repository dri entities
	validate          *validator.Validate
}

func New(repo patient.RepositoryInterface) patient.ServiceInterface { //dengan kembalian user.service
	return &patientService{
		patientRepository: repo,
		validate:          validator.New(),
	}
}

// Create implements patient.ServiceInterface
func (service *patientService) Create(input patient.CorePatient) (err error) {
	if validateERR := service.validate.Struct(input); validateERR != nil {
		return validateERR
	}

	errCreate := service.patientRepository.Create(input)
	if errCreate != nil {
		return errors.New(" Kesalahan pada input data pasien")
	}

	return nil
}

// DeleteById implements patient.ServiceInterface
func (*patientService) DeleteById(id int) error {
	panic("unimplemented")
}

// GetByPatientId implements patient.ServiceInterface
func (service *patientService) GetByPatientId(id int) (data patient.CorePatient, err error) {
	data, err = service.patientRepository.GetByPatientId(id) // memanggil struct entities repository yang ada di entities yang berisi coding logic
	return
}

// GetByUserId implements patient.ServiceInterface
func (service *patientService) GetByUserId(userid int) (data []patient.CorePatient, err error) {
	data, err = service.patientRepository.GetByUserId(userid) // memanggil struct entities repository yang ada di entities yang berisi coding logic
	return
}

// Update implements patient.ServiceInterface
func (*patientService) Update(id int, input patient.CorePatient) error {
	panic("unimplemented")
}
