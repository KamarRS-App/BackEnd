package service

import (
	"errors"

	"github.com/KamarRS-App/KamarRS-App/features/doctor"
	"github.com/go-playground/validator/v10"
)

type doctorService struct {
	doctorRepository doctor.RepositoryInterface
	validate         *validator.Validate
}

func New(repo doctor.RepositoryInterface) doctor.ServiceInterface {
	return &doctorService{
		doctorRepository: repo,
		validate:         validator.New(),
	}
}

// Post
func (service *doctorService) Create(input doctor.DoctorCore) (err error) {
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}
	_, errCreate := service.doctorRepository.Create(input)
	if errCreate != nil {
		return errors.New("failed to insert data, error logic")
	}
	return nil
}

// Get by (ID)
func (service *doctorService) GetById(id int) (data doctor.DoctorCore, err error) {
	data, errGet := service.doctorRepository.GetById(id)
	if errGet != nil {
		return data, errors.New("failed get doctor by id data, error logic")
	}
	return data, nil
}

// GetAll
func (service *doctorService) GetAll() (data []doctor.DoctorCore, err error) {
	data, err = service.doctorRepository.GetAll()
	return

}

// Update
func (service *doctorService) Update(dataCore doctor.DoctorCore, id int) (err error) {
	errUpdate := service.doctorRepository.Update(dataCore, id)
	if errUpdate != nil {
		return errors.New("failed update doctor data, error logic")
	}
	return nil

}

// Delete
func (service *doctorService) Delete(id int) (err error) {
	_, errDel := service.doctorRepository.Delete(id)
	if errDel != nil {
		return errors.New("failed delete doctor, error logic")
	}
	return nil
}
