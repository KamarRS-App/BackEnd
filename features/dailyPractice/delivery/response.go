package delivery

import (
	dailypractice "github.com/KamarRS-App/KamarRS-App/features/dailyPractice"
)

type DailyPracticeResponse struct {
	ID             uint   `json:"id"`
	TanggalPraktik string `json:"tanggal_praktik"`
	KuotaHarian    int    `json:"kuota_harian"`
	Status         string `json:"status"`
	PoliclinicID   uint   `json:"policlinic_id"`
}

// -----------------Daily Practice--------------------
func fromCore(dataCore dailypractice.DailyPracticeCore) DailyPracticeResponse {
	return DailyPracticeResponse{
		ID:             dataCore.ID,
		TanggalPraktik: dataCore.TanggalPraktik,
		KuotaHarian:    dataCore.KuotaHarian,
		Status:         dataCore.Status,
		PoliclinicID:   dataCore.PoliclinicID,
	}
}

// data dari core ke response
func fromCoreList(dataCore []dailypractice.DailyPracticeCore) []DailyPracticeResponse {
	var dataResponse []DailyPracticeResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}

//-----------------------------------------------------
