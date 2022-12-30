package checkupreservation

import (
	"time"
)

type CheckupReservationCore struct {
	ID         uint
	CreatedAt  time.Time
	PatientID  uint `validate:"required"`
	PracticeID uint `validate:"required"`
	Patient    PatientCore
	Practice   PracticeCore
}

type PatientCore struct {
	ID                    uint
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
}

type PracticeCore struct {
	ID             uint
	TanggalPraktik string
	KuotaHarian    int
	Status         string
	PoliclinicID   uint
}

type ServiceInterface interface { //sebagai contract yang dibuat di layer service

	Create(input CheckupReservationCore, userId int) (err error) // menambahkah data user berdasarkan data usercore
	// Update(id int, input CoreUser) error
	GetByPracticesId(pagination, limit, id int) (data []CheckupReservationCore, totalpage int, err error)
	// DeleteById(id int) error
	GetByreservationId(id int) (data CheckupReservationCore, err error)
}

type RepositoryInterface interface { // berkaitan database

	Create(input CheckupReservationCore, userId int) (err error)
	GetByPracticesId(offset, limit, id int) (data []CheckupReservationCore, totalpage int, err error)
	// GetById(id int) (data CoreUser, err error)
	// DeleteById(id int) error
	GetByreservationId(id int) (data CheckupReservationCore, err error)
}
