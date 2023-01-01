package delivery

import (
	"github.com/KamarRS-App/KamarRS-App/features/policlinic"
)

type ResponsePoliclinic struct {
	ID         uint             `json:"id"`
	NamaPoli   string           `json:"nama_poli"`
	JamPraktik string           `json:"jam_praktik"`
	HospitalID uint             `json:"hospital_id"`
	Doctor     []ResponseDoctor `json:"doctor"`
}

type ResponseDoctor struct {
	ID        uint   `json:"id"`
	Nama      string `json:"nama"`
	Spesialis string `json:"spesialis"`
	Email     string `json:"email"`
	NoTelpon  string `json:"no_telpon"`
	Foto      string `json:"foto"`
}

// ---------------------------
type ResponseAllPoli struct {
	ID         uint   `json:"id"`
	NamaPoli   string `json:"nama_poli"`
	JamPraktik string `json:"jam_praktik"`
	HospitalID uint   `json:"hospital_id"`
}

func FromCore(dataCore policlinic.CorePoliclinic) ResponsePoliclinic { // data user core yang ada di controller yang memanggil user repositoCorePatient
	return ResponsePoliclinic{
		ID:         dataCore.ID,
		NamaPoli:   dataCore.NamaPoli,
		JamPraktik: dataCore.JamPraktik,
		HospitalID: dataCore.HospitalID,
		Doctor:     FromCoreListD(dataCore.Doctor),
	}
}

func FromCoreList(dataCore []policlinic.CorePoliclinic) []ResponsePoliclinic { //data user.core data yang diambil dari entities ke respon struct
	var ResponData []ResponsePoliclinic

	for _, value := range dataCore { //memanggil paramete data core yang berisi data user core
		ResponData = append(ResponData, FromCore(value)) // mengambil data mapping dari user core to respon
	}
	return ResponData
}

func FromCoreD(dataCore policlinic.CoreDoctor) ResponseDoctor { // data user core yang ada di controller yang memanggil user repositoCorePatient
	return ResponseDoctor{
		ID:        dataCore.ID,
		Nama:      dataCore.Nama,
		Spesialis: dataCore.Spesialis,
		Email:     dataCore.Email,
		NoTelpon:  dataCore.NoTelpon,
		Foto:      dataCore.Foto,
	}
}

func FromCoreListD(dataCore []policlinic.CoreDoctor) []ResponseDoctor { //data user.core data yang diambil dari entities ke respon struct
	var ResponData []ResponseDoctor

	for _, value := range dataCore { //memanggil paramete data core yang berisi data user core
		ResponData = append(ResponData, FromCoreD(value)) // mengambil data mapping dari user core to respon
	}
	return ResponData
}

// All Poli by HospitalID
func FromCoreP(dataCore policlinic.CorePoliclinic) ResponseAllPoli { // data user core yang ada di controller yang memanggil user repositoCorePatient
	return ResponseAllPoli{
		ID:         dataCore.ID,
		NamaPoli:   dataCore.NamaPoli,
		JamPraktik: dataCore.JamPraktik,
		HospitalID: dataCore.HospitalID,
	}
}

func FromCoreListP(dataCore []policlinic.CorePoliclinic) []ResponseAllPoli { //data user.core data yang diambil dari entities ke respon struct
	var ResponData []ResponseAllPoli

	for _, value := range dataCore { //memanggil paramete data core yang berisi data user core
		ResponData = append(ResponData, FromCoreP(value)) // mengambil data mapping dari user core to respon
	}
	return ResponData
}
