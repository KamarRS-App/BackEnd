package repository

import (
	"errors"

	"github.com/KamarRS-App/features/hospital"
	"gorm.io/gorm"
)

type hospitalRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) hospital.RepositoryInterface {
	return &hospitalRepository{
		db: db,
	}
}

// Post
func (repo *hospitalRepository) Create(input hospital.HospitalCore) (row int, err error) {
	hospitalGorm := FromCore(input)
	tx := repo.db.Create(&hospitalGorm) // proses insert data
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("insert failed, error query")
	}
	return int(tx.RowsAffected), nil
}

// Get by (ID)
func (repo *hospitalRepository) GetById(id int) (data hospital.HospitalCore, err error) {
	var IdHospital Hospital
	var IdHospitalCore = hospital.HospitalCore{}
	IdHospital.ID = uint(id)
	tx := repo.db.First(&IdHospital, IdHospital.ID)
	if tx.Error != nil {
		return IdHospitalCore, tx.Error
	}
	IdHospitalCore = IdHospital.toCore()
	return IdHospitalCore, nil
}

// GetAll
func (repo *hospitalRepository) GetAll() (data []hospital.HospitalCore, err error) {
	var hospitals []Hospital

	tx := repo.db.Find(&hospitals)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(hospitals)
	return dataCore, nil
}

// Update
func (repo *hospitalRepository) Update(datacore hospital.HospitalCore, id int) (err error) {
	hospitalGorm := FromCore(datacore)
	tx := repo.db.Where("id= ?", id).Updates(hospitalGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("update hospital failed, error query")
	}
	return nil
}

// Delete
func (repo *hospitalRepository) Delete(id int) (row int, err error) {
	IdHospital := Hospital{}

	tx := repo.db.Delete(&IdHospital, id)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return -1, errors.New("delete hospital by id failed, error query")
	}
	return int(tx.RowsAffected), nil
}
