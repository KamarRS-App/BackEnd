package service

import (
	"github.com/KamarRS-App/features/auth"
	"github.com/KamarRS-App/features/user/repository"
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
