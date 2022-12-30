package delivery

import (
	bedreservation "github.com/KamarRS-App/KamarRS-App/features/bedReservation"
)

type BedReservationRequest struct {
	StatusPasien     string `json:"status_pasien" form:"status_pasien"`
	BiayaRegistrasi  int    `json:"biaya_registrasi" form:"biaya_registrasi"`
	KodeDaftar       string `json:"kode_daftar" form:"kode_daftar"`
	PaymentMethod    string `json:"payment_method" form:"payment_method"`
	LinkPembayaran   string `json:"link_pembayaran" form:"link_pembayaran"`
	QrString         string `json:"qr_string" form:"qr_string"`
	StatusPembayaran string `json:"status_pembayaran" form:"status_pembayaran"`
	HospitalID       uint   `json:"hospital_id" form:"hospital_id"`
	PatientID        uint   `json:"patient_id" form:"patient_id"`
	BedID            uint   `json:"bed_id" form:"bed_id"`
}

func (req *BedReservationRequest) reqToCore() bedreservation.BedReservationCore {
	return bedreservation.BedReservationCore{
		StatusPasien:     req.StatusPasien,
		BiayaRegistrasi:  req.BiayaRegistrasi,
		KodeDaftar:       req.KodeDaftar,
		PaymentMethod:    req.PaymentMethod,
		LinkPembayaran:   req.LinkPembayaran,
		QrString:         req.QrString,
		StatusPembayaran: req.StatusPembayaran,
		HospitalID:       req.HospitalID,
		PatientID:        req.PatientID,
		BedID:            req.BedID,
	}
}
