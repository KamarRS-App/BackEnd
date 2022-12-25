package delivery

import (
	"kamarRS/features/doctor"
)

type DoctorResponse struct {
	ID       uint   `json:"id"`
	Nama     string `json:"nama"`
	Bidang   string `json:"bidang"`
	Email    string `json:"email"`
	NoTelpon string `json:"no_telpon"`
}

// -----------------Doctor--------------------
func fromCore(dataCore doctor.DoctorCore) DoctorResponse {
	return DoctorResponse{
		ID:       dataCore.ID,
		Nama:     dataCore.Nama,
		Bidang:   dataCore.Bidang,
		Email:    dataCore.Email,
		NoTelpon: dataCore.NoTelpon,
	}
}

// data dari core ke response
func fromCoreList(dataCore []doctor.DoctorCore) []DoctorResponse {
	var dataResponse []DoctorResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}

//-----------------------------------------------------
