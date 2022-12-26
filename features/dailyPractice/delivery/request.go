package delivery

import "github.com/KamarRS-App/features/dailypractice"

type DailyPracticeRequest struct {
	TanggalPraktik string `json:"tanggal_praktik" form:"tanggal_praktik"`
	KuotaHarian    int    `json:"kuota_harian" form:"kuota_harian"`
	Status         string `json:"status" form:"status"`
	PoliclinicID   uint   `json:"policlinic_id" form:"policlinic_id"`
}

func (req *DailyPracticeRequest) reqToCore() dailypractice.DailyPracticeCore {
	return dailypractice.DailyPracticeCore{
		TanggalPraktik: req.TanggalPraktik,
		KuotaHarian:    req.KuotaHarian,
		Status:         req.Status,
		PoliclinicID:   req.PoliclinicID,
	}
}
