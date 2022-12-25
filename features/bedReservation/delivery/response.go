package delivery

import (
	bedreservation "kamarRS/features/bedReservation"
)

type BedReservationResponse struct {
	ID               uint            `json:"id"`
	StatusPasien     string          `json:"status_pasien"`
	BiayaRegistrasi  int             `json:"biaya_registrasi"`
	OrderID          string          `json:"order_id"`
	PaymentMethod    string          `json:"payment_method"`
	LinkPembayaran   string          `json:"link_pembayaran"`
	StatusPembayaran string          `json:"status_pembayaran"`
	HospitalID       uint            `json:"hospital_id"`
	BedID            uint            `json:"bed_id"`
	Patients         PatientResponse `json:"patients"`
}

type PatientResponse struct {
	ID                    uint   `json:"id"`
	NoKk                  string `json:"no_kk"`
	Nik                   string `json:"nik"`
	NamaPasien            string `json:"nama_pasien"`
	JenisKelamin          string `json:"jenis_kelamin"`
	TanggalLahir          string `json:"tanggal_lahir"`
	Usia                  int    `json:"usia" form:"usia"`
	NamaWali              string `json:"nama_wali"`
	EmailWali             string `json:"email_wali"`
	NoTelponWali          string `json:"no_telpon_wali"`
	AlamatKtp             string `json:"alamat_ktp"`
	ProvinsiKtp           string `json:"provinsi_ktp"`
	KabupatenKotaKtp      string `json:"kabupaten_kota_ktp"`
	AlamatDomisili        string `json:"alamat_domisili"`
	ProvinsiDomisili      string `json:"provinsi_domisili"`
	KabupatenKotaDomisili string `json:"kabupaten_kota_domisili"`
	NoBpjs                string `json:"no_bpjs"`
	KelasBpjs             string `json:"kelas_bpjs"`
	FotoKtp               string `json:"foto_ktp"`
	FotoBpjs              string `json:"foto_bpjs"`
	// BedReservation          BedReservation
	// CheckupReservation      CheckupReservation
}

// -----------------Bed Reserve from Core--------------------------------
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

//-----------------------Patient from Core---------------------------------------

func PatientCoreToPatientRespon(dataCore bedreservation.PatientCore) PatientResponse { // data user core yang ada di controller yang memanggil user repositoCorePatient
	return PatientResponse{
		ID:                    dataCore.ID,
		NoKk:                  dataCore.NoKk,
		Nik:                   dataCore.Nik,
		NamaPasien:            dataCore.NamaPasien,
		JenisKelamin:          dataCore.JenisKelamin,
		TanggalLahir:          dataCore.TanggalLahir,
		Usia:                  dataCore.Usia,
		NamaWali:              dataCore.NamaWali,
		EmailWali:             dataCore.EmailWali,
		NoTelponWali:          dataCore.NoTelponWali,
		AlamatKtp:             dataCore.AlamatKtp,
		ProvinsiKtp:           dataCore.ProvinsiKtp,
		KabupatenKotaKtp:      dataCore.KabupatenKotaKtp,
		AlamatDomisili:        dataCore.AlamatDomisili,
		ProvinsiDomisili:      dataCore.ProvinsiDomisili,
		KabupatenKotaDomisili: dataCore.KabupatenKotaDomisili,
		NoBpjs:                dataCore.NoBpjs,
		KelasBpjs:             dataCore.KelasBpjs,
		FotoKtp:               dataCore.FotoKtp,
		FotoBpjs:              dataCore.FotoBpjs,
	}
}
func ListpatientCoreTopatientRespon(dataCore []bedreservation.PatientCore) []PatientResponse { //data user.core data yang diambil dari entities ke respon struct
	var ResponData []PatientResponse

	for _, value := range dataCore { //memanggil paramete data core yang berisi data user core
		ResponData = append(ResponData, PatientCoreToPatientRespon(value)) // mengambil data mapping dari user core to respon
	}
	return ResponData
}

//----------------------------------------------------------------------------------
