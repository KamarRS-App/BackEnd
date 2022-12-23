package repository

import "gorm.io/gorm"

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
	HospitalStaffs      []HospitalStaff
	Beds                []Bed
	Policlinics         []Policlinic
}

type HospitalStaff struct {
	gorm.Model
	Nama       string
	Email      string
	Kata_Sandi string
	Peran      string
	HospitalID uint
}

type Bed struct {
	gorm.Model
	Nama_Tempat_Tidur string
	Ruangan           string
	Kelas             string
	Status            string
	HospitalID        uint
}

type Policlinic struct {
	gorm.Model
	Nama_Poli   string
	Jam_Praktik string
	HospitalID  uint
}
