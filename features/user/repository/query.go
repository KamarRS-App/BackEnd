package repository

import (
	"errors"

	"github.com/KamarRS-App/features/user"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// Create implements user.RepositoryInterface
func (repo *userRepository) Create(input user.CoreUser) (err error) {
	userGorm := FromUserCoreToModel(input)

	tx := repo.db.Create(&userGorm) // proses insert data

	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}
	return nil
}

func New(db *gorm.DB) user.RepositoryInterface { // user.repository mengimplementasikan interface repository yang ada di entities
	return &userRepository{
		db: db,
	}

}
