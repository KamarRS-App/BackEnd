package delivery

type HospitalStaffRequest struct {
	Nama       string `json:"nama" form:"nama"`
	Email      string `json:"email" form:"email"`
	KataSandi  string `json:"kata_sandi" form:"kata_sandi"`
	HospitalID uint   `json:"hospital_id" form:"hospital_id"`
}
