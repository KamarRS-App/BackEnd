package repository

import (
	"errors"

	"github.com/KamarRS-App/KamarRS-App/features/bed"
	"gorm.io/gorm"
)

type bedRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) bed.RepositoryInterface {
	return &bedRepository{
		db: db,
	}
}

// Post
func (repo *bedRepository) Create(input bed.BedCore) (row int, err error) {
	bedGorm := FromCore(input)
	tx := repo.db.Create(&bedGorm) // proses insert data
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("insert failed, error query")
	}
	return int(tx.RowsAffected), nil
}

// Get by (ID)
func (repo *bedRepository) GetById(id int) (data bed.BedCore, err error) {
	var IdBed Bed
	var IdBedCore = bed.BedCore{}
	IdBed.ID = uint(id)
	tx := repo.db.First(&IdBed, IdBed.ID)
	if tx.Error != nil {
		return IdBedCore, tx.Error
	}
	IdBedCore = IdBed.ToCore()
	return IdBedCore, nil
}

// GetAll
func (repo *bedRepository) GetAll() (data []bed.BedCore, err error) {
	var beds []Bed

	tx := repo.db.Find(&beds)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = ToCoreList(beds)
	return dataCore, nil
}

// Update
func (repo *bedRepository) Update(datacore bed.BedCore, id int) (err error) {
	bedGorm := FromCore(datacore)
	tx := repo.db.Where("id= ?", id).Updates(bedGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("update bed failed, error query")
	}
	return nil
}

// Delete
func (repo *bedRepository) Delete(id int) (row int, err error) {
	IdBed := Bed{}

	tx := repo.db.Delete(&IdBed, id)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return -1, errors.New("delete bed by id failed, error query")
	}
	return int(tx.RowsAffected), nil
}
