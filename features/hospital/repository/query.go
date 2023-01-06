package repository

import (
	"errors"

	"github.com/KamarRS-App/KamarRS-App/features/hospital"
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
func (repo *hospitalRepository) Create(input hospital.HospitalCore) (err error) {
	hospitalGorm := FromCore(input)
	tx := repo.db.Create(&hospitalGorm) // proses insert data
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed, error query")
	}
	return nil
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
	IdHospitalCore = IdHospital.ToCore()
	return IdHospitalCore, nil
}

// GetAll
func (repo *hospitalRepository) GetAll(provinsi, kabKota, nama string, limit, offset int) (data []hospital.HospitalCore, totalPage int, err error) {
	var hospitals []Hospital

	var count int64
	tx0 := repo.db.Model(&hospitals).Where("provinsi LIKE ? AND kabupaten_kota LIKE ? AND nama LIKE ?", "%"+provinsi+"%", "%"+kabKota+"%", "%"+nama+"%").Count(&count)
	if tx0.Error != nil {
		return nil, 0, tx0.Error
	}
	if tx0.RowsAffected == 0 {
		return nil, 0, errors.New("error query count")
	}

	if count < 10 {
		totalPage = 1
	} else if int(count)%limit == 0 {
		totalPage = int(count) / limit
	} else {
		totalPage = (int(count) / limit) + 1
	}

	tx := repo.db.Where("provinsi LIKE ? AND kabupaten_kota LIKE ? AND nama LIKE ?", "%"+provinsi+"%", "%"+kabKota+"%", "%"+nama+"%").Limit(limit).Offset(offset).Find(&hospitals)

	if tx.Error != nil {
		return nil, 0, tx.Error
	}
	var dataCore = ToCoreList(hospitals)
	return dataCore, totalPage, nil
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
