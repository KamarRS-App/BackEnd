package bedreservation

type BedReservationCore struct {
	ID               uint
	StatusPasien     string
	BiayaRegistrasi  int
	KodeDaftar       string
	PaymentMethod    string
	LinkPembayaran   string
	TransactionId    string
	VirtualAccount   string
	BankPenerima     string
	WaktuKedaluarsa  string
	QrString         string
	StatusPembayaran string
	HospitalID       uint
	PatientID        uint
	BedID            uint
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
	BedReservation        BedReservationCore
}

type ServiceInterface interface {
	Create(input BedReservationCore, userId uint) (data BedReservationCore, err error)
	GetPayment(kodeDaftar string) (data BedReservationCore, err error)
	CreatePayment(input BedReservationCore) (data BedReservationCore, err error)
	PaymentNotif(callback BedReservationCore) (err error)
	GetRegistrations(page, limit, hospitalId int) (data []BedReservationCore, totalpage int, err error)
	GetById(bedResId uint) (data BedReservationCore, err error)
	Delete(bedResId uint) error
	UpdateBedReservation(input BedReservationCore) error
}

type RepositoryInterface interface {
	Create(input BedReservationCore, userId uint) (data BedReservationCore, err error)
	GetPayment(kodeDaftar string) (data BedReservationCore, err error)
	CreatePayment(input BedReservationCore) (data BedReservationCore, err error)
	PaymentNotif(callback BedReservationCore) (err error)
	GetRegistrations(limit, offset, hospitalId int) (data []BedReservationCore, totalpage int, err error)
	GetById(bedResId uint) (data BedReservationCore, err error)
	Delete(bedResId uint) error
	UpdateBedReservation(input BedReservationCore) error
}
