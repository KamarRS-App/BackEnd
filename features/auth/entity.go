package auth

import "github.com/KamarRS-App/KamarRS-App/features/user/repository"

type ServiceInterface interface {
	Login(email string, pass string) (string, repository.User, error)
	// LoginOauth(email string) (string, repository.User, error)
	// Create(input user.CoreUser) (err error)
}

type RepositoryInterface interface {
	Login(email string, pass string) (string, repository.User, error)
	// LoginOauth(email string) (string, repository.User, error)
	// Create(input user.CoreUser) (err error)
}
