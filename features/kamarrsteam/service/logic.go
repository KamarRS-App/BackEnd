package service

import (
	"strings"

	"github.com/KamarRS-App/features/kamarrsteam"
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
		return errValidate
	}

	_, errFindEmail := s.kamarrsteamRepository.FindTeam(input.Email)

	if errFindEmail != nil && !strings.Contains(errFindEmail.Error(), "found") {
		return errFindEmail
	}

	hashPass, errEncrypt := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if errEncrypt != nil {
		log.Error(errEncrypt.Error())
		return errEncrypt
	}

	input.Password = string(hashPass)
	errCreate := s.kamarrsteamRepository.Create(input)
	if errCreate != nil {
		log.Error(errCreate.Error())
		return errCreate
	}

	return nil
}
