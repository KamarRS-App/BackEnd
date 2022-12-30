package mysql

import (
	"fmt"
	"log"

	"github.com/KamarRS-App/KamarRS-App/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}

	return db
}

func DBMigration(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Patient{})
	db.AutoMigrate(&Hospital{})
	db.AutoMigrate(&HospitalStaff{})
	db.AutoMigrate(&Bed{})
	db.AutoMigrate(&BedReservation{})
	db.AutoMigrate(&Doctor{})
	db.AutoMigrate(&Policlinic{})
	db.AutoMigrate(&Practice{})
	db.AutoMigrate(&CheckupReservation{})
	db.AutoMigrate(&KamarRsTeam{})
}

type Bed struct {
	gorm.Model
	NamaTempatTidur string
	Ruangan         string
	Kelas           string
	Status          string
	HospitalID      uint
	// BedReservation  BedReservation `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Hospital struct {
	gorm.Model
	KodeRs            string
	Nama              string
	Foto              string
	Alamat            string
	Provinsi          string
	KabupatenKota     string
	Kecamatan         string
	KodePos           string
	NoTelpon          string
	Email             string
	KelasRs           string
	Pengelola         string
	JumlahTempatTidur int
	StatusPenggunaan  string
	BiayaPendaftaran  int
	HospitalStaffs    []HospitalStaff `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Beds              []Bed           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Policlinics       []Policlinic    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type BedReservation struct {
	gorm.Model
	HospitalId       uint
	StatusPasien     string
	BiayaRegistrasi  int
	KodeDaftar       string
	PaymentMethod    string
	LinkPembayaran   string
	QrString         string
	StatusPembayaran string
	PatientID        uint
	BedID            uint
	// Patient          Patient `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Patient struct {
	gorm.Model
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
	// BedReservationID        uint
	BedReservation BedReservation
	// CheckupReservation CheckupReservation
}

type CheckupReservation struct {
	gorm.Model
	PatientID  uint
	PracticeID uint
	Patient    Patient  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Practice   Practice `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Practice struct {
	gorm.Model
	TanggalPraktik      string
	KuotaHarian         int
	Status              string
	PoliclinicID        uint
	CheckupReservations []CheckupReservation `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Policlinic struct {
	gorm.Model
	NamaPoli   string
	JamPraktik string
	HospitalID uint
	Practices  []Practice `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Doctors    []Doctor   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Doctor struct {
	gorm.Model
	Nama         string
	Spesialis    string
	Email        string
	NoTelpon     string
	Foto         string
	PoliclinicID uint
}

type HospitalStaff struct {
	gorm.Model
	Nama       string
	Email      string
	KataSandi  string
	Peran      string
	HospitalID uint
}

type User struct {
	gorm.Model
	Nama      string
	Email     string `gorm:"unique"`
	Nokk      string
	Nik       string
	KataSandi string
	NoTelpon  string
	Patients  []Patient `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type KamarRsTeam struct {
	gorm.Model
	Email     string
	KataSandi string
	Peran     string
}
