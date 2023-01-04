package repository

import (
	checkupreservation "github.com/KamarRS-App/KamarRS-App/features/checkupReservation"

	"gorm.io/gorm"
)

type CheckupReservation struct {
	gorm.Model
	PatientID  uint
	PracticeID uint
	NamaDokter string
	NoAntrian  string
	Patient    Patient  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Practice   Practice `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
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
	CheckupReservations   []CheckupReservation
}

type Practice struct {
	gorm.Model
	TanggalPraktik      string
	KuotaHarian         int
	Status              string
	PoliclinicID        uint
	CheckupReservations []CheckupReservation `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func FromCoreToModel(dataCore checkupreservation.CheckupReservationCore) CheckupReservation {
	checkupGorm := CheckupReservation{
		PatientID:  dataCore.PatientID,
		PracticeID: dataCore.PracticeID,
		NamaDokter: dataCore.NamaDokter,
		NoAntrian:  dataCore.NoAntrian,
	}
	return checkupGorm //insert checkup from core
}

//---------------------Checkup Reservation----------------------------------

func (dataModel *CheckupReservation) toCore() checkupreservation.CheckupReservationCore {
	return checkupreservation.CheckupReservationCore{
		ID:         dataModel.ID,
		PatientID:  dataModel.PatientID,
		NamaDokter: dataModel.NamaDokter,
		CreatedAt:  dataModel.CreatedAt,
		NoAntrian:  dataModel.NoAntrian,
		Patient: checkupreservation.PatientCore{
			ID:                    dataModel.Patient.ID,
			NoKk:                  dataModel.Patient.NoKk,
			Nik:                   dataModel.Patient.Nik,
			NamaPasien:            dataModel.Patient.NamaPasien,
			JenisKelamin:          dataModel.Patient.JenisKelamin,
			TanggalLahir:          dataModel.Patient.TanggalLahir,
			Usia:                  dataModel.Patient.Usia,
			NamaWali:              dataModel.Patient.NamaWali,
			EmailWali:             dataModel.Patient.EmailWali,
			NoTelponWali:          dataModel.Patient.NoTelponWali,
			AlamatKtp:             dataModel.Patient.AlamatKtp,
			ProvinsiKtp:           dataModel.Patient.ProvinsiKtp,
			KabupatenKotaKtp:      dataModel.Patient.KabupatenKotaKtp,
			AlamatDomisili:        dataModel.Patient.AlamatDomisili,
			ProvinsiDomisili:      dataModel.Patient.ProvinsiDomisili,
			KabupatenKotaDomisili: dataModel.Patient.KabupatenKotaDomisili,
			NoBpjs:                dataModel.Patient.NoBpjs,
			KelasBpjs:             dataModel.Patient.KelasBpjs,
			FotoKtp:               dataModel.Patient.FotoKtp,
			FotoBpjs:              dataModel.Patient.FotoBpjs,
		},
		Practice: checkupreservation.PracticeCore{
			ID:             dataModel.Practice.ID,
			TanggalPraktik: dataModel.Practice.TanggalPraktik,
			Status:         dataModel.Practice.Status,
			PoliclinicID:   dataModel.Practice.PoliclinicID,
		},
	}
}

// mengubah slice struct model gorm ke slice struct core
func toCoreList(dataModel []CheckupReservation) []checkupreservation.CheckupReservationCore {
	var dataCore []checkupreservation.CheckupReservationCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}

//----------------------------------------------------------------------------
