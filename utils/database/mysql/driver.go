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
	Nama_Tempat_Tidur string
	Ruangan           string
	Kelas             string
	Status            string
	HospitalID        uint
	BedReservation    BedReservation
}

type Hospital struct {
	gorm.Model
	Kode_Rs             string
	Nama                string
	Foto                string
	Alamat              string
	Provinsi            string
	Kabupaten_Kota      string
	Kecamatan           string
	No_Telpon           string
	Email               string
	Kelas_Rs            string
	Pengelola           string
	Jumlah_Tempat_Tidur int
	Status_Penggunaan   string
	Biaya_Pendaftaran   int
	HospitalStaffs      []HospitalStaff
	Beds                []Bed
	Policlinics         []Policlinic
}

type BedReservation struct {
	gorm.Model
	Hospital_Id       uint
	Status_Pasien     string
	Biaya_Registrasi  int
	Order_Id          string
	Payment_Method    string
	Link_Pembayaran   string
	Status_Pembayaran string
	PatientID         uint
	BedID             uint
	Patient           Patient
}

type Patient struct {
	gorm.Model
	No_Kk                   string
	Nik                     string
	Nama_Pasien             string
	Jenis_Kelamin           string
	Tanggal_Lahir           string
	Usia                    int
	Nama_Wali               string
	Email_Wali              string
	No_Telpon_Wali          string
	Alamat_Ktp              string
	Provinsi_Ktp            string
	Kabupaten_Kota_Ktp      string
	Alamat_Domisili         string
	Provinsi_Domisili       string
	Kabupaten_Kota_Domisili string
	No_Bpjs                 string
	Kelas_Bpjs              string
	Foto_Ktp                string
	Foto_Bpjs               string
	UserID                  uint
	// BedReservationID        uint
	// BedReservation     BedReservation
	// CheckupReservation CheckupReservation
}

type CheckupReservation struct {
	gorm.Model
	PatientID  uint
	PracticeID uint
	Patient    Patient
	Practice   Practice
}

type Practice struct {
	gorm.Model
	Tanggal_Praktik     string
	Kuota_Harian        int
	Status              string
	PoliclinicID        uint
	CheckupReservations []CheckupReservation
}

type Policlinic struct {
	gorm.Model
	Nama_Poli   string
	Jam_Praktik string
	HospitalID  uint
	DoctorID    uint
	Practices   []Practice
	Doctor      Doctor
}

type Doctor struct {
	gorm.Model
	Nama        string
	Spesialis   string
	Email       string
	No_Telpon   string
	Foto        string
	Policlinics []Policlinic
}

type HospitalStaff struct {
	gorm.Model
	Nama       string
	Email      string
	Kata_Sandi string
	Peran      string
	HospitalID uint
}

type User struct {
	gorm.Model
	Nama       string
	Email      string
	No_kk      string
	Nik        string
	Kata_Sandi string
	No_Telpon  string
	Patients   []Patient
}

type KamarRsTeam struct {
	gorm.Model
	Email    string
	Password string
	Peran    string
}
