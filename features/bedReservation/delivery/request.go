package delivery

type BedReservationRequest struct {
	StatusPasien     string `json:"status_pasien" form:"status_pasien"`
	BiayaRegistrasi  int    `json:"biaya_registrasi" form:"status_registrasi"`
	OrderID          string `json:"order_id" form:"order_id"`
	PaymentMethod    string `json:"payment_method" form:"payment_method"`
	LinkPembayaran   string `json:"link_pembayaran" form:"link_pembayaran"`
	StatusPembayaran string `json:"status_pembayaran" form:"status_pembayaran"`
	HospitalID       uint   `json:"hospital_id" form:"hospital_id"`
	PatientID        uint   `json:"patient_id" form:"patient_id"`
	BedID            uint   `json:"bed_id" form:"bed_id"`
}
