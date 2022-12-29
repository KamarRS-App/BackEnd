package delivery

import (
	"github.com/KamarRS-App/KamarRS-App/features/hospitalstaff"
)

type HospitalStaffResponse struct {
	ID         uint   `json:"id"`
	Nama       string `json:"nama"`
	Email      string `json:"email"`
	Peran      string `json:"peran"`
	HospitalID uint   `json:"hospital_id"`
	// Hospital   HospitalResponse
}

func StaffCoreToStaffRespon(dataCore hospitalstaff.HospitalStaffCore) HospitalStaffResponse { // data user core yang ada di controller yang memanggil user repositoCorePatient
	return HospitalStaffResponse{
		ID:         dataCore.ID,
		Nama:       dataCore.Nama,
		Email:      dataCore.Email,
		Peran:      dataCore.Peran,
		HospitalID: dataCore.HospitalID,
		// Hospital: HospitalResponse{
		// 	Nama: dataCore.Hospital.Nama,
		// },
	}
}
func ListStaffCoreToStaffRespon(dataCore []hospitalstaff.HospitalStaffCore) []HospitalStaffResponse { //data user.core data yang diambil dari entities ke respon struct
	var ResponData []HospitalStaffResponse

	for _, value := range dataCore { //memanggil paramete data core yang berisi data user core
		ResponData = append(ResponData, StaffCoreToStaffRespon(value)) // mengambil data mapping dari user core to respon
	}
	return ResponData
}

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
