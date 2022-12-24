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
	FotoKtp               string `validate:"required"`
	FotoBpjs              string
	UserID                uint `validate:"required"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	// BedReservation          BedReservation
	// CheckupReservation      CheckupReservation
}
