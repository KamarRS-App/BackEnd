package delivery

import "github.com/KamarRS-App/features/bed"

type BedResponse struct {
	ID              uint   `json:"id"`
	NamaTempatTidur string `json:"nama_tempat_tidur"`
	Ruangan         string `json:"ruangan"`
	Kelas           string `json:"kelas"`
	Status          string `json:"status"`
	HospitalID      uint   `json:"hospital_id"`
	// Hospital        HospitalResponse `json:"hospital"`
}

// type HospitalResponse struct {
// 	ID   uint   `json:"id"`
// 	Nama string `json:"nama"`
// }

// -----------------Bed--------------------------------
func fromCore(dataCore bed.BedCore) BedResponse {
	return BedResponse{
		ID:              dataCore.ID,
		NamaTempatTidur: dataCore.NamaTempatTidur,
		Ruangan:         dataCore.Ruangan,
		Kelas:           dataCore.Kelas,
		Status:          dataCore.Status,
		HospitalID:      dataCore.HospitalID,
	}
}

// data dari core ke response
func fromCoreList(dataCore []bed.BedCore) []BedResponse {
	var dataResponse []BedResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}

// // -----------------Hospital--------------------------------
// func fromCore2(dataCore bed.HospitalCore) HospitalResponse {
// 	return HospitalResponse{
// 		ID:   dataCore.ID,
// 		Nama: dataCore.Nama,
// 	}
// }

// // data dari core ke response
// func fromCoreList2(dataCore []bed.HospitalCore) []HospitalResponse {
// 	var dataResponse []HospitalResponse
// 	for _, v := range dataCore {
// 		dataResponse = append(dataResponse, fromCore2(v))
// 	}
// 	return dataResponse
// }
