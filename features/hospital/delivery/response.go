package delivery

import (
	"github.com/KamarRS-App/features/hospital"
)

type HospitalResponse struct {
	ID                uint   `json:"id"`
	KodeRs            string `json:"kode_rs"`
	Nama              string `json:"nama"`
	Foto              string `json:"foto"`
	Alamat            string `json:"alamat"`
	Provinsi          string `json:"provinsi"`
	KabupatenKota     string `json:"kabupaten_kota"`
	Kecamatan         string `json:"kecamatan"`
	NoTelpon          string `json:"no_telpon"`
	Email             string `json:"email"`
	KelasRs           string `json:"kelas_rs"`
	PemilikPengelola  string `json:"pemilik_pengelola"`
	JumlahTempatTidur int    `json:"jumlah_tempat_tidur"`
	StatusPenggunaan  string `json:"status_penggunaan"`
	BiayaRegistrasi   int    `json:"biaya_registrasi"`
}

// -----------------Hospital--------------------
func FromCore(dataCore hospital.HospitalCore) HospitalResponse {
	return HospitalResponse{
		ID:                dataCore.ID,
		KodeRs:            dataCore.KodeRs,
		Nama:              dataCore.Nama,
		Foto:              dataCore.Foto,
		Alamat:            dataCore.Alamat,
		Provinsi:          dataCore.Provinsi,
		KabupatenKota:     dataCore.KabupatenKota,
		Kecamatan:         dataCore.Kecamatan,
		NoTelpon:          dataCore.NoTelpon,
		Email:             dataCore.Email,
		KelasRs:           dataCore.KelasRs,
		PemilikPengelola:  dataCore.PemilikPengelola,
		JumlahTempatTidur: dataCore.JumlahTempatTidur,
		StatusPenggunaan:  dataCore.StatusPenggunaan,
		BiayaRegistrasi:   dataCore.BiayaRegistrasi,
	}
}

// data dari core ke response
func FromCoreList(dataCore []hospital.HospitalCore) []HospitalResponse {
	var dataResponse []HospitalResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, FromCore(v))
	}
	return dataResponse
}

//-----------------------------------------------------
