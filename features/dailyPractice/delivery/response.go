package delivery

type DailyPracticeResponse struct {
	ID             uint   `json:"id"`
	TanggalPraktik string `json:"tanggal_praktik"`
	KuotaHarian    int    `json:"kuota_harian"`
	Status         string `json:"status"`
	PoliclinicID   uint   `json:"policlinic_id"`
}
