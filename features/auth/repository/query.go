package repository

import (
	"errors"

	"github.com/KamarRS-App/KamarRS-App/features/auth"
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

	token, errToken := middlewares.CreateToken(int(userData.ID), "")
	if errToken != nil {
		return "", repository.User{}, errToken
	}

	return token, userData, nil
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

	token, errToken := middlewares.CreateToken(int(staffs.ID), staffs.Peran)
	if errToken != nil {
		return "", staff.HospitalStaff{}, errToken
	}

	return token, staffs, nil
}
