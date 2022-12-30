package delivery

import (
	bedreservation "github.com/KamarRS-App/KamarRS-App/features/bedReservation"
)

type BedReservationRequest struct {
	StatusPasien     string `json:"status_pasien" form:"status_pasien"`
	BiayaRegistrasi  int    `json:"biaya_registrasi" form:"biaya_registrasi"`
	KodeDaftar       string `json:"kode_daftar" form:"kode_daftar"`
	PaymentMethod    string `json:"metode_pembayaran" form:"metode_pembayaran"`
	LinkPembayaran   string `json:"link_pembayaran" form:"link_pembayaran"`
	TransactionId    string `json:"transaction_id" form:"transaction_id"`
	VirtualAccount   string `json:"virtual_account" form:"virtual_account"`
	BankPenerima     string `json:"bank_penerima" form:"bank_penerima"`
	WaktuKedaluarsa  string `json:"waktu_kedaluarsa" form:"waktu_kedaluarsa"`
	QrString         string `json:"qr_string" form:"qr_string"`
	StatusPembayaran string `json:"status_pembayaran" form:"status_pembayaran"`
	HospitalID       uint   `json:"hospital_id" form:"hospital_id"`
	PatientID        uint   `json:"patient_id" form:"patient_id"`
}

type UpdateMidtrans struct {
	KodeDaftar       string `json:"order_id" form:"order_id"`
	StatusPembayaran string `json:"status_pembayaran" form:"status_pembayaran"`
}

func (req *BedReservationRequest) reqToCore() bedreservation.BedReservationCore {
	return bedreservation.BedReservationCore{
		StatusPasien:     req.StatusPasien,
		BiayaRegistrasi:  req.BiayaRegistrasi,
		KodeDaftar:       req.KodeDaftar,
		PaymentMethod:    req.PaymentMethod,
		LinkPembayaran:   req.LinkPembayaran,
		TransactionId:    req.TransactionId,
		VirtualAccount:   req.VirtualAccount,
		BankPenerima:     req.BankPenerima,
		WaktuKedaluarsa:  req.WaktuKedaluarsa,
		QrString:         req.QrString,
		StatusPembayaran: req.StatusPembayaran,
		HospitalID:       req.HospitalID,
		PatientID:        req.PatientID,
	}
}

func (req *UpdateMidtrans) reqToCore() bedreservation.BedReservationCore {
	return bedreservation.BedReservationCore{
		KodeDaftar:       req.KodeDaftar,
		StatusPembayaran: req.StatusPembayaran,
	}
}
