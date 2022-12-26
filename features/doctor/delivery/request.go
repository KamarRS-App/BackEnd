package delivery

import "github.com/KamarRS-App/features/doctor"

type DoctorRequest struct {
	Nama      string `json:"nama" form:"nama"`
	Spesialis string `json:"spesialis" form:"spesialis"`
	Email     string `json:"email" form:"email"`
	NoTelpon  string `json:"no_telpon" form:"no_telpon"`
	Foto      string `json:"foto" form:"foto"`
}

func (req *DoctorRequest) reqToCore() doctor.DoctorCore {
	return doctor.DoctorCore{
		Nama:      req.Nama,
		Spesialis: req.Spesialis,
		Email:     req.Email,
		NoTelpon:  req.NoTelpon,
		Foto:      req.Foto,
	}
}
