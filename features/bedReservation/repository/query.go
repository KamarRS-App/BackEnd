package repository

import (
	"errors"
	"time"

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

// CreatePayment implements bedreservation.RepositoryInterface
func (r *bedReservationRepository) CreatePayment(input bedreservation.BedReservationCore) (data bedreservation.BedReservationCore, err error) {
	// var bedReservation BedReservation
	var regisInfo BedReservation

	tx := r.db.Where("kode_daftar = ?", input.KodeDaftar).First(&regisInfo)
	if tx.Error != nil {
		return bedreservation.BedReservationCore{}, tx.Error
	}
	input.BiayaRegistrasi = regisInfo.BiayaRegistrasi
	input.HospitalID = regisInfo.HospitalID
	midtransInfo := helper.CreateInvoice(input.KodeDaftar, int64(regisInfo.BiayaRegistrasi), input.PaymentMethod)
	input.LinkPembayaran = midtransInfo.RedirectURL
	input.StatusPembayaran = midtransInfo.TransactionStatus
	input.QrString = midtransInfo.QRString

	expirationTimeParse, _ := time.Parse("2006-01-02 15:04:05", midtransInfo.TransactionTime)
	expirationTime := expirationTimeParse.Add(time.Hour * 24).String()
	input.WaktuKedaluarsa = expirationTime

	input.TransactionId = midtransInfo.TransactionID
	if input.PaymentMethod == "transfer_va_permata" {
		input.VirtualAccount = midtransInfo.PermataVaNumber
	} else if input.PaymentMethod == "qris" {
		input.VirtualAccount = ""
	} else {
		input.VirtualAccount = midtransInfo.VaNumbers[0].VANumber
	}

	switch {
	case input.PaymentMethod == "transfer_va_permata":
		input.LinkPembayaran = "https://simulator.sandbox.midtrans.com/permata/va/index"
		input.BankPenerima = "bank permata"
	case input.PaymentMethod == "transfer_va_bca":
		input.LinkPembayaran = "https://simulator.sandbox.midtrans.com/bca/va/index"
		input.BankPenerima = "bank bca"
	case input.PaymentMethod == "transfer_va_bri":
		input.LinkPembayaran = "https://simulator.sandbox.midtrans.com/bri/va/index"
		input.BankPenerima = "bank bri"
	case input.PaymentMethod == "transfer_va_bni":
		input.LinkPembayaran = "https://simulator.sandbox.midtrans.com/bni/va/index"
		input.BankPenerima = "bank bni"
	case input.PaymentMethod == "qris":
		input.LinkPembayaran = "https://simulator.sandbox.midtrans.com/qris/index"
		input.BankPenerima = "bank? boekan, ini qris"
	}

	input.StatusPasien = "waiting list"
	helper.LogDebug("\n midtrans = ", *midtransInfo)

	if midtransInfo == nil {
		return bedreservation.BedReservationCore{}, errors.New("failed create payment")
	}

	inputGorm := FromCoreToModel(input)
	tx1 := r.db.Where("kode_daftar = ?", input.KodeDaftar).Updates(inputGorm)
	if tx1.Error != nil {
		return bedreservation.BedReservationCore{}, tx1.Error
	}
	if tx1.RowsAffected == 0 {
		return bedreservation.BedReservationCore{}, errors.New("create payment failed, error query")
	}
	return input, nil
}
