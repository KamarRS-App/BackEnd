package delivery

type HospitalStaffResponse struct {
	ID         uint   `json:"id"`
	Nama       string `json:"nama"`
	Email      string `json:"email"`
	KataSandi  string `json:"kata_sandi"`
	HospitalID uint   `json:"hospital_id"`
}
