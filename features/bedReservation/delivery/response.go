package delivery

import (
	bedreservation "kamarRS/features/bedReservation"
)

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

// -----------------Bed Reserve--------------------------------
func fromCore(dataCore bedreservation.BedReservationCore) BedReservationResponse {
	return BedReservationResponse{
		ID:               dataCore.ID,
		StatusPasien:     dataCore.StatusPasien,
		BiayaRegistrasi:  dataCore.BiayaRegistrasi,
		OrderID:          dataCore.OrderID,
		PaymentMethod:    dataCore.PaymentMethod,
		LinkPembayaran:   dataCore.LinkPembayaran,
		StatusPembayaran: dataCore.StatusPembayaran,
		HospitalID:       dataCore.HospitalID,
		PatientID:        dataCore.PatientID,
		BedID:            dataCore.BedID,
	}
}

// data dari core ke response
func fromCoreList(dataCore []bedreservation.BedReservationCore) []BedReservationResponse {
	var dataResponse []BedReservationResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
