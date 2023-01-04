package repository

import (
	"errors"
	"fmt"
	"log"

	checkupreservation "github.com/KamarRS-App/KamarRS-App/features/checkupReservation"
	dailypractice "github.com/KamarRS-App/KamarRS-App/features/dailyPractice/repository"
	hospital "github.com/KamarRS-App/KamarRS-App/features/hospital/repository"
	patient "github.com/KamarRS-App/KamarRS-App/features/patient/repository"
	policlinic "github.com/KamarRS-App/KamarRS-App/features/policlinic/repository"

	userRep "github.com/KamarRS-App/KamarRS-App/features/user/repository"
	"github.com/KamarRS-App/KamarRS-App/utils/helper"
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
	patients := userRep.Patient{}

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
		tx6 := repo.db.Model(&practice).Where("id = ?", input.PracticeID).Updates(&practice)
		if tx6.Error == nil {
			fmt.Println("kuota habis")
			return errors.New("kuota habis")
		}

	}
	if users.Nokk != patients.NoKk {
		fmt.Println("no kk salah")
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

	dataCheckup := CheckupReservation{}
	tx9 := repo.db.Last(&dataCheckup)
	if tx9.Error != nil {
		return tx9.Error
	}

	dataPatients := patient.Patient{}
	tx5 := repo.db.First(&dataPatients, input.PatientID)
	if tx5.Error != nil {
		return tx5.Error
	}

	dataPractice := dailypractice.Practice{}
	tx6 := repo.db.Where("id = ?", input.PracticeID).Preload("Policlinic").Find(&dataPractice)
	if tx6.Error != nil {
		return tx6.Error
	}

	dataPoliclinic := policlinic.Policlinic{}
	tx7 := repo.db.Where("id = ?", dataPractice.PoliclinicID).Preload("Hospital").Find(&dataPoliclinic)
	if tx7.Error != nil {
		return tx7.Error
	}

	dataHospital := hospital.Hospital{}
	tx8 := repo.db.Where("id = ?", dataPoliclinic.HospitalID).Find(&dataHospital)
	if tx8.Error != nil {
		return tx8.Error
	}

	Appointment, errAp := helper.Calendar(dataPatients.EmailWali, dataPractice.TanggalPraktik, dataHospital.Alamat)
	if errAp != nil {
		fmt.Println("gagal menambahkan ke kalender")
		return errors.New("gagal menambahkan ke kalender")
	}
	fmt.Println("link", Appointment)

	dataEmail := struct {
		RumahSakit     string
		Policlinic     string
		JamPraktik     string
		TanggalPraktik string
		NamaPasien     string
		NamaDokter     string
		NoAntri        string
	}{
		RumahSakit:     dataHospital.Nama,
		Policlinic:     dataPoliclinic.NamaPoli,
		JamPraktik:     dataPoliclinic.JamPraktik,
		TanggalPraktik: dataPractice.TanggalPraktik,
		NamaPasien:     dataPatients.NamaPasien,
		NamaDokter:     dataCheckup.NamaDokter,
		NoAntri:        dataCheckup.NoAntrian,
	}

	emailTo := dataPatients.EmailWali

	errMail := helper.SendEmailSMTPCheckup([]string{emailTo}, dataEmail, "emailCheckup.txt") //send mail
	if errMail != nil {
		log.Println(errMail, "Pengiriman Email Gagal")
	}

	practices := dailypractice.Practice{}
	fmt.Println(practice.KuotaHarian)
	practices.KuotaHarian = practice.KuotaHarian - 1
	if practices.KuotaHarian == 0 {
		practices.KuotaHarian = -1
	}
	fmt.Println(practices.KuotaHarian)
	tx4 := repo.db.Model(&practices).Where("id = ?", input.PracticeID).Updates(&practices)
	if tx4.Error != nil {
		return tx4.Error
	}
	return nil
}

// GetByPracticesId implements checkupreservation.RepositoryInterface
func (repo *CheckUpRepository) GetByPracticesId(limit int, offset int, id int) (data []checkupreservation.CheckupReservationCore, totalpage int, err error) {
	var check []CheckupReservation
	var count int64
	rx := repo.db.Model(&check).Where("practice_id", id).Count(&count)
	if rx.Error != nil {
		return nil, 0, rx.Error
	}
	fmt.Println("count", count)
	if rx.RowsAffected == 0 {
		return nil, 0, errors.New("error query count")
	}
	// var totalpage int
	if count < 10 {
		totalpage = 1
	} else if int(count)%limit == 0 {
		totalpage = int(count) / limit
	} else {
		totalpage = (int(count) / limit) + 1
	}
	var checks []CheckupReservation

	tx := repo.db.Where("practice_id=?", id).Preload("Patient").Preload("Practice").Limit(limit).Offset(offset).Find(&checks)

	if tx.Error != nil {

		return nil, 0, tx.Error
	}
	gorms := toCoreList(checks)
	return gorms, totalpage, nil
}

// GetByreservationId implements checkupreservation.RepositoryInterface
func (repo *CheckUpRepository) GetByreservationId(id int) (data checkupreservation.CheckupReservationCore, err error) {
	var check CheckupReservation

	tx := repo.db.Preload("Patient").Preload("Practice").First(&check, id)

	if tx.Error != nil {

		return checkupreservation.CheckupReservationCore{}, tx.Error
	}

	gorms := check.toCore()
	return gorms, nil
}
