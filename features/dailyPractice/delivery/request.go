package delivery

type DailyPracticeRequest struct {
	TanggalPraktik string `json:"tanggal_praktik" form:"tanggal_praktik"`
	KuotaHarian    int    `json:"kuota_harian" form:"kuota_harian"`
	Status         string `json:"status" form:"status"`
	PoliclinicID   uint   `json:"policlinic_id" form:"policlinic_id"`
}
