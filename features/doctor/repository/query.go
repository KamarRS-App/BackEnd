package repository

import (
	"errors"

	"github.com/KamarRS-App/KamarRS-App/features/doctor"
	"gorm.io/gorm"
)

type doctorRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) doctor.RepositoryInterface {
	return &doctorRepository{
		db: db,
	}
}

// Post
func (repo *doctorRepository) Create(input doctor.DoctorCore) (row int, err error) {
	doctorGorm := FromCore(input)
	tx := repo.db.Create(&doctorGorm) // proses insert data
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("insert failed, error query")
	}
	return int(tx.RowsAffected), nil
}

// Get by (ID)
func (repo *doctorRepository) GetById(id int) (data doctor.DoctorCore, err error) {
	var IdDoctor Doctor
	var IdDoctorCore = doctor.DoctorCore{}
	IdDoctor.ID = uint(id)
	tx := repo.db.First(&IdDoctor, IdDoctor.ID)
	if tx.Error != nil {
		return IdDoctorCore, tx.Error
	}
	IdDoctorCore = IdDoctor.ToCore()
	return IdDoctorCore, nil
}

// GetAll
func (repo *doctorRepository) GetAll() (data []doctor.DoctorCore, err error) {
	var doctors []Doctor

	tx := repo.db.Find(&doctors)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = ToCoreList(doctors)
	return dataCore, nil
}

// Update
func (repo *doctorRepository) Update(datacore doctor.DoctorCore, id int) (err error) {
	doctorGorm := FromCore(datacore)
	tx := repo.db.Where("id= ?", id).Updates(doctorGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("update doctor failed, error query")
	}
	return nil
}

// Delete
func (repo *doctorRepository) Delete(id int) (row int, err error) {
	IdDoctor := Doctor{}

	tx := repo.db.Delete(&IdDoctor, id)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return -1, errors.New("delete doctor by id failed, error query")
	}
	return int(tx.RowsAffected), nil
}
