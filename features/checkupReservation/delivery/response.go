package delivery

import (
	checkupreservation "kamarRS/features/checkupReservation"
)

type CheckupReservationResponse struct {
	ID         uint `json:"id"`
	PatientID  uint `json:"patient_id"`
	PracticeID uint `json:"practice_id"`
}

// -----------------Checkup Reserve--------------------
func fromCore(dataCore checkupreservation.CheckupReservationCore) CheckupReservationResponse {
	return CheckupReservationResponse{
		ID:         dataCore.ID,
		PatientID:  dataCore.PatientID,
		PracticeID: dataCore.PracticeID,
	}
}

// data dari core ke response
func fromCoreList(dataCore []checkupreservation.CheckupReservationCore) []CheckupReservationResponse {
	var dataResponse []CheckupReservationResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}

//-----------------------------------------------------
