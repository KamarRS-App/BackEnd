package delivery

type CheckupReservationResponse struct {
	PatientID  uint `json:"patient_id"`
	PracticeID uint `json:"practice_id"`
}
