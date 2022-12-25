package repository

import (
	"kamarRS/features/bedreservation"

	"gorm.io/gorm"
)

type BedReservation struct {
	gorm.Model
	Hospital_Id       uint
	Status_Pasien     string
	Biaya_Registrasi  int
	Order_Id          string
	Payment_Method    string
	Link_Pembayaran   string
	Status_Pembayaran string
	PatientID         uint
	BedID             uint
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
	BedReservation          BedReservation
}

type Bed struct {
	gorm.Model
	Nama_Tempat_Tidur string
	Ruangan           string
	Kelas             string
	Status            string
	HospitalID        uint
	BedReservation    BedReservation
}

func FromCoreBedToModel(dataCore bedreservation.BedReservationCore) BedReservation {
	bedresGorm := BedReservation{
		Status_Pasien:     dataCore.StatusPasien,
		Biaya_Registrasi:  dataCore.BiayaRegistrasi,
		Order_Id:          dataCore.OrderID,
		Payment_Method:    dataCore.PaymentMethod,
		Link_Pembayaran:   dataCore.LinkPembayaran,
		Status_Pembayaran: dataCore.StatusPembayaran,
		Hospital_Id:       dataCore.HospitalID,
		PatientID:         dataCore.PatientID,
		BedID:             dataCore.BedID,
	}
	return bedresGorm //insert bed from core
}

//
