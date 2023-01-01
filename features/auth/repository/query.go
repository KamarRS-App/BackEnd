package repository

import (
	"errors"

	"github.com/KamarRS-App/KamarRS-App/features/auth"
	teamrepo "github.com/KamarRS-App/KamarRS-App/features/kamarrsteam/repository"

	staff "github.com/KamarRS-App/KamarRS-App/features/hospitalstaff/repository"

	"github.com/KamarRS-App/KamarRS-App/features/user/repository"
	middlewares "github.com/KamarRS-App/KamarRS-App/middlewares"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) auth.RepositoryInterface {
	return &authRepository{
		db: db,
	}
}

// Login implements auth.RepositoryInterface
func (repo *authRepository) Login(email string, pass string) (string, repository.User, error) {
	var userData repository.User
	tx := repo.db.Where("email = ?", email).First(&userData)
	if tx.Error != nil {
		return "", repository.User{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return "", repository.User{}, errors.New("login failed")
	}

	token, errToken := middlewares.CreateTokenTeam(int(userData.ID), "", "")
	if errToken != nil {
		return "", repository.User{}, errToken
	}

	return token, userData, nil
}

// LoginTeam implements auth.RepositoryInterface
func (r *authRepository) LoginTeam(email string, password string) (string, teamrepo.KamarRsTeam, error) {
	var teamData teamrepo.KamarRsTeam
	tx := r.db.Where("email = ?", email).First(&teamData)
	if tx.Error != nil {
		return "", teamrepo.KamarRsTeam{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return "", teamrepo.KamarRsTeam{}, errors.New("logifn failed")
	}

	token, errToken := middlewares.CreateTokenTeam(int(teamData.ID), teamData.Peran, teamData.Email)
	if errToken != nil {
		return "", teamrepo.KamarRsTeam{}, errToken
	}

	return token, teamData, nil
}

// LoginStaff implements auth.RepositoryInterface
func (repo *authRepository) LoginStaff(email string, pass string) (string, staff.HospitalStaff, error) {
	var staffs staff.HospitalStaff
	tx := repo.db.Where("email = ?", email).First(&staffs)
	if tx.Error != nil {
		return "", staff.HospitalStaff{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return "", staff.HospitalStaff{}, errors.New("login failed")
	}

	token, errToken := middlewares.CreateTokenTeam(int(staffs.ID), staffs.Peran, "")
	if errToken != nil {
		return "", staff.HospitalStaff{}, errToken
	}

	return token, staffs, nil

}

// LoginOauth implements auth.RepositoryInterface
func (repo *authRepository) LoginOauth(auths auth.Oauth) (string, repository.User, error) {
	var userData repository.User

	tx := repo.db.Where("email = ?", auths.Email).First(&userData)
	user := repository.User{}
	user.Email = auths.Email
	user.Nama = auths.Name

	if tx.Error != nil {

		tx1 := repo.db.Create(&user) // proses insert data

		if tx1.Error != nil {
			return "", repository.User{}, tx1.Error
		}
		if tx1.RowsAffected == 0 {
			return "", repository.User{}, errors.New("insert failed")
		}

	}

	tx3 := repo.db.Where("email = ?", auths.Email).First(&userData)
	if tx3.Error != nil {
		return "", repository.User{}, tx3.Error
	}

	// if tx.RowsAffected == 0 {
	// 	return "", repository.User{}, errors.New("login failed")
	// }

	token, errToken := middlewares.CreateTokenTeam(int(userData.ID), "", "")
	if errToken != nil {
		return "", repository.User{}, errToken
	}

	return token, userData, nil
}
