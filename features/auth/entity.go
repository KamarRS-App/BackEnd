package auth

import (
	teamrepo "github.com/KamarRS-App/KamarRS-App/features/kamarrsteam/repository"
	"github.com/KamarRS-App/KamarRS-App/features/user/repository"
)

type ServiceInterface interface {
	Login(email string, pass string) (string, repository.User, error)
	LoginTeam(email string, password string) (string, teamrepo.KamarRsTeam, error)
	// LoginOauth(email string) (string, repository.User, error)
	// Create(input user.CoreUser) (err error)
}

type RepositoryInterface interface {
	Login(email string, pass string) (string, repository.User, error)
	LoginTeam(email string, password string) (string, teamrepo.KamarRsTeam, error)
	// LoginOauth(email string) (string, repository.User, error)
	// Create(input user.CoreUser) (err error)
}
