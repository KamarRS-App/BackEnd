package auth

import (

	teamrepo "github.com/KamarRS-App/KamarRS-App/features/kamarrsteam/repository"
  staff "github.com/KamarRS-App/KamarRS-App/features/hospitalstaff/repository"
	user "github.com/KamarRS-App/KamarRS-App/features/user/repository"
	"github.com/KamarRS-App/KamarRS-App/features/user/repository"
)

type ServiceInterface interface {
	Login(email string, pass string) (string, user.User, error)
	LoginTeam(email string, password string) (string, teamrepo.KamarRsTeam, error)
  LoginStaff(email string, pass string) (string, staff.HospitalStaff, error)

	// LoginOauth(email string) (string, repository.User, error)
	// Create(input user.CoreUser) (err error)
}

type RepositoryInterface interface {

	Login(email string, pass string) (string, user.User, error)
	LoginTeam(email string, password string) (string, teamrepo.KamarRsTeam, error)
  LoginStaff(email string, pass string) (string, staff.HospitalStaff, error)

	// LoginOauth(email string) (string, repository.User, error)
	// Create(input user.CoreUser) (err error)
}
