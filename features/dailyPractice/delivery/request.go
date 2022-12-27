package delivery

import dailypractice "github.com/KamarRS-App/KamarRS-App/features/dailyPractice"

type PracticeRequest struct {
	TanggalPraktik string `json:"tanggal_praktik" form:"tanggal_praktik"`
	KuotaHarian    int    `json:"kuota_harian" form:"kuota_harian"`
	Status         string `json:"status" form:"status"`
	PoliclinicID   uint   `json:"policlinic_id" form:"policlinic_id"`
}

func (req *PracticeRequest) ToCore() dailypractice.PracticeCore {
	return dailypractice.PracticeCore{
		TanggalPraktik: req.TanggalPraktik,
		KuotaHarian:    req.KuotaHarian,
		Status:         req.Status,
		PoliclinicID:   req.PoliclinicID,
	}
}
