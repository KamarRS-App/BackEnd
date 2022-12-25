package checkupreservation

import "time"

type CheckupReservationCore struct {
	ID         uint
	CreatedAt  time.Time
	PatientID  uint
	PracticeID uint
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
