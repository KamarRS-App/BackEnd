package repository

import (
	"errors"
	"fmt"

	"github.com/KamarRS-App/KamarRS-App/features/policlinic"
	"gorm.io/gorm"
)

type policlinicRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) policlinic.RepositoryInterface {
	return &policlinicRepository{
		db: db,
	}
}

// Post
func (repo *policlinicRepository) Create(input policlinic.CorePoliclinic) (row int, err error) {
	policlinicGorm := FromCore(input)
	tx := repo.db.Create(&policlinicGorm) // proses insert data
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("insert failed, error query")
	}
	return int(tx.RowsAffected), nil
}

// Get by (ID)
func (repo *policlinicRepository) GetById(id int) (data policlinic.CorePoliclinic, err error) {
	var IdPoliclinic Policlinic
	var IdPoliclinicCore = policlinic.CorePoliclinic{}
	IdPoliclinic.ID = uint(id)
	tx := repo.db.Preload("Doctors").First(&IdPoliclinic, IdPoliclinic.ID)
	if tx.Error != nil {
		return IdPoliclinicCore, tx.Error
	}
	IdPoliclinicCore = IdPoliclinic.ToCore()
	return IdPoliclinicCore, nil
}

// GetAll
func (repo *policlinicRepository) GetAll() (data []policlinic.CorePoliclinic, err error) {
	var policlinics []Policlinic

	tx := repo.db.Preload("Doctors").Find(&policlinics)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = ToCoreList(policlinics)
	return dataCore, nil
}

// GetAll
func (repo *policlinicRepository) GetAllbyHospitalID(limit, offset, id int) (data []policlinic.CorePoliclinic, totalpage int, err error) {
	var policlinics []Policlinic
	var count int64
	rx := repo.db.Model(&policlinics).Where("hospital_id = ?", id).Count(&count)
	if rx.Error != nil {
		return nil, 0, rx.Error
	}
	if rx.RowsAffected == 0 {
		return nil, 0, errors.New("error query count")
	}
	fmt.Println("+++++++++RX KING+++++++++", rx)
	fmt.Println("======COOUUUNNTTTTT=====", count)
	fmt.Println("#######ROWS AFFECTED#######", rx.RowsAffected)

	// var totalpage int
	if rx.RowsAffected < 10 {
		totalpage = 1
	} else if int(rx.RowsAffected)%limit == 0 {
		totalpage = int(rx.RowsAffected) / limit
	} else {
		totalpage = (int(rx.RowsAffected) / limit) + 1
	}

	tx := repo.db.Where("hospital_id = ?", id).Limit(limit).Offset(offset).Find(&policlinics)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, 0, errors.New("get all data failed, error query data")
	}
	data = ToCoreList(policlinics)
	return data, totalpage, nil
}

// Update
func (repo *policlinicRepository) Update(datacore policlinic.CorePoliclinic, id int) (err error) {
	policlinicGorm := FromCore(datacore)
	tx := repo.db.Where("id= ?", id).Updates(policlinicGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("update policlinic failed, error query")
	}
	return nil
}

// Delete
func (repo *policlinicRepository) Delete(id int) (row int, err error) {
	IdPoliclinic := Policlinic{}

	tx := repo.db.Delete(&IdPoliclinic, id)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return -1, errors.New("delete policlinic by id failed, error query")
	}
	return int(tx.RowsAffected), nil
}
