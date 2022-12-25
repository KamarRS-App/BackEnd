package delivery

type DoctorRequest struct {
	Nama     string `json:"nama" form:"nama"`
	Bidang   string `json:"bidang" form:"bidang"`
	Email    string `json:"email" form:"email"`
	NoTelpon string `json:"no_telpon" form:"no_telpon"`
}
