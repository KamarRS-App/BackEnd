package checkupreservation

import (
	"time"

	dailypractice "github.com/KamarRS-App/KamarRS-App/features/dailyPractice/repository"
	"github.com/KamarRS-App/KamarRS-App/features/patient/repository"
	userRep "github.com/KamarRS-App/KamarRS-App/features/user/repository"
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
	// GetById(id int) (data CoreUser, err error)
	// DeleteById(id int) error
}

type RepositoryInterface interface { // berkaitan database

	Create(input CheckupReservationCore, userId int) (users userRep.User, patients repository.Patient, practice dailypractice.Practice, err error)
	// Update(id int, input CoreUser) error
	// GetById(id int) (data CoreUser, err error)
	// DeleteById(id int) error
}
