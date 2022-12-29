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
}
type HospitalStaffResponsePreload struct {
	ID         uint             `json:"id"`
	Nama       string           `json:"nama"`
	Email      string           `json:"email"`
	Peran      string           `json:"peran"`
	HospitalID uint             `json:"hospital_id"`
	Hospital   HospitalResponse `json:"hospital"`
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
	KodeRs string `json:"kode_rs"`
	Nama   string `json:"nama"`
	Foto   string `json:"foto"`
}

func StaffCoreToStaffResponPreload(dataCore hospitalstaff.HospitalStaffCore) HospitalStaffResponsePreload { // data user core yang ada di controller yang memanggil user repositoCorePatient
	return HospitalStaffResponsePreload{
		ID:         dataCore.ID,
		Nama:       dataCore.Nama,
		Email:      dataCore.Email,
		Peran:      dataCore.Peran,
		HospitalID: dataCore.HospitalID,
		Hospital: HospitalResponse{
			KodeRs: dataCore.Hospital.KodeRs,
			Nama:   dataCore.Hospital.Nama,
			Foto:   dataCore.Hospital.Foto,
		},
	}
}
func ListStaffCoreToStaffResponStaffCoreToStaffResponPreload(dataCore []hospitalstaff.HospitalStaffCore) []HospitalStaffResponsePreload { //data user.core data yang diambil dari entities ke respon struct
	var ResponData []HospitalStaffResponsePreload

	for _, value := range dataCore { //memanggil paramete data core yang berisi data user core
		ResponData = append(ResponData, StaffCoreToStaffResponPreload(value)) // mengambil data mapping dari user core to respon
	}
	return ResponData
}
