package delivery

type BedRequest struct {
	NamaTempatTidur string `json:"nama_tempat_tidur" form:"nama_tempat_tidur"`
	Ruangan         string `json:"ruangan" form:"ruangan"`
	Kelas           string `json:"kelas" form:"kelas"`
	Status          string `json:"status" form:"status"`
	HospitalID      uint   `json:"hospital_id" form:"hospital_id"`
}
