package delivery

type BedResponse struct {
	ID              uint   `json:"id"`
	NamaTempatTidur string `json:"nama_tempat_tidur"`
	Ruangan         string `json:"ruangan"`
	Kelas           string `json:"kelas"`
	Status          string `json:"status"`
	HospitalID      uint   `json:"hospital_id"`
}
