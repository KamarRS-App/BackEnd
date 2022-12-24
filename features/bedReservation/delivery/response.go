package delivery

type BedReservationResponse struct {
	ID               uint   `json:"id"`
	StatusPasien     string `json:"status_pasien"`
	BiayaRegistrasi  int    `json:"biaya_registrasi"`
	OrderID          string `json:"order_id"`
	PaymentMethod    string `json:"payment_method"`
	LinkPembayaran   string `json:"link_pembayaran"`
	StatusPembayaran string `json:"status_pembayaran"`
	HospitalID       uint   `json:"hospital_id"`
	PatientID        uint   `json:"patient_id"`
	BedID            uint   `json:"bed_id"`
}
