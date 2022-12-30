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
	// Patient          PatientCore
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

type ServiceInterface interface {
	Create(input BedReservationCore) (data BedReservationCore, err error)
	GetPayment(kodeDaftar string) (data BedReservationCore, err error)
	CreatePayment(input BedReservationCore) (data BedReservationCore, err error)
}

type RepositoryInterface interface {
	Create(input BedReservationCore) (data BedReservationCore, err error)
	GetPayment(kodeDaftar string) (data BedReservationCore, err error)
	CreatePayment(input BedReservationCore) (data BedReservationCore, err error)
}
