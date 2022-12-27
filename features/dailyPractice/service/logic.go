package service

import (
	"errors"

	dailypractice "github.com/KamarRS-App/KamarRS-App/features/dailyPractice"
	"github.com/go-playground/validator/v10"
)

type practiceService struct {
	practiceRepository dailypractice.RepositoryInterface
	validate           *validator.Validate
}

func New(repo dailypractice.RepositoryInterface) dailypractice.ServiceInterface {
	return &practiceService{
		practiceRepository: repo,
		validate:           validator.New(),
	}
}

// Post
func (service *practiceService) Create(input dailypractice.PracticeCore) (err error) {
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}
	_, errCreate := service.practiceRepository.Create(input)
	if errCreate != nil {
		return errors.New("failed to insert data, error logic")
	}
	return nil
}

// Get by (ID)
func (service *practiceService) GetById(id int) (data dailypractice.PracticeCore, err error) {
	data, errGet := service.practiceRepository.GetById(id)
	if errGet != nil {
		return data, errors.New("failed get practice by id data, error logic")
	}
	return data, nil
}

// GetAll
func (service *practiceService) GetAll() (data []dailypractice.PracticeCore, err error) {
	data, err = service.practiceRepository.GetAll()
	return

}

// Update
func (service *practiceService) Update(dataCore dailypractice.PracticeCore, id int) (err error) {
	errUpdate := service.practiceRepository.Update(dataCore, id)
	if errUpdate != nil {
		return errors.New("failed update practice data, error logic")
	}
	return nil

}
