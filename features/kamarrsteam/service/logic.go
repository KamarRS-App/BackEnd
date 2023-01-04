package service

import (
	"errors"

	"github.com/KamarRS-App/KamarRS-App/features/kamarrsteam"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

type kamarrsteamService struct {
	kamarrsteamRepository kamarrsteam.RepositoryInterface
	validate              *validator.Validate
}

func New(repo kamarrsteam.RepositoryInterface) kamarrsteam.ServiceInterface {
	return &kamarrsteamService{
		kamarrsteamRepository: repo,
		validate:              validator.New(),
	}
}

// Create implements kamarrsteam.ServiceInterface
func (s *kamarrsteamService) Create(input kamarrsteam.KamarRsTeamCore) error {
	errValidate := s.validate.Struct(input)
	if errValidate != nil {
		log.Error(errValidate.Error())
		return errValidate
	}

	foundData, _ := s.kamarrsteamRepository.FindTeam(input.Email)

	if foundData.Email == input.Email {
		return errors.New("use another email")
	}

	hashPass, errEncrypt := bcrypt.GenerateFromPassword([]byte(input.KataSandi), 10)
	if errEncrypt != nil {
		return errEncrypt
	}

	input.KataSandi = string(hashPass)
	errCreate := s.kamarrsteamRepository.Create(input)
	if errCreate != nil {
		log.Error(errCreate.Error())
		return errCreate
	}

	return nil
}
