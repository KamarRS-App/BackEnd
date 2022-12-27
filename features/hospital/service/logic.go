package service

import (
	"errors"

	"github.com/KamarRS-App/KamarRS-App/features/hospital"
	"github.com/go-playground/validator/v10"
)

type hospitalService struct {
	hospitalRepository hospital.RepositoryInterface
	validate           *validator.Validate
}

func New(repo hospital.RepositoryInterface) hospital.ServiceInterface {
	return &hospitalService{
		hospitalRepository: repo,
		validate:           validator.New(),
	}
}

// Post
func (service *hospitalService) Create(input hospital.HospitalCore) (err error) {
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}
	_, errCreate := service.hospitalRepository.Create(input)
	if errCreate != nil {
		return errors.New("failed to insert data, error query")
	}
	return nil
}

// Get by (ID)
func (service *hospitalService) GetById(id int) (data hospital.HospitalCore, err error) {
	data, errGet := service.hospitalRepository.GetById(id)
	if errGet != nil {
		return data, errors.New("failed get hospital by id data, error query")
	}
	return data, nil
}

// GetAll
func (service *hospitalService) GetAll() (data []hospital.HospitalCore, err error) {
	data, err = service.hospitalRepository.GetAll()
	return

}

// Update
func (service *hospitalService) Update(dataCore hospital.HospitalCore, id int) (err error) {
	errUpdate := service.hospitalRepository.Update(dataCore, id)
	if errUpdate != nil {
		return errors.New("failed update hospital data, error query")
	}
	return nil

}

// Delete
func (service *hospitalService) Delete(id int) (err error) {
	_, errDel := service.hospitalRepository.Delete(id)
	if errDel != nil {
		return errors.New("failed delete hospital, error query")
	}
	return nil
}
