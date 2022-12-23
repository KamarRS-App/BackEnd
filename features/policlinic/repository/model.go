package repository

import "gorm.io/gorm"

type Policlinic struct {
	gorm.Model
	Nama_Poli   string
	Jam_Praktik string
	HospitalID  uint
	DoctorID    uint
	Practices   []Practice
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
	Policlinics         []Policlinic
}

type Doctor struct {
	gorm.Model
	Nama        string
	Bidang      string
	Email       string
	No_Telpon   string
	Policlinics []Policlinic
}

type Practice struct {
	gorm.Model
	Tanggal_Praktik string
	Kuota_Harian    int
	Status          string
	PoliclinicID    uint
}
