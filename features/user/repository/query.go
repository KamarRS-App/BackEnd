package repository

import (
	"errors"

	"github.com/KamarRS-App/features/user"
	"github.com/KamarRS-App/features/user/service"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.RepositoryInterface { // user.repository mengimplementasikan interface repository yang ada di entities
	return &userRepository{
		db: db,
	}

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

// Update implements user.RepositoryInterface
func (repo *userRepository) Update(id int, input user.CoreUser) error {
	var users User

	tx1 := repo.db.First(&users, id)

	if tx1.Error != nil {

		return tx1.Error
	}

	if input.KataSandi == "" {
		input.KataSandi = users.Kata_Sandi
	}

	userGorm := FromUserCoreToModel(input)
	input.KataSandi = service.Bcript(input.KataSandi)

	tx := repo.db.Model(&userGorm).Where("id = ?", id).Updates(&userGorm)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
