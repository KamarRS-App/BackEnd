package delivery

import (
	"kamarRS/features/policlinic"
)

type ResponsePoliclinic struct {
	ID         uint   `json:"policlinic_id"`
	NamaPoli   string `json:"nama_poli"`
	JamPraktik string `json:"jam_praktik"`
	HospitalID uint   `json:"hospital_id"`
	DoctorID   uint   `json:"doctor_id"`
}

func PoliCoreToPoliRespon(dataCore policlinic.CorePoliclinic) ResponsePoliclinic { // data user core yang ada di controller yang memanggil user repositoCorePatient
	return ResponsePoliclinic{
		ID:         dataCore.ID,
		NamaPoli:   dataCore.NamaPoli,
		JamPraktik: dataCore.JamPraktik,
		HospitalID: dataCore.HospitalID,
		DoctorID:   dataCore.DoctorID,
	}
}
func ListpoliCoreToPoliRespon(dataCore []policlinic.CorePoliclinic) []ResponsePoliclinic { //data user.core data yang diambil dari entities ke respon struct
	var ResponData []ResponsePoliclinic

	for _, value := range dataCore { //memanggil paramete data core yang berisi data user core
		ResponData = append(ResponData, PoliCoreToPoliRespon(value)) // mengambil data mapping dari user core to respon
	}
	return ResponData
}
