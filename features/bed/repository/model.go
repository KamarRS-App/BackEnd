package repository

import (
	"kamarRS/features/bed"

	"gorm.io/gorm"
)

type Bed struct {
	gorm.Model
	Nama_Tempat_Tidur string
	Ruangan           string
	Kelas             string
	Status            string
	HospitalID        uint
	BedReservation    BedReservation
}

type Hospital struct {
	gorm.Model
	Kode_Rs             string
	Nama                string
	Foto                string
	Alamat              string
	Provinsi            string
	Kabupaten_Kota      string
	Kecamatan           string
	No_Telpon           string
	Email               string
	Kelas_Rs            string
	Pengelola           string
	Jumlah_Tempat_Tidur string
	Status_Penggunaan   string
	Biaya_Pendaftaran   string
	Beds                []Bed
}

type BedReservation struct {
	gorm.Model
	Hospital_Id       uint
	Status_Pasien     string
	Biaya_Registrasi  int
	Order_Id          string
	Link_Pembayaran   string
	Status_Pembayaran string
	PatientID         uint
	BedID             uint
}

func FromCoreToModel(dataCore bed.BedCore) Bed {
	bedGorm := Bed{
		Nama_Tempat_Tidur: dataCore.NamaTempatTidur,
		Ruangan:           dataCore.Ruangan,
		Kelas:             dataCore.Kelas,
		Status:            dataCore.Status,
		HospitalID:        dataCore.HospitalID,
	}
	return bedGorm //insert bed from core
}

//-------------------------------------------------------
// Bed aja

func (dataModel *Bed) toCore() bed.BedCore {
	return bed.BedCore{
		ID:              dataModel.ID,
		NamaTempatTidur: dataModel.Nama_Tempat_Tidur,
		Ruangan:         dataModel.Ruangan,
		Kelas:           dataModel.Kelas,
		Status:          dataModel.Status,
		HospitalID:      dataModel.HospitalID,
	}
}

// mengubah slice struct model gorm ke slice struct core
func toCoreList(dataModel []Bed) []bed.BedCore {
	var dataCore []bed.BedCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}

//---------------------------------------------------------
// Hospital Aja

// func (dataModel *Hospital) toHospitalCore() bed.HospitalCore {
// 	return bed.HospitalCore{
// 		ID:   dataModel.ID,
// 		Nama: dataModel.Nama,
// 	}
// }

// // mengubah slice struct model gorm ke slice struct core
// func toCoreList2(dataModel []Hospital) []bed.HospitalCore {
// 	var dataCore []bed.HospitalCore
// 	for _, v := range dataModel {
// 		dataCore = append(dataCore, v.toHospitalCore())
// 	}
// 	return dataCore
// }
