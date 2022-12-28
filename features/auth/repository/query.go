package repository

import (
	"errors"

	"github.com/KamarRS-App/KamarRS-App/features/auth"
	teamrepo "github.com/KamarRS-App/KamarRS-App/features/kamarrsteam/repository"
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

	token, errToken := middlewares.CreateToken(int(userData.ID), userData.Nama)
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
