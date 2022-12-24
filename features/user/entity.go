package user

import (
	"time"
)

type CoreUser struct {
	ID        uint
	Nama      string
	Email     string
	Nokk      string
	Nik       string
	KataSandi string
	NoTelpon  string
	Patients  []CorePatient
	CreatedAt time.Time
	UpdatedAt time.Time
}

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
