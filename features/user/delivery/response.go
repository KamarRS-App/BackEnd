package delivery

import (
	"github.com/KamarRS-App/KamarRS-App/features/user"
)

type ResponseUser struct {
	ID       uint   `json:"id"`
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Nokk     string `json:"no_kk"`
	Nik      string `json:"nik"`
	NoTelpon string `json:"no_telpon"`
	// Patients []ResponsePatient
}

func UserCoreToUserRespon(dataCore user.CoreUser) ResponseUser { // data user core yang ada di controller yang memanggil user repository
	return ResponseUser{
		ID:       dataCore.ID,
		Nama:     dataCore.Nama,
		Email:    dataCore.Email,
		Nokk:     dataCore.Nokk,
		Nik:      dataCore.Nik,
		NoTelpon: dataCore.NoTelpon,
		// Patients: LoadpatientCoretoResponse(dataCore.Patients),
	}
}

func ListUserCoreToUserRespon(dataCore []user.CoreUser) []ResponseUser { //data user.core data yang diambil dari entities ke respon struct
	var ResponData []ResponseUser

	for _, value := range dataCore { //memanggil paramete data core yang berisi data user core
		ResponData = append(ResponData, UserCoreToUserRespon(value)) // mengambil data mapping dari user core to respon
	}
	return ResponData
}

//--------------------------------------------------------------------------------------------------------------
// type ResponsePatient struct {
// 	ID                    uint
// 	NoKk                  string `json:"no_kk"`
// 	Nik                   string `json:"nik"`
// 	NamaPasien            string `json:"nama_pasien"`
// 	JenisKelamin          string `json:"jenis_kelamin"`
// 	TanggalLahir          string `json:"tanggal_lahir"`
// 	Usia                  int    `json:"usia" form:"usia"`
// 	NamaWali              string `json:"nama_wali"`
// 	EmailWali             string `json:"email_wali"`
// 	NoTelponWali          string `json:"no_telpon_wali"`
// 	AlamatKtp             string `json:"alamat_ktp"`
// 	ProvinsiKtp           string `json:"provinsi_ktp"`
// 	KabupatenKotaKtp      string `json:"kabupaten_kota_ktp"`
// 	AlamatDomisili        string `json:"alamat_domisili"`
// 	ProvinsiDomisili      string `json:"provinsi_domisili"`
// 	KabupatenKotaDomisili string `json:"kabupaten_kota_domisili"`
// 	NoBpjs                string `json:"no_bpjs"`
// 	KelasBpjs             string `json:"kelas_bpjs"`
// 	FotoKtp               string `json:"foto_ktp"`
// 	FotoBpjs              string `json:"foto_bpjs"`
// 	UserID                uint   `json:"user_id"`
// 	// BedReservation          BedReservation
// 	// CheckupReservation      CheckupReservation
// }

// func LoadpatientCoretoResponse(core []user.CorePatient) []ResponsePatient {
// 	var res []ResponsePatient
// 	for _, v := range core {
// 		res = append(res, PatientCoreToPatientRespon(v))
// 	}
// 	return res

// }

// func PatientCoreToPatientRespon(dataCore user.CorePatient) ResponsePatient { // data user core yang ada di controller yang memanggil user repositoCorePatient
// 	return ResponsePatient{
// 		ID:                    dataCore.ID,
// 		NoKk:                  dataCore.NoKk,
// 		Nik:                   dataCore.Nik,
// 		NamaPasien:            dataCore.NamaPasien,
// 		JenisKelamin:          dataCore.JenisKelamin,
// 		TanggalLahir:          dataCore.TanggalLahir,
// 		Usia:                  dataCore.Usia,
// 		NamaWali:              dataCore.NamaWali,
// 		EmailWali:             dataCore.EmailWali,
// 		NoTelponWali:          dataCore.NoTelponWali,
// 		AlamatKtp:             dataCore.AlamatKtp,
// 		ProvinsiKtp:           dataCore.ProvinsiKtp,
// 		KabupatenKotaKtp:      dataCore.KabupatenKotaKtp,
// 		AlamatDomisili:        dataCore.AlamatDomisili,
// 		ProvinsiDomisili:      dataCore.ProvinsiDomisili,
// 		KabupatenKotaDomisili: dataCore.KabupatenKotaDomisili,
// 		NoBpjs:                dataCore.NoBpjs,
// 		KelasBpjs:             dataCore.KelasBpjs,
// 		FotoKtp:               dataCore.FotoKtp,
// 		FotoBpjs:              dataCore.FotoBpjs,
// 		UserID:                dataCore.UserID,
// 	}
// }

//--------------------------------------------------------------------------------------------------------------
// func UserCoreToUserResponNoPreload(dataCore user.CoreUser) ResponseUser { // data user core yang ada di controller yang memanggil user repository
// 	return ResponseUser{
// 		ID:       dataCore.ID,
// 		Nama:     dataCore.Nama,
// 		Email:    dataCore.Email,
// 		Nokk:     dataCore.Nokk,
// 		Nik:      dataCore.Nik,
// 		NoTelpon: dataCore.NoTelpon,
// 	}
// }

// func ListUserCoreToUserResponNoPreload(dataCore []user.CoreUser) []ResponseUser { //data user.core data yang diambil dari entities ke respon struct
// 	var ResponData []ResponseUser

// 	for _, value := range dataCore { //memanggil paramete data core yang berisi data user core
// 		ResponData = append(ResponData, UserCoreToUserResponNoPreload(value)) // mengambil data mapping dari user core to respon
// 	}
// 	return ResponData
// }
