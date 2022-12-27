package repository

import (
	"errors"

	"github.com/KamarRS-App/KamarRS-App/features/kamarrsteam"
	"gorm.io/gorm"
)

type kamarrsteamRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) kamarrsteam.RepositoryInterface {
	return &kamarrsteamRepository{
		db: db,
	}
}

// Create implements kamarrsteam.RepositoryInterface
func (r *kamarrsteamRepository) Create(input kamarrsteam.KamarRsTeamCore) error {
	teamGorm := FromKamarRsTeamCoretoModel(input)
	tx := r.db.Create(&teamGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}
	return nil
}

// FindTeam implements kamarrsteam.RepositoryInterface
func (r *kamarrsteamRepository) FindTeam(email string) (result kamarrsteam.KamarRsTeamCore, err error) {
	var TeamData KamarRsTeam
	tx := r.db.Where("email = ?", email).Find(&TeamData)
	if tx.Error != nil {
		return kamarrsteam.KamarRsTeamCore{}, tx.Error
	}

	result = TeamData.ToKamarRsTeamCore()

	return result, nil
}
