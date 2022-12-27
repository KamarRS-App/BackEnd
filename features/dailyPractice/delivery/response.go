package delivery

import (
	"github.com/KamarRS-App/features/dailypractice"
)

type PracticeResponse struct {
	ID             uint   `json:"id"`
	TanggalPraktik string `json:"tanggal_praktik"`
	KuotaHarian    int    `json:"kuota_harian"`
	Status         string `json:"status"`
	PoliclinicID   uint   `json:"policlinic_id"`
}

// -----------------Daily Practice--------------------
func FromCore(dataCore dailypractice.PracticeCore) PracticeResponse {
	return PracticeResponse{
		ID:             dataCore.ID,
		TanggalPraktik: dataCore.TanggalPraktik,
		KuotaHarian:    dataCore.KuotaHarian,
		Status:         dataCore.Status,
		PoliclinicID:   dataCore.PoliclinicID,
	}
}

// data dari core ke response
func FromCoreList(dataCore []dailypractice.PracticeCore) []PracticeResponse {
	var dataResponse []PracticeResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, FromCore(v))
	}
	return dataResponse
}

//-----------------------------------------------------
