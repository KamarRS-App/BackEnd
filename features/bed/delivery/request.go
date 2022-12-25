package delivery

import "kamarRS/features/bed"

type BedRequest struct {
	NamaTempatTidur string `json:"nama_tempat_tidur" form:"nama_tempat_tidur"`
	Ruangan         string `json:"ruangan" form:"ruangan"`
	Kelas           string `json:"kelas" form:"kelas"`
	Status          string `json:"status" form:"status"`
	HospitalID      uint   `json:"hospital_id" form:"hospital_id"`
}

func (req *BedRequest) reqToCore() bed.BedCore {
	return bed.BedCore{
		NamaTempatTidur: req.NamaTempatTidur,
		Ruangan:         req.Ruangan,
		Kelas:           req.Kelas,
		Status:          req.Status,
		HospitalID:      req.HospitalID,
	}

}
