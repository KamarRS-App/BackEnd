package delivery

import "kamarRS/features/doctor"

type DoctorRequest struct {
	Nama     string `json:"nama" form:"nama"`
	Bidang   string `json:"bidang" form:"bidang"`
	Email    string `json:"email" form:"email"`
	NoTelpon string `json:"no_telpon" form:"no_telpon"`
}

func (req *DoctorRequest) reqToCore() doctor.DoctorCore {
	return doctor.DoctorCore{
		Nama:     req.Nama,
		Bidang:   req.Bidang,
		Email:    req.Email,
		NoTelpon: req.NoTelpon,
	}
}
