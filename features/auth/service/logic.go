package service

import (
	"github.com/KamarRS-App/KamarRS-App/features/auth"
	staff "github.com/KamarRS-App/KamarRS-App/features/hospitalstaff/repository"
	"github.com/KamarRS-App/KamarRS-App/features/user/repository"
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
	return token, data, err
}

// LoginStaff implements auth.ServiceInterface
func (service *authService) LoginStaff(email string, pass string) (string, staff.HospitalStaff, error) {
	token, data, err := service.authRepository.LoginStaff(email, pass)
	return token, data, err
}
