package bedreservation

type BedReservationCore struct {
	ID               uint
	StatusPasien     string
	BiayaRegistrasi  int
	OrderID          string
	PaymentMethod    string
	LinkPembayaran   string
	StatusPembayaran string
	HospitalID       uint
	PatientID        uint
	BedID            uint
}
