package repository

import (
	"errors"
	"fmt"

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
func (repo *CheckUpRepository) Create(input checkupreservation.CheckupReservationCore, userId int) (err error) {
	patients := repository.Patient{}

	tx1 := repo.db.First(&patients, input.PatientID)
	if tx1.Error != nil {
		return tx1.Error
	}
	users := userRep.User{}
	tx2 := repo.db.First(&users, userId)
	if tx2.Error != nil {
		return tx2.Error
	}

	practice := dailypractice.Practice{}

	tx3 := repo.db.First(&practice, input.PracticeID)
	if tx3.Error != nil {
		return tx3.Error
	}

	if practice.KuotaHarian == -1 {
		practice.Status = "Not Available"
		tx6 := repo.db.Model(&practice).Where("id = ?", input.PatientID).Updates(&practice)
		if tx6.Error != nil {
			return tx6.Error
		}
		return errors.New("kuota habis")

	}
	if users.Nokk != patients.NoKk {
		return errors.New("no kk salah")
	}

	CheckUp := FromCoreToModel(input)

	tx := repo.db.Create(&CheckUp) // proses insert data

	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}

	practices := dailypractice.Practice{}
	fmt.Println(practice.KuotaHarian)
	practices.KuotaHarian = practice.KuotaHarian - 1
	if practices.KuotaHarian == 0 {
		practices.KuotaHarian = -1
	}
	fmt.Println(practices.KuotaHarian)
	tx4 := repo.db.Model(&practices).Where("id = ?", input.PatientID).Updates(&practices)
	if tx4.Error != nil {
		return tx4.Error
	}
	return nil
}
