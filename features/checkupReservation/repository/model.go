package repository

import (
	checkupreservation "github.com/KamarRS-App/KamarRS-App/features/checkupReservation"

	"gorm.io/gorm"
)

type CheckupReservation struct {
	gorm.Model
	PatientID  uint
	PracticeID uint
	Patient    Patient
	Practice   Practice
}

type Patient struct {
	gorm.Model
	No_Kk                   string
	Nik                     string
	Nama_Pasien             string
	Jenis_Kelamin           string
	Tanggal_Lahir           string
	Usia                    int
	Nama_Wali               string
	Email_Wali              string
	No_Telpon_Wali          string
	Alamat_Ktp              string
	Provinsi_Ktp            string
	Kabupaten_Kota_Ktp      string
	Alamat_Domisili         string
	Provinsi_Domisili       string
	Kabupaten_Kota_Domisili string
	No_Bpjs                 string
	Kelas_Bpjs              string
	Foto_Ktp                string
	Foto_Bpjs               string
	UserID                  uint
	// CheckupReservation      CheckupReservation
}

type Practice struct {
	gorm.Model
	Tanggal_Praktik     string
	Kuota_Harian        int
	Status              string
	PoliclinicID        uint
	CheckupReservations []CheckupReservation
}

func FromCoreToModel(dataCore checkupreservation.CheckupReservationCore) CheckupReservation {
	checkupGorm := CheckupReservation{
		PatientID:  dataCore.PatientID,
		PracticeID: dataCore.PracticeID,
	}
	return checkupGorm //insert checkup from core
}

//---------------------Checkup Reservation----------------------------------

func (dataModel *CheckupReservation) toCore() checkupreservation.CheckupReservationCore {
	return checkupreservation.CheckupReservationCore{
		ID:        dataModel.ID,
		PatientID: dataModel.PatientID,
		CreatedAt: dataModel.CreatedAt,
		Patient: checkupreservation.PatientCore{
			ID:                    dataModel.Patient.ID,
			NoKk:                  dataModel.Patient.No_Kk,
			Nik:                   dataModel.Patient.Nik,
			NamaPasien:            dataModel.Patient.Nama_Pasien,
			JenisKelamin:          dataModel.Patient.Jenis_Kelamin,
			TanggalLahir:          dataModel.Patient.Tanggal_Lahir,
			Usia:                  dataModel.Patient.Usia,
			NamaWali:              dataModel.Patient.Nama_Wali,
			EmailWali:             dataModel.Patient.Email_Wali,
			NoTelponWali:          dataModel.Patient.No_Telpon_Wali,
			AlamatKtp:             dataModel.Patient.Alamat_Ktp,
			ProvinsiKtp:           dataModel.Patient.Provinsi_Ktp,
			KabupatenKotaKtp:      dataModel.Patient.Kabupaten_Kota_Ktp,
			AlamatDomisili:        dataModel.Patient.Alamat_Domisili,
			ProvinsiDomisili:      dataModel.Patient.Provinsi_Domisili,
			KabupatenKotaDomisili: dataModel.Patient.Kabupaten_Kota_Domisili,
			NoBpjs:                dataModel.Patient.No_Bpjs,
			KelasBpjs:             dataModel.Patient.Kelas_Bpjs,
			FotoKtp:               dataModel.Patient.Foto_Ktp,
			FotoBpjs:              dataModel.Patient.Foto_Bpjs,
		},
		Practice: checkupreservation.PracticeCore{
			ID:             dataModel.Practice.ID,
			TanggalPraktik: dataModel.Practice.Tanggal_Praktik,
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
