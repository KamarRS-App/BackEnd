package repository

import (
	bedreservation "github.com/KamarRS-App/KamarRS-App/features/bedReservation"

	"gorm.io/gorm"
)

type BedReservation struct {
	gorm.Model
	StatusPasien     string
	BiayaRegistrasi  int
	KodeDaftar       string
	PaymentMethod    string
	LinkPembayaran   string
	TransactionId    string
	VirtualAccount   string
	BankPenerima     string
	WaktuKedaluarsa  string
	QrString         string
	StatusPembayaran string
	PatientID        uint
	HospitalID       uint
	BedID            uint
	Hospital         Hospital `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Patient struct {
	gorm.Model
	NoKk                  string
	Nik                   string
	NamaPasien            string
	JenisKelamin          string
	TanggalLahir          string
	Usia                  int
	NamaWali              string
	EmailWali             string
	NoTelponWali          string
	AlamatKtp             string
	ProvinsiKtp           string
	KabupatenKotaKtp      string
	AlamatDomisili        string
	ProvinsiDomisili      string
	KabupatenKotaDomisili string
	NoBpjs                string
	KelasBpjs             string
	FotoKtp               string
	FotoBpjs              string
	UserID                uint
	// BedReservationID        uint
	BedReservations []BedReservation `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Bed struct {
	gorm.Model
	NamaTempatTidur string
	Ruangan         string
	Kelas           string
	Status          string
	HospitalID      uint
	// BedReservation  BedReservation `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Hospital struct {
	gorm.Model
	KodeRs            string
	Nama              string
	Foto              string
	Alamat            string
	Provinsi          string
	KabupatenKota     string
	Kecamatan         string
	KodePos           string
	NoTelpon          string
	Email             string
	KelasRs           string
	Pengelola         string
	JumlahTempatTidur int
	StatusPenggunaan  string
	BiayaPendaftaran  int
	BedReservations   []BedReservation `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Beds              []Bed            `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func FromCoreToModel(dataCore bedreservation.BedReservationCore) BedReservation {
	bedresGorm := BedReservation{
		StatusPasien:     dataCore.StatusPasien,
		BiayaRegistrasi:  dataCore.BiayaRegistrasi,
		KodeDaftar:       dataCore.KodeDaftar,
		PaymentMethod:    dataCore.PaymentMethod,
		LinkPembayaran:   dataCore.LinkPembayaran,
		VirtualAccount:   dataCore.VirtualAccount,
		BankPenerima:     dataCore.BankPenerima,
		TransactionId:    dataCore.TransactionId,
		WaktuKedaluarsa:  dataCore.WaktuKedaluarsa,
		QrString:         dataCore.QrString,
		StatusPembayaran: dataCore.StatusPembayaran,
		HospitalID:       dataCore.HospitalID,
		PatientID:        dataCore.PatientID,
		BedID:            dataCore.BedID,
	}
	return bedresGorm //insert bedreserve from core
}

//----------------------BedReserve Aja-------------------------------

func (dataModel *BedReservation) toCore() bedreservation.BedReservationCore {
	return bedreservation.BedReservationCore{
		ID:               dataModel.ID,
		StatusPasien:     dataModel.StatusPasien,
		BiayaRegistrasi:  dataModel.BiayaRegistrasi,
		KodeDaftar:       dataModel.KodeDaftar,
		PaymentMethod:    dataModel.PaymentMethod,
		LinkPembayaran:   dataModel.LinkPembayaran,
		VirtualAccount:   dataModel.VirtualAccount,
		BankPenerima:     dataModel.BankPenerima,
		TransactionId:    dataModel.TransactionId,
		WaktuKedaluarsa:  dataModel.WaktuKedaluarsa,
		QrString:         dataModel.QrString,
		StatusPembayaran: dataModel.StatusPembayaran,
		HospitalID:       dataModel.HospitalID,
		BedID:            dataModel.BedID,
		// Patient: bedreservation.PatientCore{
		// 	ID:                    dataModel.Patient.ID,
		// 	NoKk:                  dataModel.Patient.NoKk,
		// 	Nik:                   dataModel.Patient.Nik,
		// 	NamaPasien:            dataModel.Patient.NamaPasien,
		// 	JenisKelamin:          dataModel.Patient.JenisKelamin,
		// 	TanggalLahir:          dataModel.Patient.TanggalLahir,
		// 	Usia:                  dataModel.Patient.Usia,
		// 	NamaWali:              dataModel.Patient.NamaWali,
		// 	EmailWali:             dataModel.Patient.EmailWali,
		// 	NoTelponWali:          dataModel.Patient.NoTelponWali,
		// 	AlamatKtp:             dataModel.Patient.AlamatKtp,
		// 	ProvinsiKtp:           dataModel.Patient.ProvinsiKtp,
		// 	KabupatenKotaKtp:      dataModel.Patient.KabupatenKotaKtp,
		// 	AlamatDomisili:        dataModel.Patient.AlamatDomisili,
		// 	ProvinsiDomisili:      dataModel.Patient.ProvinsiDomisili,
		// 	KabupatenKotaDomisili: dataModel.Patient.KabupatenKotaDomisili,
		// 	NoBpjs:                dataModel.Patient.NoBpjs,
		// 	KelasBpjs:             dataModel.Patient.KelasBpjs,
		// 	FotoKtp:               dataModel.Patient.FotoKtp,
		// 	FotoBpjs:              dataModel.Patient.FotoBpjs,
		// UserID:                dataModel.Patient.UserID,
		// },
	}
}

// mengubah slice struct model gorm ke slice struct core
func toCoreList(dataModel []BedReservation) []bedreservation.BedReservationCore {
	var dataCore []bedreservation.BedReservationCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}

// //----------------------Patient Aja-------------------------------

func (dataModel *Patient) toCoreP() bedreservation.PatientCore {
	// var arrBedReservations []bedreservation.BedReservationCore
	// for _, v := range dataModel.BedReservation {
	// 	arrBedReservations = append(arrBedReservations, bedreservation.BedReservationCore{
	// 		ID:               v.ID,
	// 		StatusPasien:     v.StatusPasien,
	// 		BiayaRegistrasi:  v.BiayaRegistrasi,
	// 		KodeDaftar:       v.KodeDaftar,
	// 		PaymentMethod:    v.PaymentMethod,
	// 		LinkPembayaran:   v.LinkPembayaran,
	// 		VirtualAccount:   v.VirtualAccount,
	// 		BankPenerima:     v.BankPenerima,
	// 		TransactionId:    v.TransactionId,
	// 		WaktuKedaluarsa:  v.WaktuKedaluarsa,
	// 		QrString:         v.QrString,
	// 		StatusPembayaran: v.StatusPembayaran,
	// 		HospitalID:       v.HospitalID,
	// 		BedID:            v.BedID,
	// 	})
	// }

	return bedreservation.PatientCore{
		ID:                    dataModel.ID,
		NoKk:                  dataModel.NoKk,
		Nik:                   dataModel.Nik,
		NamaPasien:            dataModel.NamaPasien,
		JenisKelamin:          dataModel.JenisKelamin,
		TanggalLahir:          dataModel.TanggalLahir,
		Usia:                  dataModel.Usia,
		NamaWali:              dataModel.NamaWali,
		EmailWali:             dataModel.EmailWali,
		NoTelponWali:          dataModel.NoTelponWali,
		AlamatKtp:             dataModel.AlamatKtp,
		ProvinsiKtp:           dataModel.ProvinsiKtp,
		KabupatenKotaKtp:      dataModel.KabupatenKotaKtp,
		AlamatDomisili:        dataModel.AlamatDomisili,
		ProvinsiDomisili:      dataModel.ProvinsiDomisili,
		KabupatenKotaDomisili: dataModel.KabupatenKotaDomisili,
		NoBpjs:                dataModel.NoBpjs,
		KelasBpjs:             dataModel.KelasBpjs,
		FotoKtp:               dataModel.FotoKtp,
		FotoBpjs:              dataModel.FotoBpjs,
		// BedReservation:       arrBedReservations,
	}
}

// // mengubah slice struct model gorm ke slice struct core
// func toCoreListP(dataModel []Patient) []bedreservation.PatientCore {
// 	var dataCore []bedreservation.PatientCore
// 	for , v := range dataModel {
// 		dataCore = append(dataCore, v.toCoreP())
// 	}
// 	return dataCore
// }

// //---------------------------------------------------------------------------------
