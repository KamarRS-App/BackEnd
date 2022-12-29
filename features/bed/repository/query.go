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

// GetAll
func (repo *bedRepository) GetAll(limit, offset, id int) (data []bed.BedCore, totalpage int, err error) {
	var beds []Bed
	var count int64
	rx := repo.db.Model(&Bed{}).Where("hospital_id = ?", id).Count(&count)
	if rx.Error != nil {
		return nil, 0, rx.Error
	}
	if rx.RowsAffected == 0 {
		return nil, 0, errors.New("error query count")
	}

	// var totalpage int
	if int(count)%limit == 0 {
		totalpage = int(count) / limit
	} else {
		totalpage = (int(count) / limit) + 1
	}

	tx := repo.db.Where("hospital_id = ?", id).Limit(limit).Offset(offset).Find(&beds)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, 0, errors.New("get all data failed, error query data")
	}
	var dataCore = ToCoreList(beds)
	return dataCore, totalpage, nil
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
