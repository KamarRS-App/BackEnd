package service

import (
	"errors"

	"github.com/KamarRS-App/KamarRS-App/features/auth"

	staff "github.com/KamarRS-App/KamarRS-App/features/hospitalstaff/repository"
	teamrepo "github.com/KamarRS-App/KamarRS-App/features/kamarrsteam/repository"

	"github.com/KamarRS-App/KamarRS-App/features/user/repository"
	// "github.com/KamarRS-App/KamarRS-App/utils/helper"
)

type authService struct {
	authRepository auth.RepositoryInterface
}

func New(repo auth.RepositoryInterface) auth.ServiceInterface {
	return &authService{
		authRepository: repo,
	}
}

// Login implements auth.ServiceInterface
func (service *authService) Login(email string, pass string) (string, repository.User, error) {
	token, data, err := service.authRepository.Login(email, pass)
	if err != nil {
		return "", repository.User{}, errors.New("Login Failed, incorrect input")
	}
	return token, data, err
}

// LoginTeam implements auth.ServiceInterface
func (s *authService) LoginTeam(email string, password string) (string, teamrepo.KamarRsTeam, error) {
	if email == "" || password == "" {
		return "", teamrepo.KamarRsTeam{}, errors.New("field must be filled")
	}

	token, data, err := s.authRepository.LoginTeam(email, password)
	if err != nil {
		return "", teamrepo.KamarRsTeam{}, err
	}

	// passCheck := helper.CheckPasswordHash(password, data.Password)
	// if !passCheck {
	// 	return "", teamrepo.KamarRsTeam{}, errors.New("login failed")
	// }

	return token, data, nil
}

// LoginStaff implements auth.ServiceInterface
func (service *authService) LoginStaff(email string, pass string) (string, staff.HospitalStaff, error) {
	token, data, err := service.authRepository.LoginStaff(email, pass)
	if err != nil {
		return "", staff.HospitalStaff{}, errors.New("Login Failed, incorrect input")
	}
	return token, data, err

}

// LoginOauth implements auth.ServiceInterface
func (service *authService) LoginOauth(auths auth.Oauth) (string, repository.User, error) {
	token, data, err := service.authRepository.LoginOauth(auths)
	if err != nil {
		return "", repository.User{}, errors.New("Login Failed, incorrect input")
	}
	return token, data, err
}
