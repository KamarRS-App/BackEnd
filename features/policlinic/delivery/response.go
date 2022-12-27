package delivery

import (
	"github.com/KamarRS-App/features/policlinic"
)

type ResponsePoliclinic struct {
	ID         uint           `json:"id"`
	NamaPoli   string         `json:"nama_poli"`
	JamPraktik string         `json:"jam_praktik"`
	HospitalID uint           `json:"hospital_id"`
	Doctor     ResponseDoctor `json:"doctor"`
}

type ResponseDoctor struct {
	ID        uint   `json:"id"`
	Nama      string `json:"nama"`
	Spesialis string `json:"spesialis"`
	Email     string `json:"email"`
	NoTelpon  string `json:"no_telpon"`
	Foto      string `json:"foto"`
}

func FromCore(dataCore policlinic.CorePoliclinic) ResponsePoliclinic { // data user core yang ada di controller yang memanggil user repositoCorePatient
	return ResponsePoliclinic{
		ID:         dataCore.ID,
		NamaPoli:   dataCore.NamaPoli,
		JamPraktik: dataCore.JamPraktik,
		HospitalID: dataCore.HospitalID,
		Doctor: ResponseDoctor{
			ID:        dataCore.Doctor.ID,
			Nama:      dataCore.Doctor.Nama,
			Spesialis: dataCore.Doctor.Spesialis,
			Email:     dataCore.Doctor.Email,
			NoTelpon:  dataCore.Doctor.NoTelpon,
			Foto:      dataCore.Doctor.Foto,
		},
	}
}
func FromCoreList(dataCore []policlinic.CorePoliclinic) []ResponsePoliclinic { //data user.core data yang diambil dari entities ke respon struct
	var ResponData []ResponsePoliclinic

	for _, value := range dataCore { //memanggil paramete data core yang berisi data user core
		ResponData = append(ResponData, FromCore(value)) // mengambil data mapping dari user core to respon
	}
	return ResponData
}
