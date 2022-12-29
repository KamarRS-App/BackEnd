package patient

import "time"

type CorePatient struct {
	ID                    uint
	NoKk                  string `validate:"required"`
	Nik                   string `validate:"required"`
	NamaPasien            string `validate:"required"`
	JenisKelamin          string `validate:"required"`
	TanggalLahir          string `validate:"required"`
	Usia                  int    `validate:"required"`
	NamaWali              string `validate:"required"`
	EmailWali             string `validate:"required"`
	NoTelponWali          string `validate:"required"`
	AlamatKtp             string `validate:"required"`
	ProvinsiKtp           string `validate:"required"`
	KabupatenKotaKtp      string `validate:"required"`
	AlamatDomisili        string `validate:"required"`
	ProvinsiDomisili      string `validate:"required"`
	KabupatenKotaDomisili string `validate:"required"`
	NoBpjs                string
	KelasBpjs             string
	FotoKtp               string
	FotoBpjs              string
	UserID                uint `validate:"required"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	// BedReservation          BedReservation
	// CheckupReservation      CheckupReservation
}

type ServiceInterface interface { //sebagai contract yang dibuat di layer service

	Create(input CorePatient) (err error) // menambahkah data user berdasarkan data usercore
	Update(id int, input CorePatient) error
	GetByPatientId(id int) (data CorePatient, err error)
	GetByUserId(pagination, limit, id int) (data []CorePatient, totalpage int, err error)
	GetAllPatient() (data []CorePatient, err error)
	DeleteById(id int) error
}

type RepositoryInterface interface { // berkaitan database

	Create(input CorePatient) (err error)
	Update(id int, input CorePatient) error
	GetByPatientId(id int) (data CorePatient, err error)
	GetByUserId(limit, offset, id int) (data []CorePatient, totalpage int, err error)
	GetAllPatient() (data []CorePatient, err error)
	DeleteById(id int) error
}
