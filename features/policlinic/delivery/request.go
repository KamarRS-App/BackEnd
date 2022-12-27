package delivery

import (
	"github.com/KamarRS-App/features/policlinic"
)

type RequestPoliclinic struct {
	NamaPoli   string `json:"nama_poli" form:"nama_poli"`
	JamPraktik string `json:"jam_praktik" form:"jam_praktik"`
	HospitalID uint   `json:"hospital_id" form:"hospital_id"`
	DoctorID   uint   `json:"doctor_id" form:"doctor_id"`
}

func (req *RequestPoliclinic) ToCore() policlinic.CorePoliclinic {
	return policlinic.CorePoliclinic{
		NamaPoli:   req.NamaPoli,
		JamPraktik: req.JamPraktik,
		HospitalID: req.HospitalID,
		DoctorID:   req.DoctorID,
	}

}
