package service

import (
	"errors"

	"github.com/KamarRS-App/features/policlinic"
	"github.com/go-playground/validator/v10"
)

type policlinicService struct {
	policlinicRepository policlinic.RepositoryInterface
	validate             *validator.Validate
}

func New(repo policlinic.RepositoryInterface) policlinic.ServiceInterface {
	return &policlinicService{
		policlinicRepository: repo,
		validate:             validator.New(),
	}
}

// Post
func (service *policlinicService) Create(input policlinic.CorePoliclinic) (err error) {
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}
	_, errCreate := service.policlinicRepository.Create(input)
	if errCreate != nil {
		return errors.New("failed to insert data, error logic")
	}
	return nil
}

// Get by (ID)
func (service *policlinicService) GetById(id int) (data policlinic.CorePoliclinic, err error) {
	data, errGet := service.policlinicRepository.GetById(id)
	if errGet != nil {
		return data, errors.New("failed get policlinic by id data, error logic")
	}
	return data, nil
}

// GetAll
func (service *policlinicService) GetAll() (data []policlinic.CorePoliclinic, err error) {
	data, err = service.policlinicRepository.GetAll()
	return

}

// Update
func (service *policlinicService) Update(dataCore policlinic.CorePoliclinic, id int) (err error) {
	errUpdate := service.policlinicRepository.Update(dataCore, id)
	if errUpdate != nil {
		return errors.New("failed update policlinic data, error logic")
	}
	return nil

}

// Delete
func (service *policlinicService) Delete(id int) (err error) {
	_, errDel := service.policlinicRepository.Delete(id)
	if errDel != nil {
		return errors.New("failed delete policlinic, error logic")
	}
	return nil
}
