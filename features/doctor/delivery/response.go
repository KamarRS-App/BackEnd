package delivery

import (
	"github.com/KamarRS-App/KamarRS-App/features/doctor"
)

type DoctorResponse struct {
	ID           uint   `json:"id"`
	Nama         string `json:"nama"`
	Spesialis    string `json:"spesialis"`
	Email        string `json:"email"`
	NoTelpon     string `json:"no_telpon"`
	Foto         string `json:"foto"`
	PoliclinicID uint   `json:"policlinic_id"`
}

// -----------------Doctor--------------------
func FromCore(dataCore doctor.DoctorCore) DoctorResponse {
	return DoctorResponse{
		ID:           dataCore.ID,
		Nama:         dataCore.Nama,
		Spesialis:    dataCore.Spesialis,
		Email:        dataCore.Email,
		NoTelpon:     dataCore.NoTelpon,
		Foto:         dataCore.Foto,
		PoliclinicID: dataCore.PoliclinicID,
	}
}

// data dari core ke response
func FromCoreList(dataCore []doctor.DoctorCore) []DoctorResponse {
	var dataResponse []DoctorResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, FromCore(v))
	}
	return dataResponse
}

//-----------------------------------------------------
