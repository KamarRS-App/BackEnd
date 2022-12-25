package repository

import (
	checkupreservation "kamarRS/features/checkupReservation"

	"gorm.io/gorm"
)

type CheckupReservation struct {
	gorm.Model
	PatientID  uint
	PracticeID uint
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
	CheckupReservation      CheckupReservation
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
		ID:         dataModel.ID,
		PatientID:  dataModel.PatientID,
		PracticeID: dataModel.PracticeID,
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
