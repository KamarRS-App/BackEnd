package repository

import (
	"github.com/KamarRS-App/KamarRS-App/features/patient"

	"gorm.io/gorm"
)

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
	BedReservation        BedReservation     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CheckupReservation    CheckupReservation `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type User struct {
	gorm.Model
	Nama      string
	Email     string
	Nokk      string
	Nik       string
	KataSandi string
	NoTelpon  string
	Patients  []Patient `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type BedReservation struct {
	gorm.Model
	HospitalId       uint
	StatusPasien     string
	BiayaRegistrasi  int
	OrderId          string
	LinkPembayaran   string
	StatusPembayaran string
	PatientID        uint
}

type CheckupReservation struct {
	gorm.Model
	PatientID  uint
	PracticeID uint
	NamaDokter string
}

func FromPatientCore(dataCore patient.CorePatient) Patient { //fungsi yang mengambil data dari entities usercore dan merubah data ke user gorm(model.go)
	patientGorm := Patient{
		NoKk:                  dataCore.NoKk,
		Nik:                   dataCore.Nik,
		NamaPasien:            dataCore.NamaPasien,
		JenisKelamin:          dataCore.JenisKelamin,
		TanggalLahir:          dataCore.TanggalLahir,
		Usia:                  dataCore.Usia,
		NamaWali:              dataCore.NamaWali,
		EmailWali:             dataCore.EmailWali,
		NoTelponWali:          dataCore.NoTelponWali,
		AlamatKtp:             dataCore.AlamatKtp,
		ProvinsiKtp:           dataCore.ProvinsiKtp,
		KabupatenKotaKtp:      dataCore.KabupatenKotaKtp,
		AlamatDomisili:        dataCore.AlamatDomisili,
		ProvinsiDomisili:      dataCore.ProvinsiDomisili,
		KabupatenKotaDomisili: dataCore.KabupatenKotaDomisili,
		NoBpjs:                dataCore.NoBpjs,
		KelasBpjs:             dataCore.KelasBpjs,
		FotoKtp:               dataCore.FotoKtp,
		FotoBpjs:              dataCore.FotoBpjs,
		UserID:                dataCore.UserID,
	} ///formating data berdasarkan data gorm dan kita mapping data yang kita butuhkan untuk inputan  klien
	return patientGorm //insert user
}
func (dataModel *Patient) ModelsToCore() patient.CorePatient { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	return patient.CorePatient{
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
		UserID:                dataModel.UserID,
		CreatedAt:             dataModel.CreatedAt,
		UpdatedAt:             dataModel.UpdatedAt,
	}
}
func ListModelTOCore(dataModel []Patient) []patient.CorePatient { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	var dataCore []patient.CorePatient
	for _, value := range dataModel {
		dataCore = append(dataCore, value.ModelsToCore())
	}
	return dataCore //  untuk menampilkan data ke controller
}
