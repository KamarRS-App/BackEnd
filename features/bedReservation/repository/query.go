package repository

import (
	bedreservation "github.com/KamarRS-App/KamarRS-App/features/bedReservation"
	"github.com/KamarRS-App/KamarRS-App/utils/helper"
	"gorm.io/gorm"
)

type bedReservationRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) bedreservation.RepositoryInterface {
	return &bedReservationRepository{
		db: db,
	}
}

// Create implements bedreservation.RepositoryInterface
func (r *bedReservationRepository) Create(input bedreservation.BedReservationCore) (data bedreservation.BedReservationCore, err error) {
	var patient Patient
	tx1 := r.db.Where("id = ?", input.PatientID).First(&patient)
	if tx1.Error != nil {
		return bedreservation.BedReservationCore{}, tx1.Error
	}
	if patient.NoBpjs != "" {
		input.BiayaRegistrasi = 0
	} else {
		input.BiayaRegistrasi = 25000
	}
	randString := helper.FileName(5)
	input.KodeDaftar = "order-" + randString
	input.StatusPembayaran = "belum dibayar"
	inputGorm := FromCoreToModel(input)
	tx2 := r.db.Create(&inputGorm)
	if tx2.Error != nil {
		return bedreservation.BedReservationCore{}, tx2.Error
	}
	return input, nil
}

// GetPayment implements bedreservation.RepositoryInterface
func (r *bedReservationRepository) GetPayment(kodeDaftar string) (data bedreservation.BedReservationCore, err error) {
	var registration BedReservation
	tx := r.db.Where("kode_daftar = ?", kodeDaftar).First(&registration)
	if tx.Error != nil {
		return bedreservation.BedReservationCore{}, tx.Error
	}
	data = registration.toCore()
	return data, nil
}
