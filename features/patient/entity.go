package patient

import "time"

type CorePatient struct {
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
	UserID                uint
	CreatedAt             time.Time
	UpdatedAt             time.Time
	// BedReservation          BedReservation
	// CheckupReservation      CheckupReservation
}

type ServiceInterface interface { //sebagai contract yang dibuat di layer service

	Create(input CorePatient) (err error) // menambahkah data user berdasarkan data usercore
	Update(id, userId int, input CorePatient) error
	GetByPatientId(id int) (data CorePatient, err error)
	GetByUserId(pagination, limit, id int) (data []CorePatient, totalpage int, err error)
	GetAllPatient() (data []CorePatient, err error)
	DeleteById(id int) error
}

type RepositoryInterface interface { // berkaitan database

	Create(input CorePatient) (err error)
	Update(id, userId int, input CorePatient) error
	GetByPatientId(id int) (data CorePatient, err error)
	GetByUserId(limit, offset, id int) (data []CorePatient, totalpage int, err error)
	GetAllPatient() (data []CorePatient, err error)
	DeleteById(id int) error
}
