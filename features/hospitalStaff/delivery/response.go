package delivery

import (
	hospitalstaff "github.com/KamarRS-App/features/hospitalStaff"
)

type HospitalStaffResponse struct {
	ID         uint   `json:"id"`
	Nama       string `json:"nama"`
	Email      string `json:"email"`
	Peran      string `json:"peran"`
	HospitalID uint   `json:"hospital_id"`
}

func StaffCoreToStaffRespon(dataCore hospitalstaff.HospitalStaffCore) HospitalStaffResponse { // data user core yang ada di controller yang memanggil user repositoCorePatient
	return HospitalStaffResponse{
		ID:         dataCore.ID,
		Nama:       dataCore.Nama,
		Email:      dataCore.Email,
		Peran:      dataCore.Peran,
		HospitalID: dataCore.HospitalID,
	}
}
func ListStaffCoreToStaffRespon(dataCore []hospitalstaff.HospitalStaffCore) []HospitalStaffResponse { //data user.core data yang diambil dari entities ke respon struct
	var ResponData []HospitalStaffResponse

	for _, value := range dataCore { //memanggil paramete data core yang berisi data user core
		ResponData = append(ResponData, StaffCoreToStaffRespon(value)) // mengambil data mapping dari user core to respon
	}
	return ResponData
}
