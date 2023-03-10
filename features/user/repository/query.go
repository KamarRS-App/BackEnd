package repository

import (
	"errors"
	"fmt"

	"github.com/KamarRS-App/KamarRS-App/features/user"
	"github.com/KamarRS-App/KamarRS-App/utils/helper"

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
	// var users []User

	// tx1 := repo.db.Find(&users)
	// if tx1.Error != nil {
	// 	return tx1.Error
	// }

	// // for _, v := range users {
	// // 	if input.Email == v.Email {
	// // 		return errors.New("email sudah pernah terdaftar silahkan mendaftar dengan email yang lain")
	generatePass := helper.Bcript(input.KataSandi)

	input.KataSandi = generatePass

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
	// var user []User

	// tx := repo.db.Find(&user)
	// if tx.Error != nil {
	// 	return tx.Error
	// }

	// for _, v := range user {
	// 	if input.Email == v.Email {
	// 		return errors.New("email sudah pernah terdaftar silahkan mendaftar dengan email yang lain")
	// 	}

	// }
	var users User

	tx1 := repo.db.First(&users, id)

	if tx1.Error != nil {

		return tx1.Error
	}

	if input.KataSandi == "" {
		input.KataSandi = users.KataSandi
	} else {
		input.KataSandi = helper.Bcript(input.KataSandi)

	}
	userGorm := FromUserCoreToModel(input)

	fmt.Println(input.KataSandi)
	tx2 := repo.db.Model(&userGorm).Where("id = ?", id).Updates(&userGorm)

	if tx2.Error != nil {
		return tx2.Error
	}

	return nil
}

// GetById implements user.RepositoryInterface
func (repo *userRepository) GetById(id int) (data user.CoreUser, err error) {
	var users User

	tx := repo.db.First(&users, id)

	if tx.Error != nil {

		return user.CoreUser{}, tx.Error
	}
	gorms := users.ModelsToCore()
	return gorms, nil
}

// DeleteById implements user.RepositoryInterface
func (repo *userRepository) DeleteById(id int) error {
	users := User{}

	tx1 := repo.db.Unscoped().Delete(&users, id)
	if tx1.Error != nil {
		return tx1.Error
	}

	if tx1.RowsAffected == 0 {
		return errors.New("id not found")

	}

	return nil
}
