package repository

import (
	"errors"

	"github.com/KamarRS-App/KamarRS-App/features/dailypractice"
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
func (repo *practiceRepository) GetAll() (data []dailypractice.PracticeCore, err error) {
	var practices []Practice

	tx := repo.db.Find(&practices)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = ToCoreList(practices)
	return dataCore, nil
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
