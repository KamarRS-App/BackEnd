package delivery

type CheckupReservationRequest struct {
	PatientID  uint `json:"patient_id" form:"patient_id"`
	PracticeID uint `json:"practice_id" form:"practice_id"`
}
