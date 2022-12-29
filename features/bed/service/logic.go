package service

import (
	"errors"

	"github.com/KamarRS-App/KamarRS-App/features/bed"
	"github.com/go-playground/validator/v10"
)

type bedService struct {
	bedRepository bed.RepositoryInterface
	validate      *validator.Validate
}

func New(repo bed.RepositoryInterface) bed.ServiceInterface {
	return &bedService{
		bedRepository: repo,
		validate:      validator.New(),
	}
}

// Post
func (service *bedService) Create(input bed.BedCore) (err error) {
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}
	_, errCreate := service.bedRepository.Create(input)
	if errCreate != nil {
		return errors.New("failed to insert data, error logic")
	}
	return nil
}

// GetAll
func (service *bedService) GetAll(pagination, limit, id int) (data []bed.BedCore, totalpage int, err error) {
	offset := (pagination - 1) * limit
	data, totalpage, err = service.bedRepository.GetAll(limit, offset, id)
	if err != nil {
		return nil, 0, errors.New("failed get bed by hospital id, error logic")
	}
	return

}

// Get by (ID)
func (service *bedService) GetById(id int) (data bed.BedCore, err error) {
	data, errGet := service.bedRepository.GetById(id)
	if errGet != nil {
		return data, errors.New("failed get bed by id data, error logic")
	}
	return data, nil
}

// Update
func (service *bedService) Update(dataCore bed.BedCore, id int) (err error) {
	errUpdate := service.bedRepository.Update(dataCore, id)
	if errUpdate != nil {
		return errors.New("failed update bed data, error logic")
	}
	return nil

}

// Delete
func (service *bedService) Delete(id int) (err error) {
	_, errDel := service.bedRepository.Delete(id)
	if errDel != nil {
		return errors.New("failed delete bed, error logic")
	}
	return nil
}
