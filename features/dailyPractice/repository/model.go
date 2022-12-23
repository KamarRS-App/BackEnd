package repository

import "gorm.io/gorm"

type Practice struct {
	gorm.Model
	Tanggal_Praktik string
	Kuota_Harian    int
	Status          string
	PoliclinicID    uint
}

type Policlinic struct {
	gorm.Model
	Nama_Poli   string
	Jam_Praktik string
	HospitalID  uint
	DoctorID    uint
	Practices   []Practice
}
