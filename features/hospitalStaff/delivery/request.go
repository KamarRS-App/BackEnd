package delivery

import hospitalstaff "kamarRS/features/hospitalStaff"

type HospitalStaffRequest struct {
	Nama       string `json:"nama" form:"nama"`
	Email      string `json:"email" form:"email"`
	KataSandi  string `json:"kata_sandi" form:"kata_sandi"`
	HospitalID uint   `json:"hospital_id" form:"hospital_id"`
}

func (req *HospitalStaffRequest) reqToCore() hospitalstaff.HospitalStaffCore {
	return hospitalstaff.HospitalStaffCore{
		Nama:       req.Nama,
		Email:      req.Email,
		KataSandi:  req.KataSandi,
		HospitalID: req.HospitalID,
	}

}
