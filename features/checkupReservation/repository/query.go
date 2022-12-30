package repository

import (
	"errors"

	checkupreservation "github.com/KamarRS-App/KamarRS-App/features/checkupReservation"
	dailypractice "github.com/KamarRS-App/KamarRS-App/features/dailyPractice/repository"
	"github.com/KamarRS-App/KamarRS-App/features/patient/repository"
	userRep "github.com/KamarRS-App/KamarRS-App/features/user/repository"
	"gorm.io/gorm"
)

type CheckUpRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) checkupreservation.RepositoryInterface { // user.repository mengimplementasikan interface repository yang ada di entities
	return &CheckUpRepository{
		db: db,
	}

}

// Create implements checkupreservation.RepositoryInterface
func (repo *CheckUpRepository) Create(input checkupreservation.CheckupReservationCore, userId int) (users userRep.User, patients repository.Patient, practice dailypractice.Practice, err error) {
	patients = repository.Patient{}

	tx1 := repo.db.First(&patients, input.PatientID)
	if tx1.Error != nil {
		return userRep.User{}, repository.Patient{}, dailypractice.Practice{}, tx1.Error
	}
	users = userRep.User{}
	tx2 := repo.db.First(&users, userId)
	if tx2.Error != nil {
		return userRep.User{}, repository.Patient{}, dailypractice.Practice{}, tx2.Error
	}

	practice = dailypractice.Practice{}

	tx3 := repo.db.First(&practice, input.PracticeID)
	if tx3.Error != nil {
		return userRep.User{}, repository.Patient{}, dailypractice.Practice{}, tx3.Error
	}

	CheckUp := FromCoreToModel(input)

	tx := repo.db.Create(&CheckUp) // proses insert data

	if tx.Error != nil {
		return userRep.User{}, repository.Patient{}, dailypractice.Practice{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return userRep.User{}, repository.Patient{}, dailypractice.Practice{}, errors.New("insert failed")
	}

	practices := dailypractice.Practice{}
	practices.KuotaHarian = practice.KuotaHarian - 1
	tx4 := repo.db.Model(&practices).Where("id = ?", input.PatientID).Updates(&practices)
	if tx4.Error != nil {
		return userRep.User{}, repository.Patient{}, dailypractice.Practice{}, tx4.Error
	}
	return users, patients, practice, nil
}
