package delivery

import "github.com/KamarRS-App/KamarRS-App/features/hospitalstaff"

type HospitalStaffRequest struct {
	Nama       string `json:"nama" form:"nama"`
	Email      string `json:"email" form:"email"`
	KataSandi  string `json:"kata_sandi" form:"kata_sandi"`
	Peran      string `json:"peran" form:"peran"`
	HospitalID uint   `json:"hospital_id" form:"hospital_id"`
}

func (req *HospitalStaffRequest) reqToCore() hospitalstaff.HospitalStaffCore {
	return hospitalstaff.HospitalStaffCore{
		Nama:       req.Nama,
		Email:      req.Email,
		KataSandi:  req.KataSandi,
		Peran:      req.Peran,
		HospitalID: req.HospitalID,
	}

}
