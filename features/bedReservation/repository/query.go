package repository

import (
	"errors"
	"log"
	"time"

	bedreservation "github.com/KamarRS-App/KamarRS-App/features/bedReservation"
	user "github.com/KamarRS-App/KamarRS-App/features/user/repository"
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

// UpdateBedReservation implements bedreservation.RepositoryInterface
func (r *bedReservationRepository) UpdateBedReservation(input bedreservation.BedReservationCore) error {
	var regisInfo BedReservation
	var patientInfo Patient
	var hospitalInfo Hospital
	var inputGorm BedReservation = FromCoreToModel(input)

	tx := r.db.First(&regisInfo, input.ID)
	if tx.Error != nil {
		return tx.Error
	}

	tx0 := r.db.First(&patientInfo, regisInfo.PatientID)
	if tx0.Error != nil {
		return tx0.Error
	}

	tx1 := r.db.First(&hospitalInfo, regisInfo.HospitalID)
	if tx1.Error != nil {
		return tx1.Error
	}

	tx2 := r.db.Where("id = ?", input.ID).Updates(inputGorm)
	if tx2.Error != nil {
		return tx2.Error
	}

	if tx2.RowsAffected == 0 {
		return errors.New("error on update")
	}
	if inputGorm.StatusPasien == "pemeriksaan awal" {

		dataEmail := struct {
			RumahSakit string
			NamaPasien string
		}{
			RumahSakit: hospitalInfo.Nama,
			NamaPasien: patientInfo.NamaPasien,
		}

		emailTo := patientInfo.EmailWali

		errMail := helper.SendEmailSMTPBed([]string{emailTo}, dataEmail, "bedReservationEmail.txt") //send mail
		if errMail != nil {
			log.Println(errMail, "Pengiriman Email Gagal")
		}
	}
	return nil

}

// Delete implements bedreservation.RepositoryInterface
func (r *bedReservationRepository) Delete(bedResId uint) error {
	var bedReservation BedReservation
	tx := r.db.Delete(&bedReservation, bedResId)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("id not found")

	}

	return nil
}

// GetById implements bedreservation.RepositoryInterface
func (r *bedReservationRepository) GetById(bedResId uint) (data bedreservation.BedReservationCore, err error) {
	var regisDetail BedReservation

	tx := r.db.Where("id = ?", bedResId).First(&regisDetail)
	if tx.Error != nil {
		return bedreservation.BedReservationCore{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return bedreservation.BedReservationCore{}, errors.New("get data failed, error query data")
	}

	data = regisDetail.toCore()
	return data, nil
}

// GetRegistrations implements bedreservation.RepositoryInterface
func (r *bedReservationRepository) GetRegistrations(limit, offset, hospitalId int) (data []bedreservation.BedReservationCore, totalpage int, err error) {
	var reservations []BedReservation
	var count int64
	tx0 := r.db.Model(&reservations).Where("hospital_id = ? AND bed_id = 0", hospitalId).Count(&count)
	if tx0.Error != nil {
		return nil, 0, tx0.Error
	}
	if tx0.RowsAffected == 0 {
		return nil, 0, errors.New("error query count")
	}

	if count < 10 {
		totalpage = 1
	} else if int(count)%limit == 0 {
		totalpage = int(count) / limit
	} else {
		totalpage = (int(count) / limit) + 1
	}

	tx := r.db.Where("hospital_id = ? AND bed_id = 0", hospitalId).Limit(limit).Offset(offset).Order("status_pembayaran desc").Find(&reservations)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, 0, errors.New("get all data failed, error query data")
	}
	data = toCoreList(reservations)
	return data, totalpage, nil

}

// Create implements bedreservation.RepositoryInterface
func (r *bedReservationRepository) Create(input bedreservation.BedReservationCore, userId uint) (data bedreservation.BedReservationCore, err error) {
	var user user.User
	tx0 := r.db.Where("id = ?", userId).First(&user)
	if tx0.Error != nil {
		return bedreservation.BedReservationCore{}, tx0.Error
	}

	var patient Patient
	tx1 := r.db.Where("id = ?", input.PatientID).First(&patient)
	if tx1.Error != nil {
		return bedreservation.BedReservationCore{}, tx1.Error
	}

	var hospital Hospital
	tx2 := r.db.First(&hospital, input.HospitalID)
	if tx2.Error != nil {
		return bedreservation.BedReservationCore{}, tx2.Error
	}

	if user.Nokk != patient.NoKk {
		return bedreservation.BedReservationCore{}, errors.New("pasien hanya dapat didaftarkan oleh user dengan kk sama")
	}

	if patient.NoBpjs != "" {
		input.BiayaRegistrasi = 0
		input.StatusPembayaran = "settlement via BPJS"
	} else {
		input.BiayaRegistrasi = hospital.BiayaPendaftaran
		input.StatusPembayaran = "belum dibayar"
	}
	randString := helper.FileName(5)
	input.KodeDaftar = "order-" + randString
	inputGorm := FromCoreToModel(input)
	tx3 := r.db.Create(&inputGorm)
	if tx3.Error != nil {
		return bedreservation.BedReservationCore{}, tx3.Error
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
	var regisInfo BedReservation
	var patientInfo Patient
	var hospitalInfo Hospital

	tx := r.db.Where("kode_daftar = ?", input.KodeDaftar).First(&regisInfo)
	if tx.Error != nil {
		return bedreservation.BedReservationCore{}, tx.Error
	}

	tx0 := r.db.First(&patientInfo, regisInfo.PatientID)
	if tx0.Error != nil {
		return bedreservation.BedReservationCore{}, tx0.Error
	}

	tx1 := r.db.First(&hospitalInfo, regisInfo.HospitalID)
	if tx1.Error != nil {
		return bedreservation.BedReservationCore{}, tx1.Error
	}

	if regisInfo.BiayaRegistrasi < 1 {
		return bedreservation.BedReservationCore{}, errors.New("tidak perlu melakukan pembayaran, pembayaran anda sudah ditanggung BPJS")
	}

	input.BiayaRegistrasi = regisInfo.BiayaRegistrasi
	input.HospitalID = regisInfo.HospitalID
	midtransInfo := helper.CreateInvoice(input.KodeDaftar, int64(regisInfo.BiayaRegistrasi), input.PaymentMethod)

	switch {
	case midtransInfo.TransactionID != "":
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
		tx2 := r.db.Where("kode_daftar = ?", input.KodeDaftar).Updates(inputGorm)
		if tx2.Error != nil {
			return bedreservation.BedReservationCore{}, tx2.Error
		}
		if tx2.RowsAffected == 0 {
			return bedreservation.BedReservationCore{}, errors.New("create payment failed, error query")
		}

		switch {
		case inputGorm.PaymentMethod == "qris":
			tenggatWaktu := inputGorm.WaktuKedaluarsa[:20] + "WIB"
			dataEmail := struct {
				RumahSakit      string
				NamaPasien      string
				KodeDaftar      string
				QrString        string
				BiayaRegistrasi int
				WaktuKedaluarsa string
				LinkPembayaran  string
			}{
				RumahSakit:      hospitalInfo.Nama,
				NamaPasien:      patientInfo.NamaPasien,
				KodeDaftar:      inputGorm.KodeDaftar,
				QrString:        inputGorm.QrString,
				BiayaRegistrasi: inputGorm.BiayaRegistrasi,
				WaktuKedaluarsa: tenggatWaktu,
				LinkPembayaran:  inputGorm.LinkPembayaran,
			}

			emailTo := patientInfo.EmailWali
			errMail := helper.SendEmailInvoice([]string{emailTo}, dataEmail, "invoiceBRQris.txt") //send mail
			if errMail != nil {
				log.Println(errMail, "Pengiriman Email Gagal")
			}
		default:
			tenggatWaktu := inputGorm.WaktuKedaluarsa[:20] + "WIB"
			dataEmail := struct {
				RumahSakit      string
				NamaPasien      string
				KodeDaftar      string
				BankPenerima    string
				VirtualAccount  string
				BiayaRegistrasi int
				WaktuKedaluarsa string
				LinkPembayaran  string
			}{
				RumahSakit:      hospitalInfo.Nama,
				NamaPasien:      patientInfo.NamaPasien,
				KodeDaftar:      inputGorm.KodeDaftar,
				BankPenerima:    inputGorm.BankPenerima,
				VirtualAccount:  inputGorm.VirtualAccount,
				BiayaRegistrasi: inputGorm.BiayaRegistrasi,
				WaktuKedaluarsa: tenggatWaktu,
				LinkPembayaran:  inputGorm.LinkPembayaran,
			}

			emailTo := patientInfo.EmailWali
			errMail := helper.SendEmailInvoice([]string{emailTo}, dataEmail, "invoiceBRBank.txt") //send mail
			if errMail != nil {
				log.Println(errMail, "Pengiriman Email Gagal")
			}
		}

		return input, nil
	case midtransInfo.TransactionID == "":
		return bedreservation.BedReservationCore{}, errors.New("terjadi kesalahan pembayaran, pilih metode pembayaran lain")
	}
	return bedreservation.BedReservationCore{}, errors.New("create payment failed, error query")
}

// PaymentNotif implements bedreservation.RepositoryInterface
func (r *bedReservationRepository) PaymentNotif(callback bedreservation.BedReservationCore) (err error) {
	updatePayment := helper.UpdateMidtransPayment(callback.KodeDaftar)
	callback.StatusPembayaran = updatePayment.TransactionStatus

	updateGorm := FromCoreToModel(callback)
	tx := r.db.Where("kode_daftar = ?", callback.KodeDaftar).Updates(updateGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("update payment failed, error query")
	}
	return nil
}
