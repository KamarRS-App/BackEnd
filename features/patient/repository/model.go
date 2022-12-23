package repository

import (
	"gorm.io/gorm"
)

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
	CheckupReservation      CheckupReservation
}

type User struct {
	gorm.Model
	Nama       string
	Email      string
	No_kk      string
	Nik        string
	Kata_Sandi string
	No_Telpon  string
	Patients   []Patient
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
}

type CheckupReservation struct {
	gorm.Model
	PatientID  uint
	PracticeID uint
}
