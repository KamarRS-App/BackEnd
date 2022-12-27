package repository

import (
	"errors"

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
	tx := repo.db.First(&IdPoliclinic, IdPoliclinic.ID)
	if tx.Error != nil {
		return IdPoliclinicCore, tx.Error
	}
	IdPoliclinicCore = IdPoliclinic.ToCore()
	return IdPoliclinicCore, nil
}

// GetAll
func (repo *policlinicRepository) GetAll() (data []policlinic.CorePoliclinic, err error) {
	var policlinics []Policlinic

	tx := repo.db.Find(&policlinics)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = ToCoreList(policlinics)
	return dataCore, nil
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
