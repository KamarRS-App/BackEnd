package delivery

type DoctorResponse struct {
	ID       uint   `json:"id"`
	Nama     string `json:"nama"`
	Bidang   string `json:"bidang"`
	Email    string `json:"email"`
	NoTelpon string `json:"no_telpon"`
}
