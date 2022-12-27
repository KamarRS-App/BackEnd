package repository

import (
	"errors"
	"fmt"

	hospitalstaff "github.com/KamarRS-App/features/hospitalStaff"
	"github.com/KamarRS-App/features/hospitalStaff/service"
	"gorm.io/gorm"
)

type staffRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) hospitalstaff.RepositoryInterface { // user.repository mengimplementasikan interface repository yang ada di entities
	return &staffRepository{
		db: db,
	}

}

// Create implements hospitalstaff.RepositoryInterface
func (repo *staffRepository) Create(input hospitalstaff.HospitalStaffCore) (err error) {
	var staffs []HospitalStaff

	tx1 := repo.db.Find(&staffs)
	if tx1.Error != nil {
		return tx1.Error
	}

	for _, v := range staffs {
		if input.Email == v.Email {
			return errors.New("email sudah pernah terdaftar silahkan mendaftar dengan email yang lain")
		}

	}

	staffGorm := FromStaffCore(input)

	tx := repo.db.Create(&staffGorm) // proses insert data

	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}
	return nil
}

// DeleteById implements hospitalstaff.RepositoryInterface
func (repo *staffRepository) DeleteById(id int) error {
	var staff HospitalStaff

	tx1 := repo.db.Delete(&staff, id)
	if tx1.Error != nil {
		return tx1.Error
	}

	if tx1.RowsAffected == 0 {
		return errors.New("id not found")

	}

	return nil
}

// GetStaff implements hospitalstaff.RepositoryInterface
func (repo *staffRepository) GetStaff(id int) (data hospitalstaff.HospitalStaffCore, err error) {
	var staff HospitalStaff

	tx := repo.db.First(&staff, id)

	if tx.Error != nil {

		return hospitalstaff.HospitalStaffCore{}, tx.Error
	}
	gorms := staff.ModelsToCore()
	return gorms, nil
}

// Update implements hospitalstaff.RepositoryInterface
func (repo *staffRepository) Update(id int, input hospitalstaff.HospitalStaffCore) error {
	var staffs []HospitalStaff

	tx := repo.db.Find(&staffs)
	if tx.Error != nil {
		return tx.Error
	}

	for _, v := range staffs {
		if input.Email == v.Email {
			return errors.New("email sudah pernah terdaftar silahkan mendaftar dengan email yang lain")
		}

	}
	var staff HospitalStaff

	tx1 := repo.db.First(&staff, id)

	if tx1.Error != nil {

		return tx1.Error
	}

	if input.KataSandi == "" {
		input.KataSandi = staff.Kata_Sandi
	} else {
		input.KataSandi = service.Bcript(input.KataSandi)

	}
	userGorm := FromStaffCore(input)

	fmt.Println(input.KataSandi)
	tx2 := repo.db.Model(&userGorm).Where("id = ?", id).Updates(&userGorm)

	if tx2.Error != nil {
		return tx2.Error
	}

	return nil
}
