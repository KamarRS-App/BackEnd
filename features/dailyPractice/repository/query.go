package repository

import (
	"errors"

	dailypractice "github.com/KamarRS-App/KamarRS-App/features/dailyPractice"
	"gorm.io/gorm"
)

type practiceRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) dailypractice.RepositoryInterface {
	return &practiceRepository{
		db: db,
	}
}

// Post
func (repo *practiceRepository) Create(input dailypractice.PracticeCore) (row int, err error) {
	practiceGorm := FromCore(input)
	tx := repo.db.Create(&practiceGorm) // proses insert data
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("insert failed, error query")
	}
	return int(tx.RowsAffected), nil
}

// Get by (ID)
func (repo *practiceRepository) GetById(id int) (data dailypractice.PracticeCore, err error) {
	var IdPractice Practice
	var IdPracticeCore = dailypractice.PracticeCore{}
	IdPractice.ID = uint(id)
	tx := repo.db.First(&IdPractice, IdPractice.ID)
	if tx.Error != nil {
		return IdPracticeCore, tx.Error
	}
	IdPracticeCore = IdPractice.ToCore()
	return IdPracticeCore, nil
}

// GetAll
func (repo *practiceRepository) GetAll(limit, offset, id int) (data []dailypractice.PracticeCore, totalpage int, err error) {
	var practices []Practice
	var count int64
	rx := repo.db.Model(&practices).Where("policlinic_id = ?", id).Count(&count)
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

	tx := repo.db.Find(&practices)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}
	var dataCore = ToCoreList(practices)
	return dataCore, totalpage, nil
}

// Update
func (repo *practiceRepository) Update(datacore dailypractice.PracticeCore, id int) (err error) {
	practiceGorm := FromCore(datacore)
	tx := repo.db.Where("id= ?", id).Updates(practiceGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("update dailypractice failed, error query")
	}
	return nil
}
