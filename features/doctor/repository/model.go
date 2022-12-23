package repository

import "gorm.io/gorm"

type Doctor struct {
	gorm.Model
	Nama        string
	Bidang      string
	Email       string
	No_Telpon   string
	Policlinics []Policlinic
}

type Policlinic struct {
	gorm.Model
	Nama_Poli   string
	Jam_Praktik string
	HospitalID  uint
	DoctorID    uint
}
